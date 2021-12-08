# SCANOSS Platform 2.0 Public APIs
Welcome to the public APIs for the SCANOSS Platform 2.0

**Warning** Work In Progress

## Repository Structure
This repository is made up of the following components:
* [Protocol Buffer Definitions](protobuf)
* [Go API library Implemention](api)

## Consumption
To consume the Public APIs, simply use one of the following:

* Go - [github.com/scanoss/papi](https://github.com/scanoss/papi)
* Python - [scanoss](https://pypi.org/project/scanoss/)
* Java - TBD

Or alternatively, take the protocol buffer definintions and compile them for your language of choice.

### Go
Examples of consuming the Go library can be found in the [tests/go](tests/go) folder.

### Python
Examples of consuming the Python Library can be found [here](TBD).


## Compilation
This gRPC API can be implemented in multiple languages. SCANOSS will provide the following:

* Go - [github.com/scanoss/papi](https://github.com/scanoss/papi)
* Python - [scanoss](https://pypi.org/project/scanoss/)
* Java - TBD

### Go
In order to build the Go library implementation of the proto definitions, the following tools are required:
* [Go](https://go.dev/doc/install)
* [Protoc for Go](https://grpc.io/docs/protoc-installation/)

To build the Go code from the definitions, please run:
```bash
make build_go
```

