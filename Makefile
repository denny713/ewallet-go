SHELL := /bin/bash
BASE := $(shell pwd)
# PKGS := $(shell go list ./... | grep -v /vendor/ | grep -v /internal/mock)
# ALL_PACKAGES := $(shell go list ./... | grep -v /vendor/ | grep -v /internal/mock)
# GIT_COMMIT := $(shell git rev-list -1 HEAD)
BUILD_TIME := $(shell date +%FT%T%z)
ALL_PACKAGES := $(shell go list ./... | grep -v /vendor/)
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go-e-wallet
# BINARY_UNIX=$(BINARY_NAME)_unix
fmt:
	go fmt $(ALL_PACKAGES)

lint:
	env GO111MODULE=on golint -set_exit_status $(ALL_PACKAGES)

generate: fmt lint
	go generate -mod=vendor ./...

build_dev:
	# pwd
	#go mod tidy && go mod vendor
	go mod tidy && go mod vendor
	git status
	GO111MODULE=on $(GOBUILD) -mod=vendor -o $(BINARY_NAME) -v main.go

build_beta:
	# pwd
	#go mod tidy && go mod vendor
	go mod tidy && go mod vendor
	git status
	GO111MODULE=on $(GOBUILD) -mod=vendor -o $(BINARY_NAME_BETA) -v main.go


build_prod:
	# pwd
	go mod tidy && go mod vendor
	git status
	GO111MODULE=on $(GOBUILD) -mod=vendor -o $(BINARY_NAME_PROD) -v main.go