BIN_DIR       := $(CURDIR)/bin
CMD_DIR       := $(CURDIR)/cmd
TOOLS_BIN_DIR := $(BIN_DIR)/tools
TOOLS_SRC_DIR := $(CURDIR)/tools
INSTALL_PATH  ?= /usr/local/bin
TARGETS       := darwin/amd64 darwin/arm64 linux/amd64 windows/amd64

.PHONY: all
all: build-tools
	go build -o $(BIN_DIR) $(CMD_DIR)/pega

.PHONY: clean
clean:
	rm -Rf $(BIN_DIR)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build-tools
build-tools:
	go build -o $(TOOLS_BIN_DIR)/inebase $(TOOLS_SRC_DIR)/inebase/main.go
	go build -o $(TOOLS_BIN_DIR)/infoelectoral $(TOOLS_SRC_DIR)/infoelectoral/main.go

.PHONY: migrate-down
migrate-down:
	migrate -path database/migration -database sqlite3://database.sqlite down

.PHONY: migrate-up
migrate-up:
	migrate -path database/migration -database sqlite3://database.sqlite up

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...

.PHONY: test-coverage-report
test-coverage-report: test-coverage
	go tool cover -html=coverage.out
