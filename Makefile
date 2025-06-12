

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

build_go_native:  ## Build the Go library from protos using native tools
	bash protobuf/build.sh --native -f -t go -d .

build_go:  ## Build the Go library from protos using Docker
	bash protobuf/build.sh --docker -f -t go -d .

build_python_native:  ## Build the Python library from protos using native tools
	bash protobuf/build.sh --native -f -t python

build_python:  ## Build the Python library from protos using Docker
	bash protobuf/build.sh --docker -f -t python

build_js_native:  ## Build the Javascript library using native tools
	bash protobuf/build.sh --native -f -t js

build_js:  ## Build the Javascript library using Docker
	bash protobuf/build.sh --docker -f -t js

build_all:  ## Build all languages using Docker
	bash protobuf/build.sh --docker -f -t go -d .
	bash protobuf/build.sh --docker -f -t python
	bash protobuf/build.sh --docker -f -t js

update_scanoss_py:  ## Copy the latest Python API code to scanoss.py
	@echo "Copying Python API to scanoss.py..."
	bash protobuf/copy_python.sh -f -s python -t ../../scanoss.py/src

python_all: clean build_python update_scanoss_py ## Execute all Python actions

