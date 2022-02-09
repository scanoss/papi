# SCANOSS gRPC Library Tests
This folder contains sample go code to try out the SCANOSS gRPC PAPI APIs.

There is both client and server examples.

## Server Samples
There are two server implementations present:
* [Dependency](dep_server)
* [Scanning](scan_server)

These two servers provide Go code to instantiate a sample service.

### Dependency Server
To launch, simply run:
```bash
go run tests/go/dep_server/dep-server-test.go
```
The default port is 50051, but it can be modified by supplying the `-port` option.

### Scanning Server
To launch, simply run:
```bash
go run tests/go/scan_server/scan-server-test.go
```
The default port is 50052, but it can be modified by supplying the `-port` option.

## Client Samples
The Go client implements the Echo gRPC call on both servers. To launch, simple run:
```bash
go run tests/go/client/client-test.go
```
The default host/port used for the dependency service is localhost:50051, but it can be modified by supplying the `-d_addr` option.

The default host/port used for the scanning services is localhost:50052, but it can be modified by supplying the `-s_addr` option.
