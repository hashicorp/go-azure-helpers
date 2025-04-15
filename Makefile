GO111MODULE=on

default: test

dependencies:
	go mod download

tools:
	go install mvdan.cc/gofumpt@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH || $$GOPATH)/bin v1.64.8

lint:
	@golangci-lint -j 12 run --fast --exclude-dirs="/sdk/" ./...

test: dependencies
	go vet ./...
	go test -race -v ./...

.PHONY: default test
