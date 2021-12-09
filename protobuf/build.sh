#!/usr/bin/env bash
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

# This file contains the commands required to produce API code for the following languages
# - Python
# - Go

export b_dir=$(dirname "$0") # script location
if [ "$b_dir" = "" ]; then
  export b_dir=.
fi
export protobuf_dir=protobuf
#
# Print help info
#
help() {
  echo "Usage: proto-build [-h]
               [-d folder]
               -t <python|go>"
  exit 2
}

#
# Confirm before proceeding
# Args:
#      1 - force: true/false
#      2 - question: text
#
confirm() {
  if [ "$1" != "true" ]; then
    read -r -p "$2 [Y/n] " response
    response=$(echo "$response" | tr '[:upper:]' '[:lower:]')
    if [[ $response =~ ^(yes|y| ) ]] || [[ -z $response ]]; then
      echo "Proceeding..."
    else
      echo "Aborting."
      exit 1
    fi
  fi
}

#
# Build the Python library code for the proto definitions.
#
build_python() {
  dest_dir=$1
  echo "Writing Python APIs to: $dest_dir"
  if [ ! -d "$dest_dir" ]; then
    echo "Error: Destination directory does not exist: $dest_dir"
    exit 1
  fi
  sc_dir="$dest_dir/scanoss"
  if [ -d "$sc_dir" ]; then
    rm -rf "$sc_dir"
  fi
  if ! python3 -m grpc_tools.protoc -I$protobuf_dir --python_out="$dest_dir" --grpc_python_out="$dest_dir" $(find $protobuf_dir -type f -name "*.proto" -print); then
    echo "Error: Failed to compile Python libraries from proto files"
    exit 1
  fi
}

#
# Build the Go library code for the proto definitions.
#
build_go() {
  dest_dir=$1
  echo "Writing Go APIs to: $dest_dir"
  if [ ! -d "$dest_dir" ]; then
    echo "Error: Destination directory does not exist: $dest_dir"
    exit 1
  fi
  cd "$dest_dir"
  target_dir=api
  rm -rf "$target_dir"
  mkdir "$target_dir" || {
    echo "Error: Failed to create $target_dir"
    exit 1
  }
  if ! protoc --proto_path=$protobuf_dir --go_out="$target_dir" --go-grpc_out="$target_dir" $(find $protobuf_dir -type f -name "*.proto" -print); then
    echo "Error: Failed to compile Go libraries from proto files."
    exit 1
  fi
  if ! mv "$target_dir/github.com/scanoss/papi/api/"* "$target_dir/"; then
    echo "Error: Failed to move go library code from generated folder to $target_dir"
    exit 1
  fi
  rm -rf "$target_dir/github.com"
}

#
# Parse command options and take action
#
force=false
while getopts ":d:t:,hf" o; do
  case "${o}" in
  t)
    t=${OPTARG}
    ;;
  d)
    d=${OPTARG}
    ;;
  h)
    help
    ;;
  f)
    force=true
    ;;
  *)
    echo "Unknown option: $o"
    help
    ;;
  esac
done
shift $((OPTIND - 1))

if [ -z "${t}" ]; then
  echo "Please specify a language (python or go)."
  help
fi
if [ "$t" != "go" ] && [ "$t" != "python" ]; then
  echo "Error: Unknown build type: $t"
  help
fi

if [ "$t" = "python" ]; then
  confirm $force "Create Python library from proto?"
  cd "$b_dir/.."
  dest="python"
  if [ ! -z "$d" ]; then
    dest="$d"
  else
    mkdir -p "$dest"
  fi
  build_python "$dest"
elif [ "$t" = "go" ]; then
  confirm $force "Create Go library from proto?"
  dest="$b_dir/.."
  if [ ! -z "$d" ]; then
    dest="$d"
  fi
  build_go "$dest"
else
  echo "Error: Unknown language type: $t"
  help
fi

exit 0
