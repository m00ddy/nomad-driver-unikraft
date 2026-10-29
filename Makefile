.PHONY: fmt lint test build ci

fmt:
	go fmt ./...
	go run golang.org/x/tools/cmd/goimports -w .

lint:
	golangci-lint run ./...

test:
	go test -v -race ./...

build:
# 	go build -o bin/nomad-unikraft-plugin ./...
# 	compile check:
	go build -v ./...

clean:
	rm -rf bin/*

ci: fmt lint test build