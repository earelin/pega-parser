BIN_DIR       := $(CURDIR)/bin
TOOLS_BIN_DIR := $(BIN_DIR)/tools
TOOLS_SRC_DIR := $(CURDIR)/tools
INSTALL_PATH  ?= /usr/local/bin
TARGETS       := darwin/amd64 darwin/arm64 linux/amd64 windows/amd64

.PHONY: all
all: build-tools

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build-tools
build-tools:
	go build -o $(TOOLS_BIN_DIR)/infoelectoral $(TOOLS_SRC_DIR)/infoelectoral/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -Rf $(BIN_DIR)
