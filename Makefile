# Simple Makefile for running tests
CURRENT_DAY := $(shell date "+%d" | grep -o '[1-9]')

.PHONY: test
test:
	go test -v ./cmd/$(CURRENT_DAY)/

.PHONY: run
run:
	go run ./cmd/$(CURRENT_DAY)/main.go --input "assets/$(CURRENT_DAY)/input.txt"

.PHONY: new
new:
	mkdir -p assets/$(CURRENT_DAY) && \
		mkdir -p cmd/$(CURRENT_DAY) && \
		cp ~/Downloads/input.txt assets/$(CURRENT_DAY)/ && \
		cp tools/boilerplate/* cmd/$(CURRENT_DAY)

.PHONY: all
all: test run