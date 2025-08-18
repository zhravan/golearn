# GoLearn Makefile

SHELL := /bin/sh

GO := go
PKG := ./cmd/golearn
BIN_DIR := bin
BIN := $(BIN_DIR)/golearn
DOCKER_IMAGE := golearn:latest

.PHONY: help build run tidy init list verify hint progress reset test fmt fmt-check vet check clean docker-build docker-run ci

default: help

help:
	@echo "GoLearn - Rustlings-like Go learning CLI"
	@echo ""
	@echo "Common targets:"
	@echo "  build             Build CLI binary to $(BIN)"
	@echo "  run ARGS=...      Run CLI via 'go run' (e.g., make run ARGS=\"list\")"
	@echo "  init              Copy embedded templates into ./exercises"
	@echo "  list              List exercises"
	@echo "  verify NAME=...   Verify an exercise (or all if NAME unset)"
	@echo "  hint NAME=...     Show hints for an exercise"
	@echo "  progress          Show local progress"
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
	@echo "Docker:"
	@echo "  docker-build      Build Docker image $(DOCKER_IMAGE)"
	@echo "  docker-run        Run image with current workspace mounted"
	@echo ""
	@echo "CI:"
	@echo "  ci                Run a sensible CI pipeline"

$(BIN):
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN) $(PKG)

build: $(BIN)

run:
	$(GO) run $(PKG) $(ARGS)

tidy:
	$(GO) mod tidy

init:
	$(GO) run $(PKG) init

list:
	$(GO) run $(PKG) list

verify:
	$(GO) run $(PKG) verify $(NAME)

hint:
	@if [ -z "$(NAME)" ]; then echo "Usage: make hint NAME=<exercise-slug>"; exit 1; fi
	$(GO) run $(PKG) hint $(NAME)

progress:
	$(GO) run $(PKG) progress

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

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run --rm -it -v "$$PWD:/workspace" -w /workspace $(DOCKER_IMAGE) $(ARGS)

ci: tidy fmt-check vet build test docker-build


