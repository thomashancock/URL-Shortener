GO := go
GO_BUILD := $(GO) build
GO_TEST := $(GO) test
GO_CLEAN := $(GO) clean

NAME := url-shortener
DIST_DIR := ./dist
CMD_DIR := ./cmd

.PHONY: build
build:
	$(GO_BUILD) -v -o $(DIST_DIR)/$(NAME) $(CMD_DIR)/$(NAME)/...

test:
	$(GO_TEST) -v ./...

clean:
	rm $(DIST_DIR)/$(NAME)
