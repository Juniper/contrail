#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

[[ -z `go tool fix --diff ./pkg/` ]]

golangci-lint --config .golangci-lint.yml run ./... 2>&1
