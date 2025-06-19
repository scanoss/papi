#!/bin/bash
#
#  SPDX-License-Identifier: MIT
#
#    Copyright (c) 2021, SCANOSS
#
#    Permission is hereby granted, free of charge, to any person obtaining a copy
#    of this software and associated documentation files (the "Software"), to deal
#    in the Software without restriction, including without limitation the rights
#    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
#    copies of the Software, and to permit persons to whom the Software is
#    furnished to do so, subject to the following conditions:
#
#    The above copyright notice and this permission notice shall be included in
#    all copies or substantial portions of the Software.
#
#    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
#    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
#    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
#    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
#    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
#    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
#    THE SOFTWARE.

# Download protobuf dependencies required for gRPC-Gateway
# Downloads googleapis and protoc-gen-openapiv2 proto files from GitHub repositories
# More info: https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file#2-with-custom-annotations

# All dependencies are pinned to specific commits to ensure:
# - Reproducible builds across different environments
# - Compatibility between protobuf dependencies and generated code

# Determine repository root directory from script location
script_path=$(realpath "$0")
script_dir=$(dirname "$script_path")
repo_dir=$(dirname "$script_dir")
if [ "$repo_dir" = "" ]; then
  echo "Error: Cannot determine repository root directory" >&2
  exit 1
fi

# googleapis commit
GOOGLEAPIS_COMMIT=${GOOGLEAPIS_COMMIT:-818eab111d50df57f4b09b7044f72c039c556d42}
# gRPC-Gateway v2 version
GRPC_GATEWAY_COMMIT=${GRPC_GATEWAY_COMMIT:-v2.24.0}


# Google API files required for HTTP annotations
GOOGLEAPIS_FILES=(
    "google/api/annotations.proto"
    "google/api/field_behavior.proto"
    "google/api/http.proto"
    "google/api/httpbody.proto"
)

# protoc-gen-openapiv2 files required for OpenAPI/Swagger generation (gRPC-Gateway v2)
OPENAPIV2_FILES=(
    "protoc-gen-openapiv2/options/annotations.proto"
    "protoc-gen-openapiv2/options/openapiv2.proto"
)

echo "=== Downloading protobuf dependencies ==="

PROTOBUF_DIR="$repo_dir/protobuf"

# Download googleapis files
echo "Downloading googleapis files from commit $GOOGLEAPIS_COMMIT"
for file in "${GOOGLEAPIS_FILES[@]}"; do
    url="https://raw.githubusercontent.com/googleapis/googleapis/$GOOGLEAPIS_COMMIT/$file"

    echo "  Downloading $file"
      if ! (cd "$PROTOBUF_DIR" && wget --force-directories --no-host-directories --cut-dirs=3 --quiet "$url"); then
          echo "Error: Failed to download $file"
          exit 1
      fi
done

# Download protoc-gen-openapiv2 files
echo "Downloading protoc-gen-openapiv2 files from $GRPC_GATEWAY_COMMIT"
for file in "${OPENAPIV2_FILES[@]}"; do
    url="https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/$GRPC_GATEWAY_COMMIT/$file"
    target="$OPENAPIV2_DIR/${file#protoc-gen-openapiv2/}"
    
    echo "  Downloading $file"
      if ! (cd "$PROTOBUF_DIR" && wget --force-directories --no-host-directories --cut-dirs=3 --quiet "$url"); then
          echo "Error: Failed to download $file"
          exit 1
      fi
done

echo "Successfully downloaded all protobuf dependencies"
echo "googleapis: ${#GOOGLEAPIS_FILES[@]} files (commit: $GOOGLEAPIS_COMMIT)"
echo "grpc-gateway v2: ${#OPENAPIV2_FILES[@]} files (commit: $GRPC_GATEWAY_COMMIT)"