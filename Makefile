# Simple Makefile for getting setup, testing, and running in a conventional manner

CURRENT_DAY ?= $(shell date "+%d" | grep -o '[0-9]*')

.PHONY: all
all: test run

.PHONY: test
test:
	gotest -v -cover -failfast ./pkg/$(CURRENT_DAY)/

.PHONY: test-watch
test-watch:
	goconvey -workDir ./pkg/$(CURRENT_DAY)/

.PHONY: test-all
test-all:
	gotest -cover ./...

.PHONY: test-utils
test-utils:
	gotest -v -cover -bench=. ./pkg/utils

.PHONY: run
run:
	go run ./cmd/main.go --input "assets/$(CURRENT_DAY)/input.txt" --day "$(CURRENT_DAY)"

.PHONY: new
new:
	mkdir assets/$(CURRENT_DAY) && \
		mv ~/Downloads/input.txt assets/$(CURRENT_DAY)/input.txt && \
		touch assets/$(CURRENT_DAY)/instructions.md && \
		mkdir pkg/$(CURRENT_DAY) && \
		cp tools/boilerplate/DAYX.go pkg/$(CURRENT_DAY)/day$(CURRENT_DAY).go && \
		cp tools/boilerplate/DAYX_test.go pkg/$(CURRENT_DAY)/day$(CURRENT_DAY)_test.go && \
		gsed -i 's/DAYX/$(CURRENT_DAY)/' pkg/$(CURRENT_DAY)/day$(CURRENT_DAY)_test.go && \
		gsed -i 's/DAYX/$(CURRENT_DAY)/' pkg/$(CURRENT_DAY)/day$(CURRENT_DAY).go && \
		mkdir test/testdata/$(CURRENT_DAY) && \
		touch test/testdata/$(CURRENT_DAY)/test_input.txt

.PHONY: clean
clean:
	trash assets/$(CURRENT_DAY) && \
		trash pkg/$(CURRENT_DAY) && \
		trash test/testdata/$(CURRENT_DAY)

