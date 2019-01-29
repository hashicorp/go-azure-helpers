GO111MODULE=on

default: test

dependencies:
	go mod download

test: dependencies
	go vet ./...
	go test -race ./...

.PHONY: default test
