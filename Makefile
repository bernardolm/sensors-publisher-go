.ONESHELL:
.PHONY: build-darwin-arm64 build-linux-amd64 build-linux-armv7 build config deploy format gotools lint reset

ifneq (,$(wildcard ./.env))
	include ./.env
	export
endif

PWD = $(shell pwd)

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 . ./dev/build/build.sh

build-linux-amd64:
	GOOS=linux GOARCH=amd64 . ./dev/build/build.sh

build-linux-armv7:
	GOOS=linux GOARCH=arm GOARM=7 . ./dev/build/build.sh

build: build-darwin-arm64 build-linux-amd64 build-linux-armv7

config:
	ln -sf "$(PWD)/.githooks/pre-commit" "$(PWD)/.git/hooks/pre-commit"

deploy:
	. ./dev/deploy/dockerx.sh

format:
	golangci-lint fmt --verbose ./...

gotools:
	go get -tool $(shell /bin/cat dev/packages.golang)

lint:
	golangci-lint run --verbose ./...

reset:
	reset
