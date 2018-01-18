#!/usr/bin/env bash

goimports -w -d $(find . -type f -name '*.go' -not -path "./vendor/*") &> /dev/null
gofmt -s -w -d $(find . -type f -name '*.go' -not -path "./vendor/*") &> /dev/null