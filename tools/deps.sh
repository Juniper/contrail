#!/usr/bin/env bash

go get -u github.com/golang/dep/cmd/dep
go get github.com/alecthomas/gometalinter
go get github.com/mitchellh/gox
go get github.com/mattn/goveralls
gometalinter --install
go get github.com/gogo/protobuf/protoc-gen-gogo
go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
go get github.com/go-openapi/spec
