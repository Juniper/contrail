#!/usr/bin/env bash

set -o errexit

go install github.com/gogo/protobuf/protoc-gen-gogofaster
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
go install github.com/golang/mock/mockgen
go install github.com/mitchellh/gox
go install golang.org/x/lint/golint
go install golang.org/x/tools/cmd/goimports
go install github.com/Juniper/asf/cmd/contrailschema

curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
    bash -s -- -b $(go env GOPATH)/bin v1.10.2

if [[ $(./bin/protoc --version) == "libprotoc 3.5.1" ]]; then
    echo "Protoc already installed - skipping"
    exit 0
fi

if [ "$(uname)" == 'Darwin' ]; then
    wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-osx-x86_64.zip
    unzip -o protoc-3.5.1-osx-x86_64.zip "bin/protoc"
    rm protoc-3.5.1-osx-x86_64.zip
elif [ "$(expr substr $(uname -s) 1 5)" == 'Linux' ]; then
    wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
    unzip -o protoc-3.5.1-linux-x86_64.zip "bin/protoc"
    rm protoc-3.5.1-linux-x86_64.zip
else
    echo "Your platform ($(uname -a)) is not supported."
    echo "Please manually install protoc"
fi

