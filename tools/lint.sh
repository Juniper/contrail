#!/usr/bin/env bash
# Following static analysis tools are disabled:
# * dupl which consumes too much RAM and is too slow
# * megacheck which uses staticcheck which consumes too much RAM
# * safesql which does not omit skipped directories
# * staticcheck which consumes too much RAM
# * test which runs tests and reports failures - that is unnecessary
# * testify which runs tests and reports failures - that is unnecessary
# Concurrency is reduced in order to reduce RAM consumption.

set -o errexit
set -o nounset
set -o pipefail

[[ -z `go tool fix --diff ./pkg/` ]]

gometalinter --config .gometalinter.json ./...
