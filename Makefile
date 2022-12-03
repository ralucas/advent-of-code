# Simple Makefile for getting setup, testing, and running in a conventional manner

DAY ?= $(shell TZ=America/New_York date "+%-d")
YEAR ?= $(shell [ "12" -eq "$$(TZ=America/New_York date +%m)" ] && date "+%Y" || echo "$$(($$(date +%Y)-1))")

.PHONY: all
all: build test-all

.PHONY: build
build:
	mkdir -p bin
	go build -o bin/ -v ./...

.PHONY: test
test:
	go test -v -cover -failfast -benchmem -bench=. ./pkg/$(YEAR)/$(DAY)/

.PHONY: test-watch
test-watch:
	goconvey -workDir ./pkg/$(YEAR)/$(DAY)/

.PHONY: test-all
test-all:
	go test -cover -benchmem -bench=. ./...

.PHONY: test-current-year
test-current-year:
	go test -v -cover -benchmem -bench=. ./pkg/$(YEAR)/...

.PHONY: test-util
test-util:
	go test -v -cover -benchmem -bench=. ./pkg/util/...

.PHONY: lint
lint:
	golangci-lint run -v --enable-all ./pkg/$(YEAR)/$(DAY)

.PHONY: lint-fix
lint-fix:
	golangci-lint run -v --enable-all --fix ./pkg/$(YEAR)/$(DAY)

.PHONY: lint-all
lint-all:
	golangci-lint run -v --enable-all ./...

.PHONY: run
run:
	go run ./cmd/main.go --input "assets/$(YEAR)/$(DAY)/input.txt" --day "$(DAY)" --year "$(YEAR)"

.PHONY: new
new:
	if [ ! -f ~/Downloads/input.txt ]; then echo "no input file"; exit 1; fi; \
	if [ -d pkg/$(YEAR)/$(DAY) ]; then echo "already exists"; exit 1; fi; \
	mkdir -p assets/$(YEAR)/$(DAY) && \
		mv ~/Downloads/input.txt assets/$(YEAR)/$(DAY)/input.txt && \
		touch assets/$(YEAR)/$(DAY)/instructions.md && \
		mkdir -p pkg/$(YEAR)/$(DAY) && \
		cp tools/boilerplate/DAYX.gotpl pkg/$(YEAR)/$(DAY)/day$(DAY).go && \
		cp tools/boilerplate/DAYX_test.gotpl pkg/$(YEAR)/$(DAY)/day$(DAY)_test.go && \
		gsed -i 's/DAYX/$(DAY)/g' pkg/$(YEAR)/$(DAY)/day$(DAY)_test.go && \
		gsed -i 's/YEARX/$(YEAR)/g' pkg/$(YEAR)/$(DAY)/day$(DAY)_test.go && \
		gsed -i 's/DAYX/$(DAY)/g' pkg/$(YEAR)/$(DAY)/day$(DAY).go && \
		gsed -i 's|//newdayimport|day$(DAY) \"github.com/ralucas/advent-of-code/pkg/$(YEAR)/$(DAY)\"\n//newdayimport|' pkg/aoc/days_$(YEAR).go && \
		gsed -i 's|//newdaystruct|\&day$(DAY).Day{},\n\r//newdaystruct|' pkg/aoc/days_$(YEAR).go && \
		mkdir -p test/testdata/$(YEAR)/$(DAY) && \
		touch test/testdata/$(YEAR)/$(DAY)/test_input.txt && \
		goimports -w pkg/aoc/days_$(YEAR).go

.PHONY: clean
clean:
	trash assets/$(YEAR)/$(DAY) && \
		trash pkg/$(YEAR)/$(DAY) && \
		trash test/testdata/$(YEAR)/$(DAY) && \
		trash bin/