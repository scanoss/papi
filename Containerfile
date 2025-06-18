# All binaries are downloaded for x86_64 architecture
# Docker/QEMU handles cross-platform emulation when running on ARM (e.g., Apple Silicon)
FROM alpine:3.22

# IMPORTANT: Python gRPC Plugin Limitation
# The official protoc gRPC Python plugin (grpcio-tools) is implemented as a Python
# C extension module, NOT as a standalone binary like Go or Node.js plugins.
#
# This means:
# - You CANNOT extract a standalone 'protoc-gen-python' binary
# - You MUST use 'python3 -m grpc_tools.protoc' instead of 'protoc --plugin=protoc-gen-python'
# - Python runtime is required for Python protobuf generation
ARG PYTHON_VERSION=~3.12
ARG PIP_VERSION=~25.1

# Tool versions are defined with defaults in download-tools.sh
# Override versions at build time with: docker build --build-arg PROTOC_VERSION=32.0

# Install Python runtime required for gRPC Python plugin
RUN apk add --no-cache python3=${PYTHON_VERSION} py3-pip=${PIP_VERSION}

# Install utilities needed for downloading and extracting tools
# gcompat provides glibc compatibility for protoc-gen-js binary
RUN apk add --no-cache wget unzip bash curl gcompat

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