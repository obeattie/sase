SHELL := /bin/sh

.PHONY: compile-parser
compile-parser:
	ragel -Z -G2 query/tokeniser.rl
	gofmt -s -w query/tokeniser.go

.PHONY: generate
generate:
	go install golang.org/x/tools/cmd/stringer
	go generate "./..."

.PHONY: build
build: compile-parser generate
	go build
