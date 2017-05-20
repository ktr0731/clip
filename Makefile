SHELL := /bin/bash
VERSION := 1.2.0
REVISION := $(shell git rev-parse --short HEAD)

# Show version
version:
	@echo "Version: $(VERSION)($(REVISION))"

glide:
ifeq ($(shell which glide 2>/dev/null),)
	mkdir -p $(GOPATH)/bin
	curl -s https://glide.sh/get | sh
endif

deps: glide
ifeq ($(shell find . -depth 1 -name vendor 2>/dev/null),)
	glide install
endif

test: deps
	go test -v

build: deps
	go build -o clip
