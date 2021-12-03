# SCANOSS Protocol Buffer (gRPC) Definitions
This folder contains all the SCANOSS proto definitions for its gRPC services.

## Structure
The API is made up of different services
* Dependencies
* Scanning

These services have some shared items, which can be found in Common.

## Compilation

### Go
From the root of the project run:
```bash
protoc --proto_path=. --go_out=api --go-grpc_out=api `find protobuf/scanoss -type f -name "*.proto" -print`
mv api/github.com/scanoss/papi/api/* api/
rm -rf api/github.com
```