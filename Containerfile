FROM alpine:3.22

#All binaries are download for x86_64, QEMU handles emulation on other architectures

##Installed using apk from alpine
#More info on package pinning here: https://wiki.alpinelinux.org/wiki/Alpine_Package_Keeper#Package_pinning
ARG PYTHON_VERSION=~3.12
ARG PIP_VERSION=~25.1

#Download from Github Releases
ARG PROTOC_VERSION=31.1
ARG PROTOBUF_JS_VERSION=3.21.4
ARG PROTOC_GEN_GO_VERSION=1.36.6
ARG PROTOC_GEN_GO_GRPC_VERSION=1.5.1
ARG GRPC_GATEWAY_VERSION=2.26.3
ARG GRPC_TOOLS_VERSION=1.13.0

# Language runtimes
RUN apk add --no-cache python3=${PYTHON_VERSION} py3-pip=${PIP_VERSION}

# Extra tools
RUN apk add --no-cache wget unzip bash curl

# Copy helper scripts
COPY scripts/install-helpers.sh /usr/local/bin/
COPY scripts/download-tools.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/*.sh

# Install all protobuf tools using helper scripts
RUN /usr/local/bin/download-tools.sh

# Install Python protobuf plugins with frozen versions
RUN pip install --no-cache-dir --break-system-packages \
    grpcio-tools==1.73.0 \
    grpcio==1.73.0

WORKDIR /workspace
ENTRYPOINT ["./scripts/proto-build.sh"]