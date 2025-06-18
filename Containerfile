FROM alpine:3.22

#All binaries are download for x86_64, QEMU handles emulation on other architectures

##Installed using apk from alpine
#More info on package pinning here: https://wiki.alpinelinux.org/wiki/Alpine_Package_Keeper#Package_pinning
ARG PYTHON_VERSION=~3.12
ARG PIP_VERSION=~25.1

#Download compiled from GoLang Official Website
ARG GO_VERSION=1.23.4

#Download from Github Releases
ARG PROTOC_VERSION=31.1
ARG PROTOBUF_JS_VERSION=3.21.4

# Language runtimes
RUN apk add --no-cache python3=${PYTHON_VERSION} py3-pip=${PIP_VERSION}

# Extra tools
RUN apk add --no-cache wget unzip bash


# Install Go
RUN echo "Installing Go ${GO_VERSION}" && \
    wget -O /tmp/go.tar.gz "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

# Set up Go environment
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:/usr/local/bin:$PATH


# Install protoc compiler
RUN echo "Installing protoc ${PROTOC_VERSION}" && \
    wget -O /tmp/protoc.zip "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip" && \
    unzip /tmp/protoc.zip -d /usr/local && \
    chmod +x /usr/local/bin/protoc && \
    rm /tmp/protoc.zip

# Install protobuf-javascript plugin
RUN echo "Installing protobuf-javascript ${PROTOBUF_JS_VERSION}" && \
    wget -O /tmp/protobuf-js.tar.gz "https://github.com/protocolbuffers/protobuf-javascript/releases/download/v${PROTOBUF_JS_VERSION}/protobuf-javascript-${PROTOBUF_JS_VERSION}-linux-x86_64.tar.gz" && \
    tar -C /tmp -xzf /tmp/protobuf-js.tar.gz && \
    cp /tmp/bin/protoc-gen-js /usr/local/bin/protoc-gen-js && \
    chmod +x /usr/local/bin/protoc-gen-js && \
    rm -rf /tmp/protobuf-js.tar.gz /tmp/bin


# Install grpc-tools binary manually
RUN echo "Installing grpc-tools 1.13.0" && \
    wget -O /tmp/grpc-tools.tar.gz "https://node-precompiled-binaries.grpc.io/grpc-tools/v1.13.0/linux-x64.tar.gz" && \
    tar -C /tmp -xzf /tmp/grpc-tools.tar.gz && \
    cp /tmp/bin/* /usr/local/bin/ && \
    chmod +x /usr/local/bin/grpc_node_plugin && \
    rm -rf /tmp/grpc-tools.tar.gz /tmp/bin

RUN npm install -g protoc-gen-ts@0.8.7


RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1


# Install gRPC <-> HTTP Gateway. Used to generate the gateway code
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.26.3 && \
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.16.0

# Install Python protobuf plugins with frozen versions
RUN pip install --no-cache-dir --break-system-packages \
    grpcio-tools==1.73.0 \
    grpcio==1.73.0

RUN protoc --version

WORKDIR /workspace
ENTRYPOINT ["./scripts/proto-build.sh"]