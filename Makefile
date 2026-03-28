-include .env

VERSION := $(shell git describe --tags 2>/dev/null || echo "dev")
BUILD := $(shell git rev-parse --short HEAD 2>/dev/null || echo "dev")
PROJECT_NAME := modalsdb-graph-module
BINARY := ./bin/$(PROJECT_NAME)

LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"
PID_FILE := /tmp/.$(PROJECT_NAME).pid

.PHONY: help build run dev test clean stop

help: Makefile
# TODO: Implement help target later


clean:
	@echo "$(PROJECT_NAME) Cleaning..."
	@rm -rf bin/
	@rm -f $(PID_FILE)
	@go clean
build: clean
	@echo "$(PROJECT_NAME) Building binary to $(BINARY)..."
	go build $(LDFLAGS) -o $(BINARY) ./main.go

run: build
	@echo "$(PROJECT_NAME) Running $(PROJECT_NAME)"
	$(BINARY)

dev:
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "$(PROJECT_NAME) 'air' not found. Install it for hot-reload"; \
		$(MAKE) run; \
	fi

test:
	@echo "$(PROJECT_NAME) Running tests..."
	go test -v -race ./...

stop: 
	@if [ -f $(PID_FILE)]; then \
		kill $$(cat $(PID_FILE)) || true; \
		rm $(PID_FILE); \
		echo "$(PROJECT_NAME) Stopped."; \
	else \
		echo "$(PROJECT_NAME) No PID file found."; \
	fi