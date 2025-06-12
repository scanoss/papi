

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

build_go_native:  ## Build the Go library from protos
	@echo "Building Go API libraries..."
	bash protobuf/build.sh -f -t go -d .

build_go:  ## Build the Go library from protos using Docker
	@echo "Building Go API libraries with Docker..."
	@echo "Cleaning existing API files..."
	rm -rf api
	docker build -t scanoss-protoc -f Containerfile.protoc .
	@USER_FLAG=$$(docker info 2>/dev/null | grep -q "rootless" || echo "--user $$(id -u):$$(id -g)"); \
	docker run --rm -v $$(pwd):/workspace $$USER_FLAG scanoss-protoc -f -t go -d .

build_python:  ## Build the Python library from protos
	@echo "Building Python API libraries..."
	bash protobuf/build.sh -f -t python

update_scanoss_py:  ## Copy the latest Python API code to scanoss.py
	@echo "Copying Python API to scanoss.py..."
	bash protobuf/copy_python.sh -f -s python -t ../../scanoss.py/src

python_all: clean build_python update_scanoss_py ## Execute all Python actions

build_js:  ## Build the Javascript library with Typescript definitions from protos
	@echo "Building Javascript API libraries..."
	bash protobuf/build.sh -f -t js

