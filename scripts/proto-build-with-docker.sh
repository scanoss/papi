#!/bin/bash
#
# Protocol Buffer Docker Build Script
# 
# Builds and runs a Docker container for generating protocol buffer code.
# This script can be run from anywhere - it automatically finds the repository root.
#

# Determine repository root directory from script location
# Works regardless of where the script is executed from
script_path=$(realpath "$0")
script_dir=$(dirname "$script_path")
repo_dir=$(dirname "$script_dir")
if [ "$repo_dir" = "" ]; then
  echo "Error: Cannot determine repository root directory" >&2
  exit 1
fi

# Build Docker image from repository root
# --platform linux/amd64: Forces x86_64 architecture (required for protoc binaries)
# --tag proto-build: Tags the image for later use
# --file Containerfile: Specifies the Dockerfile to use
# "$repo_dir": Build context (repository root contains scripts/ and Containerfile)
docker buildx build --platform linux/amd64 --tag proto-build --file Containerfile "$repo_dir"

# Configure user mapping for file ownership
# Rootless Docker: Files automatically have correct ownership
# Regular Docker: Map host user to container user to avoid root-owned files
USER_FLAG=""
if ! docker info 2>/dev/null | grep -q "rootless"; then
    USER_FLAG="--user $(id -u):$(id -g)"
fi

# Run the container to generate protocol buffer code
# --rm: Remove container after execution
# -v "$repo_dir:/workspace": Mount repository root as /workspace in container
# proto-build: Use the image we just built
# "$@": Pass all script arguments to the container
docker run --rm -v "$repo_dir:/workspace" $USER_FLAG proto-build "$@"