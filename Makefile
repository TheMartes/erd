MIN_MAKE_VERSION := 3.81

# Min version
ifneq ($(MIN_MAKE_VERSION),$(firstword $(sort $(MAKE_VERSION) $(MIN_MAKE_VERSION))))
	$(error GNU Make $(MIN_MAKE_VERSION) or higher required)
endif

.DEFAULT_GOAL:=help

##@ Development
.PHONY: up down build run-cli fresh-start
up: ## Start the application containers (use "args=" to supply custom arguments)
	docker-compose up $(args)

down: ## Stop and remove the application containers
	docker-compose down --volumes

build: ## Build docker container
	docker-compose up --build

run-cli: ## Get into container
	docker-compose run --rm app bash

fresh-start: ## ️🌿 Ceate new fresh containers
	docker-compose rm --all && \
	docker-compose pull && \
	docker-compose build --no-cache && \
	docker-compose up --force-recreate

##@ Help
.PHONY: help
help: ## Show all available commands (you are looking at it)
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
