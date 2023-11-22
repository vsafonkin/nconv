.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

build: fmt
	go build -ldflags='-s -w' -o ./bin/nconv main.go
.PHONY:build
