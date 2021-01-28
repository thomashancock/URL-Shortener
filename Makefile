GO := go
GO_BUILD := $(GO) build

NAME := url-shortener
DIST_DIR := ./dist
CMD_DIR := ./cmd

.PHONY: build
build:
	$(GO_BUILD) -o $(DIST_DIR)/$(NAME) $(CMD_DIR)/$(NAME)/...
