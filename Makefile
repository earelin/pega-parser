BIN_DIR       := $(CURDIR)/bin
CMD_DIR       := $(CURDIR)/cmd
INSTALL_PATH  ?= /usr/local/bin
TARGETS       := darwin/amd64 darwin/arm64 linux/amd64 windows/amd64

.PHONY: all
all:
	go build -o $(BIN_DIR)/pega $(CMD_DIR)/pega

.PHONY: clean
clean:
	rm -Rf $(BIN_DIR)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: migrate-down
migrate-down:
	migrate -path database/migration -database sqlite3://database.sqlite down

.PHONY: migrate-up
migrate-up:
	migrate -path database/migration -database sqlite3://database.sqlite up

.PHONY: run
run:
	go run $(CMD_DIR)/pega

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...

.PHONY: test-coverage-report
test-coverage-report: test-coverage
	go tool cover -html=coverage.out
