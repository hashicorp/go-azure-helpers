default: test

dependencies:
	GO111MODULE=on go mod download

test: dependencies
	go vet ./...
	go test -race ./...

.PHONY: default test
