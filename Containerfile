# Base Image: python:3.12-alpine
#
# Why python:3.12-alpine instead of plain alpine:latest?
# - Python gRPC plugin (grpcio-tools) is a C extension, NOT a standalone binary
# - Python runtime is mandatory for protobuf Python generation
# - Cannot extract 'protoc-gen-python' binary - must use 'python3 -m grpc_tools.protoc'
# - python:3.12-alpine provides optimized setup while keeping image small
#
# All binaries downloaded for x86_64 (Docker/QEMU handles ARM emulation)
# Tool versions configurable via build args: docker build --build-arg PROTOC_VERSION=32.0
FROM python:3.9-alpine

# Install utilities needed for downloading and extracting tools
# gcompat, libstdc++, libgcc provide glibc compatibility for protoc-gen-js binary
# g++, gcc, musl-dev are required for building grpcio-tools from source
RUN apk add --no-cache wget unzip bash curl gcompat libstdc++ libgcc

# Copy installation scripts into the container
COPY scripts/install-helpers.sh /usr/local/bin/
COPY scripts/download-tools.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/*.sh

# Install all protocol buffer tools and plugins
RUN /usr/local/bin/download-tools.sh

# Set working directory to /workspace where user's project will be mounted
# This matches the volume mount in proto-build-with-docker.sh: -v "$(pwd):/workspace"
WORKDIR /workspace

# Entry point runs the proto-build.sh script from the user's project
# Script must exist at ./scripts/proto-build.sh relative to the mounted workspace
ENTRYPOINT ["./scripts/proto-build.sh"]