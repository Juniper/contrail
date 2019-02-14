#!/usr/bin/env bash

gofiles=$(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "*gen*")
goimports -v -w ${gofiles}
gofmt -s -w ${gofiles}
