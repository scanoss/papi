#!/bin/bash

# Build Docker image
docker buildx build --platform linux/amd64 --tag scanoss-protoc --file Containerfile.protoc .

# For rootless Docker installations: don't use --user flag, files will have correct ownership automatically
USER_FLAG=""
if ! docker info 2>/dev/null | grep -q "rootless"; then
    # Regular Docker: use --user flag to map host user to container user
    USER_FLAG="--user $(id -u):$(id -g)"
fi

# Run container with language and optional destination
docker run --rm -v "$(pwd):/workspace" $USER_FLAG scanoss-protoc "$LANG" "$DEST_DIR"