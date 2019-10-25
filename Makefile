SHELL := /bin/bash
PLATFORM := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)
GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

GO_PACKAGE := github.com/zhanglianxin/jop-go
CROSS_TARGETS := linux/amd64 darwin/amd64 windows/386 windows/amd64
BIN_FILE := jop-go

default: build cross gen-sha1

get-deps:
	dep ensure

cp-config:
	cp config_example.toml config.toml

build:
	go fmt ./...
	@go build

clean:
	rm -fr data/*

cross:
	gox -osarch="$(CROSS_TARGETS)" $(GO_PACKAGE)
	@$(MAKE) gen-sha1

rm-sha1:
	@rm -f $(BIN_FILE)_*.sha1

gen-sha1: rm-sha1
	@$$(for f in $$(find $(BIN_FILE)_* -type f); do shasum $$f > $$f.sha1; done)

zip:
	zip $(BIN_FILE).zip $(BIN_FILE)_* -x *.log

tar:
	tar -czvf --exclude=*.log $(BIN_FILE).tar.gz $(BIN_FILE)_*
