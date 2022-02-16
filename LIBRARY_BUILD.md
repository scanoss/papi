# SCANOSS gRPC Library Compilation
Details for producing language specific library code from the Protocol Buffer definitions can be found here.

## Protocol Buffer Details
The gRPC definitions can be found the [protobuf](protobuf) folder.

## Compilation
This gRPC API can be implemented in multiple languages. SCANOSS will provide the following:

* Go - [github.com/scanoss/papi](https://github.com/scanoss/papi)
* Python - [scanoss](https://pypi.org/project/scanoss/)
* Javascript - [scanoss](https://www.npmjs.com/package/scanoss)
* Java - TBD

### Go
In order to build the Go library implementation of the proto definitions, the following tools are required:
* [Go](https://go.dev/doc/install)
* [Protoc for Go](https://grpc.io/docs/protoc-installation/)

To build the Go code from the definitions, please run:
```bash
make build_go
```

### Python
In order to build the Python library implementation of the proto definitions, the following tools are required:
* [Python](https://www.python.org/downloads/)
* [grpcio-tools](https://pypi.org/project/grpcio-tools/)

To build the Python code from the definitions, please run:
```bash
make build_python
```

This will save the library/package files in the [python](python) folder.
These files need to be copied into the [scanoss.py](https://github.com/scanoss/scanoss.py) src folder.

Consumption of the Python gRPC APIs is provided through this package. You can install this package from [PyPI](https://pypi.org/project/scanoss/).

### Javascript
In order to build the Javascript library implementation of the proto definitions, the following tools are required:
* [Node](https://nodejs.org/en/download/)
* [protoc] (https://grpc.io/docs/protoc-installation/)
* [ts-protoc-gen](https://www.npmjs.com/package/ts-protoc-gen)



To build the Javascript code from the definitions, please run:
```bash
make build_js
```

This will save the library/package files in the [javascript](javascript) folder.
These files need to be copied into the [scanoss.js](https://github.com/scanoss/scanoss.js) src folder.

Consumption of the Javascript gRPC APIs is provided through this package. You can install this package from [NPM](https://www.npmjs.com/package/scanoss).

### Java
TBD
