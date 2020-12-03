# Simple Makefile for running tests

.PHONY: test
test:
	go test -v ./cmd/$$(date "+%d" | grep -o '[1-9]')/

.PHONY: all
all: test