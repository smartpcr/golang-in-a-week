#!/bin/bash

set -x # have bash print command been ran
set -e # fail if any command fails

sudo apt update

# Determine the location of the installed openssl
openssl_path=$(which openssl)

# Extract the version number of the installed openssl
openssl_version=$(openssl version | awk '{print $2}')

# Check if the openssl version is less than 3.1.2
if [[ "$openssl_version" < "3.1.2" ]]; then

    # Check if openssl is installed via Homebrew/Linuxbrew
    if [[ $openssl_path == *"/linuxbrew/"* ]]; then
        echo "OpenSSL installed via Linuxbrew. Upgrading..."
        brew update
        brew upgrade openssl

    # Check if openssl is installed via system/apt
    elif [[ $openssl_path == *"/usr/bin/"* ]]; then
        echo "OpenSSL installed via system package manager. Upgrading..."
        sudo apt update
        sudo apt install openssl

    else
        echo "Unrecognized OpenSSL installation. Please upgrade manually."
    fi

    # Print the current version of openssl after the upgrade
    openssl version

else
    echo "OpenSSL version is $openssl_version which is >= 3.1.2. No need to upgrade."
fi

sudo apt install goreleaser
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