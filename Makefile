.PHONY: build test

build:
		go build -v ./cmd/botServer/botServer.go

run:
		go run ./cmd/botServer/botServer.go

test:
		go test -v -race -timeout=30s ./...

.DEFAULT_GOAL := build