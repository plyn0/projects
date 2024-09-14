# VERSION = $(shell cat pyproject.toml | grep version | head -1 | awk -F= "{ print $2 }" | sed -E 's/(version)|[:=,\",]//g' | sed 's/ *//g')
# PROJECT_NAME = $(shell cat pyproject.toml | grep name | head -1 | awk -F= "{ print $2 }" | sed -E 's/(name)|[:=,\",]//g' | sed 's/ *//g')
# BRANCH := $(shell git for-each-ref --format='%(objectname) %(refname:short)' refs/heads | awk "/^$$(git rev-parse HEAD)/ {print \$$2}")

# TAG = $(strip $(PROJECT_NAME)):$(strip $(VERSION))
# VERSION_NUM = $(strip $(VERSION))
# GITLAB_PATH = gitlab.com/synapse-medicine/microservices
# IMAGE_NAME = registry.$(strip $(GITLAB_PATH))/$(strip $(TAG))
# GITLAB_REG = https://$(strip $(GITLAB_PATH))/$(strip $(PROJECT_NAME))/container_registry

.PHONY: help
help: ## display makefile commands
	@ awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)''

.PHONY: test
test:  ## run all go tests
	@ go test ./...

.PHONY: testv
testv:  ## run all go tests
	@ go test ./... -v

# .PHONY: clean
# clean: ## Clean python cache (files and folders)
# 	@echo "Cleaning python cache"
# 	@echo 'find . -name "*.py[co]" -o -name __pycache__ -o -name htmlcov -exec rm -rf {} +'
# 	@find . -name "*.py[co]" -o -name __pycache__ -o -name htmlcov -exec rm -rf {} +

# .PHONY: tests
# tests: ## Launch test suite
# 	@python -m pytest tests/ -vv --cov=synapse_cache --cov-report term-missing

# .PHONY: chatops-cmd
# chatops-cmd: ## Prepare chatops command
# 	@echo "Prepare chatops command"
# 	@echo "/gitlab ci run tag $(strip ${PROJECT_NAME}) ${BRANCH} ${VERSION}"

# .PHONY: tag
# tag: ## Push project image to Gitlab registry
# 	@echo "Setting Gitlab tag: ${VERSION_NUM}"
# 	@git tag ${VERSION_NUM}
# 	@git push origin --tags

# .PHONY: redis
# redis: ## Run non-persistent Redis container with password "mypass"
# 	@docker run --rm --name rdb -p 6379:6379 redis --requirepass mypass

# .PHONY: start
# start: ## run the example app
# 	@uvicorn example.main:app --reload
