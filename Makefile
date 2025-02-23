DOCKER_HUB:=ghcr.io/babbage88/gocolors:
BIN_NAME:=gocolors
INSTALL_PATH:=$${GOPATH}/bin
ENV_FILE:=.env
MIG:=$(shell date '+%m%d%Y.%H%M%S')
SHELL := /bin/bash
VERBOSE ?= 1
ifeq ($(VERBOSE),1)
	V = -v
endif

build:
	go build $(V) -o $(BIN_NAME) .

build-quiet:
	go build -o $(BIN_NAME)

install: build
	echo "Install Path $(INSTALL_PATH)"
	mv $(BIN_NAME) $(INSTALL_PATH)

check-builder:
	@if ! docker buildx inspect gocolorsbuilder > /dev/null 2>&1; then \
		echo "Builder gocolorsbuilder does not exist. Creating..."; \
		docker buildx create --name gocolorsbuilder --bootstrap; \
	fi

create-builder: check-builder

buildandpush: check-builder
	docker buildx use $(BIN_NAME)builder
	docker buildx build --platform linux/amd64,linux/arm64 -t $(DOCKER_HUB)$(tag) . --push

