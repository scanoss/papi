#!/bin/bash
#
# Download and install all protobuf tools
#

# Source helper functions
source /usr/local/bin/install-helpers.sh

# All tool are pinned to specific versions to ensure:
# - Reproducible builds across different environments
# - Compatibility between protobuf tools and generated code

# Core protobuf compiler
PROTOC_VERSION=${PROTOC_VERSION:-31.1}

# JavaScript/Node.js tools
JS_PROTOBUF_VERSION=${JS_PROTOBUF_VERSION:-3.21.4}
JS_GRPC_TOOLS_VERSION=${JS_GRPC_TOOLS_VERSION:-1.13.0}

# Go tools
GO_PROTOC_GEN_VERSION=${GO_PROTOC_GEN_VERSION:-1.36.6}
GO_PROTOC_GEN_GRPC_VERSION=${GO_PROTOC_GEN_GRPC_VERSION:-1.5.1}
GO_GRPC_GATEWAY_VERSION=${GO_GRPC_GATEWAY_VERSION:-1.16.0}

# Python tools
PY_GRPCIO_TOOLS_VERSION=${PY_GRPCIO_TOOLS_VERSION:-1.73.0}
PY_GRPCIO_VERSION=${PY_GRPCIO_VERSION:-1.73.0}

echo "=== Installing protobuf tools ==="

# Install protoc
echo "Installing protoc ${PROTOC_VERSION}"
download_github_release "protocolbuffers/protobuf" "v${PROTOC_VERSION}" "protoc-${PROTOC_VERSION}-linux-x86_64.zip" "/tmp/protoc.zip"
install_zip "/tmp/protoc.zip" "bin/protoc"
install_includes "/tmp/protoc.zip" "include" "/usr/local"
rm /tmp/protoc.zip

# Install protobuf-javascript (protoc-gen-js)
echo "Installing protobuf-javascript ${JS_PROTOBUF_VERSION}"
download_github_release "protocolbuffers/protobuf-javascript" "v${JS_PROTOBUF_VERSION}" "protobuf-javascript-${JS_PROTOBUF_VERSION}-linux-x86_64.tar.gz" "/tmp/protobuf-js.tar.gz"
install_tar_gz "/tmp/protobuf-js.tar.gz" "bin/protoc-gen-js"
rm /tmp/protobuf-js.tar.gz

# Install grpc-tools (from gRPC CDN)
echo "Installing grpc-tools ${JS_GRPC_TOOLS_VERSION}"
curl -L -o /tmp/grpc-tools.tar.gz "https://node-precompiled-binaries.grpc.io/grpc-tools/v${JS_GRPC_TOOLS_VERSION}/linux-x64.tar.gz"
install_tar_gz "/tmp/grpc-tools.tar.gz" "bin/grpc_node_plugin"
rm /tmp/grpc-tools.tar.gz

# Install protoc-gen-go
echo "Installing protoc-gen-go ${GO_PROTOC_GEN_VERSION}"
download_github_release "protocolbuffers/protobuf-go" "v${GO_PROTOC_GEN_VERSION}" "protoc-gen-go.v${GO_PROTOC_GEN_VERSION}.linux.amd64.tar.gz" "/tmp/protoc-gen-go.tar.gz"
install_tar_gz "/tmp/protoc-gen-go.tar.gz" "protoc-gen-go"
rm /tmp/protoc-gen-go.tar.gz

# Install protoc-gen-go-grpc (monorepo structure)
echo "Installing protoc-gen-go-grpc ${GO_PROTOC_GEN_GRPC_VERSION}"
download_github_release "grpc/grpc-go" "cmd/protoc-gen-go-grpc/v${GO_PROTOC_GEN_GRPC_VERSION}" "protoc-gen-go-grpc.v${GO_PROTOC_GEN_GRPC_VERSION}.linux.amd64.tar.gz" "/tmp/protoc-gen-go-grpc.tar.gz"
install_tar_gz "/tmp/protoc-gen-go-grpc.tar.gz" "protoc-gen-go-grpc"
rm /tmp/protoc-gen-go-grpc.tar.gz

# Install grpc-gateway tools (direct binaries)
echo "Installing grpc-gateway tools ${GO_GRPC_GATEWAY_VERSION}"
download_github_release "grpc-ecosystem/grpc-gateway" "v${GO_GRPC_GATEWAY_VERSION}" "protoc-gen-grpc-gateway-v${GO_GRPC_GATEWAY_VERSION}-linux-x86_64" "/tmp/protoc-gen-grpc-gateway"
download_github_release "grpc-ecosystem/grpc-gateway" "v${GO_GRPC_GATEWAY_VERSION}" "protoc-gen-swagger-v${GO_GRPC_GATEWAY_VERSION}-linux-x86_64" "/tmp/protoc-gen-swagger"
install_binary "/tmp/protoc-gen-grpc-gateway"
install_binary "/tmp/protoc-gen-swagger"
rm /tmp/protoc-gen-grpc-gateway /tmp/protoc-gen-swagger

# Install Python gRPC tools
# --break-system-packages flag bypasses pip's environment isolation in containerized environments
echo "Installing Python gRPC tools ${PY_GRPCIO_TOOLS_VERSION}"
pip install --no-cache-dir --break-system-packages grpcio-tools==${PY_GRPCIO_TOOLS_VERSION} grpcio==${PY_GRPCIO_VERSION}

echo "=== All tools installed successfully ==="

# Verify installations
echo "=== Verifying installations ==="
protoc --version
protoc-gen-go --version
protoc-gen-go-grpc --version
which protoc-gen-grpc-gateway
which protoc-gen-swagger
which protoc-gen-js
which grpc_node_plugin
python3 -c "import grpc_tools.protoc; print('Python gRPC tools available')"
echo "All tools are ready"