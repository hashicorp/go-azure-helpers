default: test

test:
	go vet ./...
	go test -race ./...

.PHONY: default test
