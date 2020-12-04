# Simple Makefile for running tests

.PHONY: test
test:
	go test -v ./cmd/$$(date "+%d" | grep -o '[1-9]')/

.PHONY: run
run:
	go run ./cmd/$$(date "+%d" | grep -o '[1-9]')/main.go

.PHONY: all
all: test