MIN_MAKE_VERSION := 3.81

# Min version
ifneq ($(MIN_MAKE_VERSION),$(firstword $(sort $(MAKE_VERSION) $(MIN_MAKE_VERSION))))
	$(error GNU Make $(MIN_MAKE_VERSION) or higher required)
endif

.DEFAULT_GOAL:=help

##@ Development
.PHONY: up down run-cli rebuild
run: ## Start the daemon & infrastructure containers (use "args=" to supply custom arguments)
	docker-compose up $(args)

stop: ## Stop and remove the application containers
	docker-compose down --volumes

run-cli: ## Get into daemon container
	docker-compose run --rm app bash

rebuild: ## Rebuild Docker Containers
	docker-compose down --volumes && \
	docker-compose rm --all && \
	docker-compose pull && \
	docker-compose build --no-cache && \
	docker-compose up --force-recreate

##@ Miscellaneous
.PHONY: run-bare run-infrastructure
run-bare: ## Start only daemon, assuming youre infrastructure is running already (use "args=" to supply custom arguments)
	docker-compose up app $(args)

run-infrastructure: ## Start infrastructure (good for local env)
	docker-compose up -d elastic nsqlookupd nsqd nsqadmin mongo $(args)

##@ Help
.PHONY: help
help: ## Show all available commands (you are looking at it)
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
