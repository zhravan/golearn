# GoLearn Makefile

SHELL := /bin/sh

GO := go
PKG := ./cmd/golearn
BIN_DIR := bin
BIN := $(BIN_DIR)/golearn

.PHONY: help build rebuild run tidy init list verify hint progress watch reset test fmt fmt-check vet check clean ci

default: help

help:
	@echo "GoLearn - Rustlings-like Go learning CLI"
	@echo ""
	@echo "Common targets:"
	@echo "  build             Build CLI binary to $(BIN)"
	@echo "  rebuild           Force rebuild the CLI binary"
	@echo "  run ARGS=...      Run CLI via 'go run' (e.g., make run ARGS=\"list\")"
	@echo "  init [REPO] [DIR] Initialize workspace: clone REPO into DIR or copy built-in templates"
	@echo "  list              List exercises"
	@echo "  verify NAME=...   Verify an exercise (or all if NAME unset)"
	@echo "  hint NAME=...     Show hints for an exercise"
	@echo "  progress          Show progress dashboard with ASCII bar"
	@echo "  watch             Watch exercises and auto-verify on changes"
	@echo "  reset NAME=...    Reset an exercise to starter state"
	@echo ""
	@echo "Quality targets:"
	@echo "  tidy              Run 'go mod tidy'"
	@echo "  fmt               Format code"
	@echo "  fmt-check         Fail if any files need formatting"
	@echo "  vet               Run 'go vet'"
	@echo "  test              Run all tests in ./exercises/..."
	@echo "  check             Run fmt-check, vet, build, and tests"
	@echo ""
	@echo "CI:"
	@echo "  ci                Run a sensible CI pipeline"

$(BIN):
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN) $(PKG)

# Always build the binary when invoking 'make build'
build:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN) $(PKG)

# Force rebuild convenience target
rebuild:
	rm -f $(BIN)
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN) $(PKG)

run:
	$(GO) run $(PKG) $(ARGS)

tidy:
	$(GO) mod tidy

init:
	$(GO) run $(PKG) init $(REPO) $(DIR)

list:
	$(GO) run $(PKG) list

verify:
	$(GO) run $(PKG) verify $(NAME)

hint:
	@if [ -z "$(NAME)" ]; then echo "Usage: make hint NAME=<exercise-slug>"; exit 1; fi
	$(GO) run $(PKG) hint $(NAME)

progress:
	$(GO) run $(PKG) progress

watch:
	$(GO) run $(PKG) watch

reset:
	@if [ -z "$(NAME)" ]; then echo "Usage: make reset NAME=<exercise-slug>"; exit 1; fi
	$(GO) run $(PKG) reset $(NAME)

test:
	$(GO) test ./exercises/...

fmt:
	$(GO) fmt ./...

fmt-check:
	@files=$$(gofmt -l .); \
	if [ -n "$$files" ]; then \
	  echo "Unformatted files:"; echo "$$files"; \
	  echo "Run 'make fmt' to fix"; \
	  exit 1; \
	fi

vet:
	$(GO) vet ./...

check: fmt-check vet build test

clean:
	rm -rf $(BIN_DIR)

 

ci: tidy fmt-check vet build test


