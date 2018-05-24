#!/usr/bin/env bash

go get -u github.com/golang/dep/cmd/dep
go get github.com/alecthomas/gometalinter
go get github.com/mitchellh/gox
go get github.com/mattn/goveralls
gometalinter --install
go get github.com/gogo/protobuf/protoc-gen-gogo
go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
go get github.com/go-openapi/spec
go get github.com/golang/mock/gomock
go install ./vendor/github.com/golang/mock/mockgen

if [ "$(uname)" == 'Darwin' ]; then
    wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-osx-x86_64.zip
    unzip protoc-3.5.1-osx-x86_64.zip "bin/protoc"
    rm protoc-3.5.1-osx-x86_64.zip
elif [ "$(expr substr $(uname -s) 1 5)" == 'Linux' ]; then
    wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
    unzip protoc-3.5.1-linux-x86_64.zip "bin/protoc"
    rm protoc-3.5.1-linux-x86_64.zip
else
	echo "Your platform ($(uname -a)) is not supported."
    echo "Please manually install protoc"
fi
