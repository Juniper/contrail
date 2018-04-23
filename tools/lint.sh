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
gometalinter \
	--enable-all \
	--exclude "Subprocess launching with variable.*\(gas\)$" \
	--exclude "TLS InsecureSkipVerify.*\(gas\)$" \
	--disable goconst \
	--disable gocyclo \
	--disable megacheck \
	--disable safesql \
	--disable staticcheck \
	--disable test \
	--disable testify \
	--disable vetshadow \
	--tests \
	--aggregate \
	--sort path \
	--deadline 10m \
	--concurrency 2 \
	--line-length 120 \
	--dupl-threshold=115 \
	--vendor \
	--skip pkg/services \
	--skip pkg/serviceif \
	--skip pkg/models \
	--skip pkg/db \
	--skip pkg/compilationif \
	./...

gometalinter \
	--enable-all \
	--exclude "Subprocess launching with variable.*\(gas\)$" \
	--exclude "TLS InsecureSkipVerify.*\(gas\)$" \
	--disable megacheck \
	--disable safesql \
	--disable staticcheck \
	--disable test \
	--disable testify \
	--tests \
	--aggregate \
	--sort path \
	--deadline 10m \
	--concurrency 2 \
	--line-length 120 \
	--dupl-threshold=115 \
	--vendor \
	./cmd/... ./pkg/cmd/... ./pkg/agent/... ./pkg/log/... ./pkg/watcher/...
