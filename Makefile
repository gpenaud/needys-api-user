## permanent variables
PROJECT			?= github.com/gpenaud/needys-api-user
RELEASE			?= $(shell git describe --tags --abbrev=0)
COMMIT			?= $(shell git rev-parse --short HEAD)
BUILD_TIME  ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')

## docker environment options
DOCKER_BUILD_ARGS ?= --build-arg PROJECT="${PROJECT}" --build-arg RELEASE="${RELEASE}" --build-arg COMMIT="${COMMIT}" --build-arg BUILD_TIME="${BUILD_TIME}"

## docker-compose options
DOCKER_COMPOSE_OPTIONS ?= --file deployments/docker-compose.yml --file deployments/development-override.yml

## Colors
COLOR_RESET       = $(shell tput sgr0)
COLOR_ERROR       = $(shell tput setaf 1)
COLOR_COMMENT     = $(shell tput setaf 3)
COLOR_TITLE_BLOCK = $(shell tput setab 4)

## display this help text
help:
	@printf "\n"
	@printf "${COLOR_TITLE_BLOCK}${PROJECT} Makefile${COLOR_RESET}\n"
	@printf "\n"
	@printf "${COLOR_COMMENT}Usage:${COLOR_RESET}\n"
	@printf " make build\n\n"
	@printf "${COLOR_COMMENT}Available targets:${COLOR_RESET}\n"
	@awk '/^[a-zA-Z\-_0-9@]+:/ { \
				helpLine = match(lastLine, /^## (.*)/); \
				helpCommand = substr($$1, 0, index($$1, ":")); \
				helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
				printf " ${COLOR_INFO}%-15s${COLOR_RESET} %s\n", helpCommand, helpMessage; \
		} \
		{ lastLine = $$0 }' $(MAKEFILE_LIST)
	@printf "\n"

## stack - start the entire stack in background, then follow logs type=app for only application, type=service for only service
start:
ifeq ($(type),application)
	docker-compose ${DOCKER_COMPOSE_OPTIONS} up --build needys-api-user
else ifeq ($(type),service)
	docker-compose ${DOCKER_COMPOSE_OPTIONS} up --build --detach mariadb rabbitmq
else
	docker-compose ${DOCKER_COMPOSE_OPTIONS} up --build --detach
	docker-compose ${DOCKER_COMPOSE_OPTIONS} logs --follow needys-api-user
endif

## stack - rebuild and restart needys-api-user
rebuild:
	docker-compose ${DOCKER_COMPOSE_OPTIONS} stop needys-api-user
	docker-compose ${DOCKER_COMPOSE_OPTIONS} up --build needys-api-user

## stack - stop the entire stack
stop:
	docker-compose ${DOCKER_COMPOSE_OPTIONS} down

## stack - watch the stack
watch:
	watch docker-compose ${DOCKER_COMPOSE_OPTIONS} ps

## stack - log the entire stack
logs:
	docker-compose ${DOCKER_COMPOSE_OPTIONS} logs --follow

## test - execute all api queries to check results on bash
.PHONY: test
test:
	/bin/sh scripts/test-api.sh --query

## test - execute all unit-tests defined in application
test-unit:
	@echo "Stricts unit-tests are not yet implemented !"

## test - execute all cucumber-behavior tests defined in application
test-behavior:
	go test -v ./... --godog.format=pretty --godog.random -race -covermode=atomic

test-stack:
	docker-compose --file deployments/docker-compose.yml up --build --detach

## docker - build the needys-api-user image
.PHONY: build
build:
	docker build ${DOCKER_BUILD_ARGS} --file build/package/Dockerfile --tag needys-api-user:latest .

## docker - enter into the needys-api-user container
enter:
	docker-compose ${DOCKER_COMPOSE_OPTIONS} exec needys-api-user /bin/sh

cleanup:
	sudo rm -rf needys-api-user tmp
