# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

clean:  ## Clean all dev data
	@echo "Removing dev data..."
	@rm -rf python javascript

build_go:  ## Build the Go library from protos
	@echo "Building Go API libraries..."
	@scripts/proto-build-with-docker.sh -f -t go -d .

build_python:  ## Build the Python library from protos
	@echo "Building Python API libraries..."
	@scripts/proto-build-with-docker.sh -f -t python

build_js:  ## Build the Javascript library with Typescript definitions from protos
	@echo "Building Javascript API libraries..."
	@scripts/proto-build-with-docker.sh -f -t js

update_scanoss_py:  ## Copy the latest Python API code to scanoss.py
	@echo "Copying Python API to scanoss.py..."
	@scripts/copy_python.sh -f -s python -t ../../scanoss.py/src

update_scanoss_js:  ## Copy the latest JavaScript API code to scanoss.js
	@echo "Copying JavaScript API to scanoss.js..."
	@scripts/copy_js.sh -f -s javascript -t ../../scanoss.js/src/sdk/Services/Grpc
