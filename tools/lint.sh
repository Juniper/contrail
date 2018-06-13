#!/usr/bin/env bash
# Following static analysis tools are disabled:
# * megacheck which uses staticcheck which consumes too much RAM
# * safesql which does not omit skipped directories
# * staticcheck which consumes too much RAM
# * test which runs tests and reports failures - that is unnecessary
# * testify which runs tests and reports failures - that is unnecessary
# Concurrency is reduced in order to reduce RAM consumption.

set -o errexit
set -o nounset
set -o pipefail

# Several tools for majority of the code are disabled.
# TODO(daniel): run the same set of tools for all Go files
# TODO(tomasz): enable dupl, lll, unparam
gometalinter \
    --config .gometalinter.json \
    --disable gocyclo \
    ./...

gometalinter \
    --config .gometalinter.json \
	./cmd/... ./pkg/cmd/... ./pkg/agent/... ./pkg/log/... ./pkg/sync/... ./pkg/testutil/...
