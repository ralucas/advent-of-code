# Simple Makefile for getting setup, testing, and running in a conventional manner

DAY ?= $(shell date "+%d" | grep -o '[0-9]*')

.PHONY: all
all: test run

.PHONY: test
test:
	gotest -v -cover -failfast ./pkg/$(DAY)/

.PHONY: test-watch
test-watch:
	goconvey -workDir ./pkg/$(DAY)/

.PHONY: test-all
test-all:
	gotest -cover ./...

.PHONY: test-utils
test-utils:
	gotest -v -cover -bench=. ./pkg/utils

.PHONY: run
run:
	go run ./cmd/main.go --input "assets/$(DAY)/input.txt" --day "$(DAY)"

.PHONY: new
new:
	mkdir assets/$(DAY) && \
		mv ~/Downloads/input.txt assets/$(DAY)/input.txt && \
		touch assets/$(DAY)/instructions.md && \
		mkdir pkg/$(DAY) && \
		cp tools/boilerplate/DAYX.go pkg/$(DAY)/day$(DAY).go && \
		cp tools/boilerplate/DAYX_test.go pkg/$(DAY)/day$(DAY)_test.go && \
		gsed -i 's/DAYX/$(DAY)/' pkg/$(DAY)/day$(DAY)_test.go && \
		gsed -i 's/DAYX/$(DAY)/' pkg/$(DAY)/day$(DAY).go && \
		mkdir test/testdata/$(DAY) && \
		touch test/testdata/$(DAY)/test_input.txt

.PHONY: clean
clean:
	trash assets/$(DAY) && \
		trash pkg/$(DAY) && \
		trash test/testdata/$(DAY)

