BINDIR       := $(CURDIR)/bin
INSTALL_PATH ?= /usr/local/bin
TARGETS      := darwin/amd64 darwin/arm64 linux/amd64 windows/amd64

all: build build-tools
.PHONY: all

build:
.PHONY: build

build-tools: build
.PHONY: build-tools