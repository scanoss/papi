# SCANOSS Platform 2.0 Public APIs
Welcome to the public APIs for the SCANOSS Platform 2.0

**Warning** Work In Progress **Warning**

## Repository Structure
This repository is made up of the following components:
* [Protocol Buffer Definitions](protobuf)
* [Go API library Implementation](api)

## Consumption
To consume the Public APIs, simply use one of the following:

* Go - [github.com/scanoss/papi](https://github.com/scanoss/papi)
* Python - [scanoss](https://pypi.org/project/scanoss/)
* Javascript - [scanoss](https://www.npmjs.com/package/scanoss)
* Java - TBD

Or alternatively, take the protocol buffer definitions and compile them for your language of choice.

Details on how the supported languages are compiled can be found [here](LIBRARY_BUILD.md).

### Go
Examples of consuming the Go library can be found in the [tests/go](tests/go) folder.

### Python
Examples of consuming the Python Library can be found [here](https://github.com/scanoss/scanoss.py).

### Javascript
Examples of consuming the Javascript Library can be found [here](https://github.com/scanoss/scanoss.js).

### Java
TBD

## Build
For details on how to compile the supported language libraries, please look in [LIBRARY_BUILD.md](LIBRARY_BUILD.md).