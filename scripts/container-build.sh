#!/bin/bash
set -e

LANG=$1
DEST_DIR=$2

if [ -z "$LANG" ]; then
    echo "Usage: $0 <language> [destination]"
    echo "Languages: go, python, js"
    exit 1
fi

echo "Building $LANG protobuf APIs..."

case "$LANG" in
    go)
        DEST_DIR=${DEST_DIR:-"."}
        echo "Writing Go APIs to: $DEST_DIR"
        cd "$DEST_DIR" || exit 1
        
        # Clean and create api directory
        rm -rf api
        mkdir -p api
        
        # Generate Go code
        protoc --proto_path=protobuf \
               --go_out=api \
               --go-grpc_out=api \
               --grpc-gateway_out=api \
               --grpc-gateway_opt=logtostderr=true \
               --grpc-gateway_opt=generate_unbound_methods=true \
               $(find protobuf/scanoss -type f -name "*.proto")
        
        # Move generated files to correct location
        if [ -d "api/github.com/scanoss/papi/api" ]; then
            mv api/github.com/scanoss/papi/api/* api/
            rm -rf api/github.com
        fi
        
        # Generate Swagger files
        echo "Generating Swagger documentation..."
        find protobuf/scanoss -type f -name "*.proto" -exec \
            protoc --proto_path=protobuf --swagger_out=logtostderr=true:protobuf {} \;
        ;;
        
    python)
        DEST_DIR=${DEST_DIR:-"python"}
        echo "Writing Python APIs to: $DEST_DIR"
        
        # Clean and create destination
        rm -rf "$DEST_DIR/scanoss"
        mkdir -p "$DEST_DIR"
        
        # Generate Python code
        python3 -m grpc_tools.protoc \
                -Iprotobuf \
                --python_out="$DEST_DIR" \
                --grpc_python_out="$DEST_DIR" \
                $(find protobuf/scanoss -type f -name "*.proto") \
                $(find protobuf/protoc-gen-swagger -type f -name "*.proto" 2>/dev/null || true)
        ;;
        
    js)
        DEST_DIR=${DEST_DIR:-"javascript"}
        echo "Writing JavaScript APIs to: $DEST_DIR"
        
        # Clean and create destination
        rm -rf "$DEST_DIR/scanoss"
        mkdir -p "$DEST_DIR"
        
        # Generate JavaScript code
        protoc -Iprotobuf \
               --plugin=protoc-gen-ts="$(which protoc-gen-ts)" \
               --plugin=protoc-gen-grpc="$(which grpc_tools_node_protoc_plugin)" \
               --js_out=import_style=commonjs,binary:"$DEST_DIR" \
               --ts_out=service=grpc-node,mode=grpc-js:"$DEST_DIR" \
               --grpc_out=grpc_js:"$DEST_DIR" \
               $(find protobuf/scanoss -type f -name "*.proto")
        ;;
        
    *)
        echo "Error: Unknown language '$LANG'"
        echo "Supported languages: go, python, js"
        exit 1
        ;;
esac

echo "Successfully generated $LANG protobuf APIs!"