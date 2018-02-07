#!/usr/bin/env bash

go get -u github.com/golang/dep/cmd/dep
go get -u github.com/alecthomas/gometalinter
go get -u github.com/mitchellh/gox
go get -u github.com/mattn/goveralls
gometalinter --install
