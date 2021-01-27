GO := go
GO_BUILD := $(GO) build

DIST_DIR := ./dist
CMD_DIR := ./cmd

.PHONY: build
build:
	$(GO_BUILD) -o dist/url-shortener ./cmd/...
