#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

gometalinter --disable-all -E gofmt -E golint -E vet -t --deadline 1m --vendor -s pkg/generated ./...

# TODO(nati) apply more lint tools
#gometalinter --enable-all -D safesql -D megacheck -D staticcheck -t --line-length 120 -j2 --deadline 1m \
#	--vendor -s pkg/generated ./...