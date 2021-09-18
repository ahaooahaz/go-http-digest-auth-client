VERSION=v1.0.4
PROJECT=$(shell basename $(shell pwd))
DOCKER_COMPOSE=docker-compose

.PHONY: all
all: deps test

deps:
	$(DOCKER_COMPOSE) -f tests/docker-compose.yml up -d

test:
	go test ./... -coverprofile=coverage.txt -covermode=atomic