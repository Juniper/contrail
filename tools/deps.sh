#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go get -u github.com/golang/dep/cmd/dep
go get -u github.com/alecthomas/gometalinter
go get -u github.com/mitchellh/gox
go get -u github.com/mattn/goveralls
go install ./vendor/github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
go get -u github.com/go-openapi/spec
go install ./vendor/github.com/gogo/protobuf/protoc-gen-gogo
go install ./vendor/github.com/golang/mock/mockgen
gometalinter --install

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
