package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"os"

	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lthibault/log"
	boost "github.com/primev/builder-boost/pkg"
	"github.com/primev/builder-boost/pkg/boostcli"
	"github.com/primev/builder-boost/pkg/p2p/node"
	"github.com/primev/builder-boost/pkg/preconf"
	"github.com/primev/builder-boost/pkg/rollup"
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
		EnvVars: []string{"AGENT_ADDR"},
	},
	&cli.StringFlag{
		Name:    "rollupkey",
		Usage:   "Private key to interact with rollup",
		Value:   "",
		EnvVars: []string{"ROLLUP_KEY"},
	},
	&cli.StringFlag{
		Name:    "rollupaddr",
		Usage:   "Rollup RPC address",
		Value:   "https://ethereum-sepolia.blockpi.network/v1/rpc/public",
		EnvVars: []string{"ROLLUP_ADDR"},
	},
	&cli.StringFlag{
		Name:    "rollupcontract",
		Usage:   "Rollup contract address",
		Value:   "0x6219a236EFFa91567d5ba4a0A5134297a35b0b2A",
		EnvVars: []string{"ROLLUP_CONTRACT"},
	},
	&cli.StringFlag{
		Name:    "buildertoken",
		Usage:   "Token used to authenticate request as originating from builder",
		Value:   "",
		EnvVars: []string{"BUILDER_AUTH_TOKEN"},
	},
	&cli.BoolFlag{
		Name:    "metrics",
		Usage:   "enables metrics tracking for boost",
		Value:   false,
		EnvVars: []string{"METRICS"},
	},
	&cli.StringFlag{
		Name:    "dacontract",
		Usage:   "DA contract address",
		Value:   "0xac27A2cbdBA8768D49e359ebA326fC1F27832ED4",
		EnvVars: []string{"ROLLUP_CONTRACT"},
	},
	&cli.StringFlag{
		Name:    "daaddr",
		Usage:   "DA RPC address",
		Value:   "http://54.200.76.18:8545",
		EnvVars: []string{"ROLLUP_ADDR"},
	},
	&cli.BoolFlag{
		Name:    "inclusionlist",
		Usage:   "enables inclusion list for boost",
		Value:   false,
		EnvVars: []string{"INCLUSION_LIST"},
	},
}

var (
	config = boost.Config{Log: log.New()}
	svr    http.Server
)

// TODO(@ckartik): Intercept SIGINT and SIGTERM to gracefully shutdown the server and persist state to cache

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
			Log: boostcli.Logger(c),
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
		builderKeyString := c.String("rollupkey")
		if builderKeyString == "" {
			return errors.New("rollup key is not set, use --rollupkey option or ROLLUP_KEY env variable")
		}
		builderAuthToken := c.String("buildertoken")
		if builderAuthToken == "" {
			return errors.New("builder token is not set, use --buildertoken option or BUILDER_AUTH_TOKEN env variable")
		}

		builderKeyBytes := common.FromHex(builderKeyString)
		builderKey := crypto.ToECDSAUnsafe(builderKeyBytes)

		client, err := ethclient.Dial(c.String("rollupaddr"))
		if err != nil {
			return err
		}

		contractAddress := common.HexToAddress(c.String("rollupcontract"))
		ru, err := rollup.New(client, contractAddress, builderKey, config.Log)
		if err != nil {
			return err
		}
		buildernode := node.NewBuilderNode(config.Log, builderKey, ru, nil)
		select {
		case <-buildernode.Ready():
		}
		go func() {
			client, err := ethclient.Dial("http://54.200.76.18:8545")
			if err != nil {
				log.Fatalf("Failed to connect to the Ethereum client: %v", err)
			}
			for peerMsg := range buildernode.BidReader() {
				var pc preconf.PreConfBid
				err := json.Unmarshal(peerMsg.Bytes, &pc)
				if err != nil {
					config.Log.WithError(err).Error("failed to unmarshal preconf bid")
				}
				address, err := pc.VerifySearcherSignature()
				if err != nil {
					config.Log.WithError(err).Error("failed to verify preconf bid")
				}
				config.Log.WithField("address", address.Hex()).WithField("bid_txn", pc.TxnHash).WithField("bid_amt", pc.GetBidAmt()).WithField("sending_peer", peerMsg.Peer).Info("preconf bid verified")
				commit, err := pc.ConstructCommitment(builderKey)
				if err != nil {
					config.Log.WithError(err).Error("failed to construct commitment")
				}

				config.Log.WithField("commitment", commit).Info("commitment constructed")
				commit.SendCommitmentToSearcher(buildernode)

				txn, err := commit.StoreCommitmentToDA(builderKey, "0xac27A2cbdBA8768D49e359ebA326fC1F27832ED4", client)
				if err != nil {
					config.Log.WithError(err).Error("failed to store commitment to DA")
					continue
				}

				config.Log.WithField("txn", txn.Hash().Hex()).Info("commitment stored to DA")
				// http call to send commitment to Geth Node
				// Send RPC call as follows  curl -X POST --data '{"jsonrpc":"2.0","method":"eth_sendPreconfirmationBid","params":["0x927452e78b79db883d3652284245f1de5087efabba620d601098c9ae2ac8a942"],"id":1}' -H "Content-Type: application/json" http://localhost:8545
				request := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_sendPreconfirmationBid\",\"params\":[\"" + txn.Hash().Hex() + "\"],\"id\":1}"
				config.Log.WithField("request", request).Info("sending request to Geth Node")

				// // Start json rpc client using net/rpc package
				// gethClient, err := rpc.DialHTTP("tcp", "localhost:8545")
				// if err != nil {
				// 	config.Log.WithError(err).Error("failed to dial Geth Node")
				// 	continue
				// }
				// err = gethClient.Call("eth_sendPreconfirmationBid", txn.Hash().Hex(), nil)
				resp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer([]byte(request)))
				if err != nil {
					config.Log.WithError(err).Error("failed to send request to Geth Node")
					continue
				}
				resp.Body.Close()
				// config.Log.WithField("response", resp).Info("response from Geth Node")

				// config.Log.WithField("peer", peerMsg.Peer).Info(string(peerMsg.Bytes))
			}
		}()

		g.Go(func() error {
			return ru.Run(ctx)
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

			svr.Handler = &boost.API{
				Service:        service,
				Log:            config.Log,
				Worker:         masterWorker,
				Rollup:         ru,
				BuilderToken:   c.String("buildertoken"),
				MetricsEnabled: c.Bool("metrics"),
				InclusionList:  c.Bool("inclusionlist"),
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
