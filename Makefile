

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

clean:  ## Clean all dev data
	@echo "Removing dev data..."
	@rm -rf python

build_go:  ## Build the Go library from protos
	@echo "Building Go API libraries..."
	bash protobuf/build.sh -f -t go -d .

build_python:  ## Build the Python library from protos
	@echo "Building Python API libraries..."
	bash protobuf/build.sh -f -t python
