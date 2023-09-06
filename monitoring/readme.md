## Setup

based on [blob](https://www.komu.engineer/blogs/11/opentelemetry-and-go)

1. install go-releaser
    ```bash
    sudo apt install goreleaser
    ```
2. install buf, protoc-gen-buf-lint, and protoc-gen-buf-breaking
    ```bash
    sudo apt update
    sudo apt install protobuf-compiler
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0

    VERSION="1.26.1"
    curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" -o buff
    curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/protoc-gen-buf-breaking-$(uname -s)-$(uname -m)" -o protoc-gen-buf-breaking
    curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/protoc-gen-buf-lint-$(uname -s)-$(uname -m)" -o protoc-gen-buf-lint
    chmod +x ./buf
    chmod +x ./protoc-gen-buf-breaking
    chmod +x ./protoc-gen-buf-lint
    sudo mv ./buf /usr/local/bin/buf
    sudo mv ./protoc-gen-buf-breaking /usr/local/bin/protoc-gen-buf-breaking
    sudo mv ./protoc-gen-buf-lint /usr/local/bin/protoc-gen-buf-lint
    ```