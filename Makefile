.PHONY: fmt lint test build ci

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

test:
	go test -v -race ./...

build:
	go build -o bin/nomad-unikraft-plugin ./...

ci: fmt lint test build