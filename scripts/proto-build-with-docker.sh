#!/bin/bash

# Build Docker image.
# Force darwin users to generate a linux/amd64 image (with --platform linux/amd64)
# This relies on docker engine to emulate the os/arch
# The container downloads binaries on build time
docker buildx build --platform linux/amd64 --tag proto-build --file Containerfile .

# For rootless Docker installations: don't use --user flag, files will have correct ownership automatically
USER_FLAG=""
if ! docker info 2>/dev/null | grep -q "rootless"; then
    # Regular Docker: use --user flag to map host user to container user
    USER_FLAG="--user $(id -u):$(id -g)"
fi

docker run --rm -v "$(pwd):/workspace" $USER_FLAG proto-build "$@" # Pass all arguments