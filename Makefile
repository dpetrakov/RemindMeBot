# Makefile for RemindMeBot

# Environment variables
ENV_FILE := .env

# Load environment variables
include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

# Go variables
GO ?= go
GOFLAGS ?= $(GOFLAGS:)

# Docker variables
DOCKER_COMPOSE ?= docker-compose

# Targets
.PHONY: help build clean run test

help:
	@echo "Usage:"
	@echo "  make build    - build the application"
	@echo "  make clean    - remove build artifacts"
	@echo "  make run      - run the application"
	@echo "  make test     - run tests"

build:
	$(DOCKER_COMPOSE) build

clean:
	rm -rf bin/

run:
	$(DOCKER_COMPOSE) up

test:
	$(GO) test ./...


