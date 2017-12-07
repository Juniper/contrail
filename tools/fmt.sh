#!/usr/bin/env bash

goimports -w -d $(find . -type f -name '*.go' -not -path "./vendor/*")
gofmt -s -w -d $(find . -type f -name '*.go' -not -path "./vendor/*")