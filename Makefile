# Simple Makefile for getting setup, testing, and running in a conventional manner

CURRENT_DAY ?= $(shell date "+%d" | grep -o '[1-9]')

.PHONY: all
all: test run

.PHONY: test
test:
	go test -v -cover ./pkg/$(CURRENT_DAY)/

.PHONY: test-all
test-all:
	go test -cover ./...

.PHONY: test-util
test-util:
	go test -v -cover -bench=. ./pkg/util

.PHONY: run
run:
	go run ./cmd/$(CURRENT_DAY)/main.go --input "assets/$(CURRENT_DAY)/input.txt"

.PHONY: new
new:
	mkdir assets/$(CURRENT_DAY) && \
		mkdir cmd/$(CURRENT_DAY) && \
		mkdir test/testdata/$(CURRENT_DAY) && \
		mv ~/Downloads/input.txt assets/$(CURRENT_DAY)/input.txt && \
		cp tools/boilerplate/dayDAYX.go pkg/$(CURRENT_DAY)/day$(CURRENT_DAY).go && \
		cp tools/boilerplate/dayDAYX_test.go pkg/$(CURRENT_DAY)/day$(CURRENT_DAY)_test.go && \
		gsed -i 's/DAYX/$(CURRENT_DAY)/' cmd/$(CURRENT_DAY)/main.go && \
		gsed -i 's/DAYX/$(CURRENT_DAY)/' pkg/$(CURRENT_DAY)/day$(CURRENT_DAY)_test.go && \
		gsed -i 's/DAYX/$(CURRENT_DAY)/' pkg/$(CURRENT_DAY)/day$(CURRENT_DAY).go && \
		touch test/testdata/$(CURRENT_DAY)/test_input.txt