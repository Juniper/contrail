#!/usr/bin/env bash

set -o errexit

# TODO: uncomment this once this dep is fixed
# go get -u github.com/go-openapi/spec

go install ./vendor/github.com/golang/dep/cmd/dep
go install ./vendor/github.com/mattn/goveralls
go install ./vendor/github.com/mitchellh/gox
go install ./vendor/github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
go install ./vendor/golang.org/x/tools/cmd/goimports

go install ./vendor/github.com/gogo/protobuf/protoc-gen-gogo
go install ./vendor/github.com/golang/mock/mockgen

curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
	bash -s -- -b $(go env GOPATH)/bin v1.10.2

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
