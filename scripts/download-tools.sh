#!/bin/bash
#
# Download and install all protobuf tools
#

# Source helper functions
source /usr/local/bin/install-helpers.sh

# Get versions from environment variables (set in Containerfile)
PROTOC_VERSION=${PROTOC_VERSION:-31.1}
PROTOBUF_JS_VERSION=${PROTOBUF_JS_VERSION:-3.21.4}
PROTOC_GEN_GO_VERSION=${PROTOC_GEN_GO_VERSION:-1.36.6}
PROTOC_GEN_GO_GRPC_VERSION=${PROTOC_GEN_GO_GRPC_VERSION:-1.5.1}
GRPC_GATEWAY_VERSION=${GRPC_GATEWAY_VERSION:-2.26.3}
GRPC_TOOLS_VERSION=${GRPC_TOOLS_VERSION:-1.13.0}

echo "=== Installing protobuf tools ==="

# Install protoc
echo "Installing protoc ${PROTOC_VERSION}"
download_github_release "protocolbuffers/protobuf" "v${PROTOC_VERSION}" "protoc-${PROTOC_VERSION}-linux-x86_64.zip" "/tmp/protoc.zip"
install_zip "/tmp/protoc.zip" "protoc"
rm /tmp/protoc.zip

# Install protobuf-javascript (protoc-gen-js)
echo "Installing protobuf-javascript ${PROTOBUF_JS_VERSION}"
download_github_release "protocolbuffers/protobuf-javascript" "v${PROTOBUF_JS_VERSION}" "protobuf-javascript-${PROTOBUF_JS_VERSION}-linux-x86_64.tar.gz" "/tmp/protobuf-js.tar.gz"
tar -C /tmp -xzf /tmp/protobuf-js.tar.gz
install_binary "/tmp/bin/protoc-gen-js" "protoc-gen-js"
rm -rf /tmp/protobuf-js.tar.gz /tmp/bin

# Install protoc-gen-go
echo "Installing protoc-gen-go ${PROTOC_GEN_GO_VERSION}"
download_github_release "protocolbuffers/protobuf-go" "v${PROTOC_GEN_GO_VERSION}" "protoc-gen-go.v${PROTOC_GEN_GO_VERSION}.linux.amd64.tar.gz" "/tmp/protoc-gen-go.tar.gz"
install_tar_gz "/tmp/protoc-gen-go.tar.gz" "protoc-gen-go" "protoc-gen-go"
rm -rf /tmp/protoc-gen-go*

# Install protoc-gen-go-grpc (monorepo structure)
echo "Installing protoc-gen-go-grpc ${PROTOC_GEN_GO_GRPC_VERSION}"
download_github_release "grpc/grpc-go" "cmd/protoc-gen-go-grpc/v${PROTOC_GEN_GO_GRPC_VERSION}" "protoc-gen-go-grpc.v${PROTOC_GEN_GO_GRPC_VERSION}.linux.amd64.tar.gz" "/tmp/protoc-gen-go-grpc.tar.gz"
install_tar_gz "/tmp/protoc-gen-go-grpc.tar.gz" "protoc-gen-go-grpc" "protoc-gen-go-grpc"
rm -rf /tmp/protoc-gen-go-grpc*

# Install grpc-gateway tools (direct binaries)
echo "Installing grpc-gateway tools ${GRPC_GATEWAY_VERSION}"
download_github_release "grpc-ecosystem/grpc-gateway" "v${GRPC_GATEWAY_VERSION}" "protoc-gen-grpc-gateway-v${GRPC_GATEWAY_VERSION}-linux-x86_64" "/tmp/protoc-gen-grpc-gateway"
download_github_release "grpc-ecosystem/grpc-gateway" "v${GRPC_GATEWAY_VERSION}" "protoc-gen-openapiv2-v${GRPC_GATEWAY_VERSION}-linux-x86_64" "/tmp/protoc-gen-openapiv2"
install_binary "/tmp/protoc-gen-grpc-gateway" "protoc-gen-grpc-gateway"
install_binary "/tmp/protoc-gen-openapiv2" "protoc-gen-swagger"
rm -rf /tmp/protoc-gen-*

# Install grpc-tools (from npm CDN)
echo "Installing grpc-tools ${GRPC_TOOLS_VERSION}"
curl -L -o /tmp/grpc-tools.tar.gz "https://node-precompiled-binaries.grpc.io/grpc-tools/v${GRPC_TOOLS_VERSION}/linux-x64.tar.gz"
tar -C /tmp -xzf /tmp/grpc-tools.tar.gz
cp /tmp/bin/* /usr/local/bin/
chmod +x /usr/local/bin/grpc_node_plugin
rm -rf /tmp/grpc-tools.tar.gz /tmp/bin

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
echo "All tools are ready"