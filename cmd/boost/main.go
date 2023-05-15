package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"strconv"

	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lthibault/log"
	boost "github.com/primev/builder-boost/pkg"
	"github.com/primev/builder-boost/pkg/rollup"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const (
	shutdownTimeout = 5 * time.Second
	version         = boost.Version
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "loglvl",
		Usage:   "logging level: trace, debug, info, warn, error or fatal",
		Value:   "info",
		EnvVars: []string{"LOGLVL"},
	},
	&cli.StringFlag{
		Name:    "logfmt",
		Usage:   "format logs as text, json or none",
		Value:   "text",
		EnvVars: []string{"LOGFMT"},
	},
	&cli.StringFlag{
		Name:    "addr",
		Usage:   "server listen address",
		Value:   ":18550",
		EnvVars: []string{"BOOST_ADDR"},
	},
	&cli.StringFlag{
		Name:    "env",
		Usage:   "service environment (development, production, etc.)",
		Value:   "development",
		EnvVars: []string{"ENV"},
	},
	&cli.StringFlag{
		Name:    "agentaddr",
		Usage:   "datadog agent address",
		Value:   "",
		EnvVars: []string{"AGENTADDR"},
	},
	&cli.StringFlag{
		Name:    "rollupkey",
		Usage:   "Private key to interact with rollup",
		Value:   "",
		EnvVars: []string{"ROLLUPKEY"},
	},
	&cli.StringFlag{
		Name:    "rollupaddr",
		Usage:   "Rollup RPC address",
		Value:   "https://rpc.sepolia.org",
		EnvVars: []string{"ROLLUPADDR"},
	},
	&cli.StringFlag{
		Name:    "rollupcontract",
		Usage:   "Rollup contract address",
		Value:   "0xB82c160372cd59eD689D84f2632847963F87B212",
		EnvVars: []string{"ROLLUPCONTRACT"},
	},
	&cli.StringFlag{
		Name:    "rollupblock",
		Usage:   "Block at which rollup contract was deployed",
		Value:   "3490558",
		EnvVars: []string{"ROLLUPBLOCK"},
	},
	&cli.StringFlag{
		Name:    "rollupstate",
		Usage:   "Filename of rollup state file",
		Value:   "rollup.json",
		EnvVars: []string{"ROLLUPSTATE"},
	},
}

var (
	config = boost.Config{Log: log.New()}
	svr    http.Server
)

// Main starts the primev protocol
func main() {
	app := &cli.App{
		Name:    "builder boost",
		Usage:   "entry point to primev protocol",
		Version: version,
		Flags:   flags,
		Before:  setup(),
		Action:  run(),
	}

	if err := app.Run(os.Args); err != nil {
		config.Log.Fatal(err)
	}
}

func setup() cli.BeforeFunc {
	return func(c *cli.Context) (err error) {

		config = boost.Config{
			Log: logger(c),
		}

		svr = http.Server{
			Addr:           c.String("addr"),
			ReadTimeout:    c.Duration("timeout"),
			WriteTimeout:   c.Duration("timeout"),
			IdleTimeout:    time.Second * 2,
			MaxHeaderBytes: 4096,
		}

		return
	}
}

func run() cli.ActionFunc {
	return func(c *cli.Context) error {
		g, ctx := errgroup.WithContext(c.Context)

		// setup rollup service
		builderKetString := c.String("rollupkey")
		if builderKetString == "" {
			return errors.New("rollup key is not set, use --rollupkey option or ROLLUPKEY env variable")
		}

		builderKeyBytes := common.FromHex(builderKetString)
		builderKey := crypto.ToECDSAUnsafe(builderKeyBytes)

		client, err := ethclient.Dial(c.String("rollupaddr"))
		if err != nil {
			return err
		}

		rollupBlockStr := c.String("rollupblock")
		rollupBlock, err := strconv.ParseUint(rollupBlockStr, 10, 64)
		if err != nil {
			return err
		}

		contractAddress := common.HexToAddress(c.String("rollupcontract"))
		statePath := c.String("rollupstate")
		rollup, err := rollup.New(client, contractAddress, builderKey, rollupBlock, statePath, config.Log)
		if err != nil {
			return err
		}

		g.Go(func() error {
			return rollup.Run(ctx)
		})

		// setup the boost service
		service := &boost.DefaultService{
			Log:    config.Log,
			Config: config,
		}
		g.Go(func() error {
			return service.Run(ctx)
		})

		// wait for the boost service to be ready
		select {
		case <-service.Ready():
		case <-ctx.Done():
			return g.Wait()
		}

		config.Log.Info("boost service ready")

		masterWorker := boost.NewWorker(service.Boost.GetWorkChannel(), config.Log)
		g.Go(func() error {
			return masterWorker.Run(ctx)
		})
		// wait for the boost service to be ready
		select {
		case <-masterWorker.Ready():
		case <-ctx.Done():
			return g.Wait()
		}

		config.Log.Info("master worker ready")

		// run the http server
		g.Go(func() (err error) {
			// set up datadog tracer
			agentAddr := c.String("agentaddr")
			if len(agentAddr) > 0 {
				tracer.Start(
					tracer.WithService("builder-boost"),
					tracer.WithEnv(c.String("env")),
					tracer.WithAgentAddr(agentAddr),
				)
				defer tracer.Stop()
			}

			svr.BaseContext = func(l net.Listener) context.Context {
				return ctx
			}

			// rl := rollup.New()
			// // Manually creating a rollup
			// rollup := boost.NewStubRollup()
			// rollup.SetBalance("kartik", 32)
			// rollup.SetBalance("murat", 32)
			// rollup.SetBalance("dan", 32)
			// rollup.SetBalance("kant", 32)
			// rollup.SetBalance("serhii", 32)
			// rollup.SetBalance("justenough", 31)

			svr.Handler = &boost.API{
				Service: service,
				Log:     config.Log,
				Worker:  masterWorker,
				Rollup:  rollup,
			}

			config.Log.Info("http server listening")
			if err = svr.ListenAndServe(); err == http.ErrServerClosed {
				err = nil
			}

			return err
		})

		g.Go(func() error {
			defer svr.Close()
			<-ctx.Done()

			ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
			defer cancel()

			return svr.Shutdown(ctx)
		})

		return g.Wait()
	}
}

func logger(c *cli.Context) log.Logger {
	return log.New(
		withLevel(c),
		withFormat(c),
		withErrWriter(c))
}

func withLevel(c *cli.Context) (opt log.Option) {
	var level = log.FatalLevel
	defer func() {
		opt = log.WithLevel(level)
	}()

	if c.Bool("trace") {
		level = log.TraceLevel
		return
	}

	if c.String("logfmt") == "none" {
		return
	}

	switch c.String("loglvl") {
	case "trace", "t":
		level = log.TraceLevel
	case "debug", "d":
		level = log.DebugLevel
	case "info", "i":
		level = log.InfoLevel
	case "warn", "warning", "w":
		level = log.WarnLevel
	case "error", "err", "e":
		level = log.ErrorLevel
	case "fatal", "f":
		level = log.FatalLevel
	default:
		level = log.InfoLevel
	}

	return
}

func withFormat(c *cli.Context) log.Option {
	var fmt logrus.Formatter

	switch c.String("logfmt") {
	case "none":
	case "json":
		fmt = &logrus.JSONFormatter{PrettyPrint: c.Bool("prettyprint")}
	default:
		fmt = new(logrus.TextFormatter)
	}

	return log.WithFormatter(fmt)
}

func withErrWriter(c *cli.Context) log.Option {
	return log.WithWriter(c.App.ErrWriter)
}
