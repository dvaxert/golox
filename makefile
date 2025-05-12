# Detect OS and binary extension
ifeq ($(OS),Windows_NT)
    EXE := .exe
else
    EXE :=
endif

# Go variables
GO      := go
APP     := golox
VERSION := 1.0.0

# Binaries
BIN_DIR := bin
CMD_DIR := cmd
CMDS    := lox

# Default target
.PHONY: all
all: build

# Build all binaries
.PHONY: build
build:
	@mkdir -p $(BIN_DIR)
	@for cmd in $(CMDS); do \
		echo "Building $$cmd..."; \
		$(GO) build -o $(BIN_DIR)/$$cmd$(EXE) ./$(CMD_DIR)/$$cmd; \
	done

# Clean build artifacts
.PHONY: clean
clean:
	@rm -rf $(BIN_DIR)

# Run specific binary (example: make run CMD=server)
.PHONY: run
run:
	$(GO) run ./cmd/$(CMD) $(ARGS)

# Run tests
.PHONY: test
test:
	$(GO) test ./...

# Help
.PHONY: help
help:
	@echo "Makefile targets:"
	@echo "  build     - Build all binaries"
	@echo "  clean     - Remove build artifacts"
	@echo "  run CMD=x - Run command (e.g. make run CMD=cli)"
	@echo "  test      - Run all tests"
	@echo "  tidy      - Run go mod tidy"
