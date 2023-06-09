dep:
	go mod tidy
	go mod vendor

test:
	cd pkg
	go test -v ./...

build:
	go build ./cmd/boost
	go build ./cmd/searcher

all:
	make dep
	make test
	make build