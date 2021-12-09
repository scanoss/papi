# SCANOSS Protocol Buffer (gRPC) Definitions
This folder contains all the SCANOSS proto definitions for its gRPC services.

## Structure
The API is made up of different services
* [Dependencies](scanoss/api/dependencies/v2/scanoss-dependencies.proto)
* [Scanning](scanoss/api/scanning/v2/scanoss-scanning.proto)

These services have some shared items, which can be found in [Common](scanoss/api/common/v2/scanoss-common.proto).

## Compilation
For details on how to compile the supported language libraries, please look in [LIBRARY_BUILD.md](../LIBRARY_BUILD.md).

## Support Scripts
There are two support scripts in this folder:
* [build.sh](build.sh)
* [copy_python.sh](copy_python.sh)

### build.sh
The [build.sh](build.sh) script provides the support for compiling the library code for the supported language types.

Examples for consuming this script can be found in the [Makefile](../Makefile).

### copy_python.sh
The [copy_python.sh](copy_python.sh) script provides the support for copying the generated Python library code from this repository into the scanoss.py project on the local file system.
It defaults to assuming the scanoss.py repo is in the parent folder beside the papi repo.

Examples for consuming this script can be found in the [Makefile](../Makefile).
