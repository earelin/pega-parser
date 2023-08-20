BIN_DIR       := $(CURDIR)/bin
TOOLS_BIN_DIR := $(BINDIR)/tools
TOOLS_SRC_DIR := $(CURDIR)/tools
INSTALL_PATH ?= /usr/local/bin
TARGETS      := darwin/amd64 darwin/arm64 linux/amd64 windows/amd64

.PHONY: all
all: build-tools

.PHONY: build-tools
build-tools: $(TOOLS_BIN_DIR)/infoelectoral

INFOELECTORAL_SRC_DIR := $(TOOLS_SRC_DIR)/infoelectoral
$(TOOLS_BIN_DIR)/infoelectoral: $(INFOELECTORAL_SRC_DIR)/main.go
	go build -o $(TOOL_BIN_DIR)/infoelectoral $(INFOELECTORAL_SRC_DIR)/main.go

.PHONY: clean
clean:
	rm -Rf $(BINDIR)
