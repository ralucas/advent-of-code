# Simple Makefile for getting setup, testing, and running in a conventional manner

CURRENT_DAY ?= $(shell date "+%d" | grep -o '[1-9]')

.PHONY: all
all: test run

.PHONY: test
test:
	go test -v ./cmd/$(CURRENT_DAY)/

.PHONY: test-all
test-all:
	go test -v ./...

.PHONY: run
run:
	go run ./cmd/$(CURRENT_DAY)/main.go --input "assets/$(CURRENT_DAY)/input.txt"

.PHONY: new
new:
	mkdir assets/$(CURRENT_DAY) && \
		mkdir cmd/$(CURRENT_DAY) && \
		mkdir test/testdata/$(CURRENT_DAY) && \
		cp ~/Downloads/input.txt assets/$(CURRENT_DAY)/ && \
		cp tools/boilerplate/* cmd/$(CURRENT_DAY) && \
		gsed -i 's/%%DAY%%/$(CURRENT_DAY)/' cmd/$(CURRENT_DAY)/main.go && \
		touch test/testdata/$(CURRENT_DAY)/test_input.txt