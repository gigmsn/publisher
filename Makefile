.DEFAULT_GOAL := help

PROJECT_NAME := GIGMSN Publisher

.PHONY: help
help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run tests
	@docker-compose up --build test
	@docker-compose rm -fsv test

.PHONY: server/up
server/up: ## run server container
	@docker-compose up --build -d server

.PHONY: server/stop
server/stop: ## stop and remove server container
	@docker-compose rm -fsv server

.PHONY: client/up
jsclient/up: ## run jsclient container
	@docker-compose up --build jsclient

.PHONY: client/stop
jsclient/stop: ## stop and remove jsclient container
	@docker-compose rm -fsv jsclient

.PHONY: broker/up
broker/up: ## start broker container
	@docker-compose up -d broker

.PHONY: broker/stop
broker/stop: ## stop and remove broker container
	@docker-compose rm -fsv broker
