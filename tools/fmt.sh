#!/usr/bin/env bash

goimports -v -w $(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "*pkg/generated/models/generated*")