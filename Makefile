.PHONY: build test

build:
		go build -v ./cmd/main.go

run:
		go run ./cmd/main.go

test:
		go test -v -race -timeout=30s ./...

.DEFAULT_GOAL := build