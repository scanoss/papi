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
# - JS

b_dir=$(dirname "$0") # script location
if [ "$b_dir" = "" ]; then
  b_dir=.
fi
export b_dir
export protobuf_dir=protobuf
#
# Print help info
#
help() {
  echo "Usage: proto-build [-h]
               [-d output folder]
               -t <python|go|js>"
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
# Build the Javascript library code and Typescript files for the proto definitions.
# See reference https://github.com/improbable-eng/ts-protoc-gen
#
# TODO: For TS support install grpc-web, read here: https://github.com/protobufjs/protobuf.js/issues/1327
#       or switch to protobufjs directly that has TS support
build_js() {
  dest_dir=$1
  echo "Writing Javascript APIs to: $dest_dir"
  if [ ! -d "$dest_dir" ]; then
    echo "Error: Destination directory does not exist: $dest_dir"
    exit 1
  fi
  sc_dir="$dest_dir/scanoss"
  if [ -d "$sc_dir" ]; then
    rm -rf "$sc_dir"
  fi
  if ! protoc -I$protobuf_dir \
              --plugin=protoc-gen-grpc="$(which grpc_node_plugin)" \
              --js_out=import_style=commonjs,binary:"$dest_dir" \
              --grpc_out=grpc_js:"$dest_dir" \
              $(find $protobuf_dir/scanoss -type f -name "scanoss*.proto" -print);
              then
    echo "Error: Failed to compile Javascript libraries from proto files"
    exit 1
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
  if ! python3 -m grpc_tools.protoc -I$protobuf_dir --python_out="$dest_dir" --grpc_python_out="$dest_dir" \
               $(find $protobuf_dir/scanoss -type f -name "scanoss*.proto" -print) \
               $(find $protobuf_dir/protoc-gen-swagger -type f -name "*.proto" -print)
               then
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
  cd "$dest_dir" || {
    echo "Error: failed to cd to $dest_dir"
    exit 1
  }
  target_dir=api
  rm -rf "$target_dir"
  mkdir "$target_dir" || {
    echo "Error: Failed to create $target_dir"
    exit 1
  }
  if ! protoc --proto_path=$protobuf_dir --go_out="$target_dir" --go-grpc_out="$target_dir" \
              --grpc-gateway_out="$target_dir" --grpc-gateway_opt logtostderr=true \
              --grpc-gateway_opt generate_unbound_methods=true \
              $(find $protobuf_dir/scanoss -type f -name "scanoss*.proto" -print)
              then
    echo "Error: Failed to compile Go libraries from proto files."
    exit 1
  fi
  if ! mv "$target_dir/github.com/scanoss/papi/api/"* "$target_dir/"; then
    echo "Error: Failed to move go library code from generated folder to $target_dir"
    exit 1
  fi
  rm -rf "$target_dir/github.com"
  echo "Producing Swagger files..."
  find $protobuf_dir/scanoss -type f -name "scanoss*.proto" -exec \
    protoc --proto_path=$protobuf_dir --swagger_out=logtostderr=true:$protobuf_dir "{}" \;
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
  echo "Please specify a language (python, go or js)."
  help
fi
if [ "$t" != "go" ] && [ "$t" != "python" ] && [ "$t" != "js" ]; then
  echo "Error: Unknown build type: $t"
  help
fi

if [ "$t" = "python" ]; then
  confirm $force "Create Python library from proto?"
  cd "$b_dir/.." || {
    echo "Error: Failed to cd to $b_dir"
    exit 1
  }
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
elif [ "$t" = "js" ]; then
  confirm $force "Create Javascript library from proto?"
  cd "$b_dir/.." || {
    echo "Error: Failed to cd to $b_dir"
    exit 1
  }
  dest="javascript"
  if [ -n "$d" ]; then
    dest="$d"
  else
    mkdir -p "$dest"
  fi
  build_js "$dest"
else
  echo "Error: Unknown language type: $t"
  help
fi

exit 0