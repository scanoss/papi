#!/bin/bash
#
# Helper functions for installing protobuf tools
#

# Download from GitHub releases using direct URLs
download_github_release() {
    local repo="$1"        # e.g. "protocolbuffers/protobuf"
    local tag="$2"         # e.g. "v31.1"
    local asset_name="$3"  # e.g. "protoc-31.1-linux-x86_64.zip"
    local output_path="$4" # e.g. "/tmp/protoc.zip"
    
    echo "Downloading ${asset_name} from ${repo}..."
    curl -L -o "$output_path" "https://github.com/${repo}/releases/download/${tag}/${asset_name}"
}

# Install binary from tar.gz archive
install_tar_gz() {
    local archive_path="$1"
    local binary_name="$2"
    local target_name="${3:-$binary_name}"
    
    echo "Installing ${target_name} from ${archive_path}..."
    tar -C /tmp -xzf "$archive_path"
    cp "/tmp/$binary_name" "/usr/local/bin/$target_name"
    chmod +x "/usr/local/bin/$target_name"
}

# Install binary from zip archive (extracts to /usr/local directly)
install_zip() {
    local archive_path="$1"
    local target_name="$2"
    
    echo "Installing ${target_name} from ${archive_path}..."
    unzip "$archive_path" -d /usr/local
    chmod +x "/usr/local/bin/$target_name"
}

# Install direct binary file
install_binary() {
    local source_path="$1"
    local target_name="$2"
    
    echo "Installing ${target_name} from ${source_path}..."
    cp "$source_path" "/usr/local/bin/$target_name"
    chmod +x "/usr/local/bin/$target_name"
}