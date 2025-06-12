#!/bin/bash
set -e

LANG=$1
DEST_DIR=$2

if [ -z "$LANG" ]; then
    echo "Usage: $0 <language> [destination]"
    echo "Languages: go, python, js"
    exit 1
fi

echo "Building $LANG library using Docker..."

# Build Docker image
docker build -t scanoss-protoc -f Containerfile.protoc .

# Determine user flag for Docker permissions
USER_FLAG=""
if ! docker info 2>/dev/null | grep -q "rootless"; then
    # Regular Docker: use --user flag to map host user to container user
    USER_FLAG="--user $(id -u):$(id -g)"
fi
# Rootless Docker: don't use --user flag, files will have correct ownership automatically

# Run container with language and optional destination
docker run --rm -v "$(PWD):/workspace" $USER_FLAG scanoss-protoc "$LANG" "$DEST_DIR"