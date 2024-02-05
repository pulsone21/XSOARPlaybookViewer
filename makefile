.PHONY: help
include .env
export

refresh_deps:
	@go mod tidy

build: refresh_deps
	@go build -o bin/main ./cmd/main.go

run: build
	@./bin/main

bundle: 
	@go run ./internal/react/builder.go true true
