#!/usr/bin/env bash

set -o errexit

[[ -z $(go tool fix --diff ./pkg/) ]]

golangci-lint --config .golangci.yml run ./... 2>&1
