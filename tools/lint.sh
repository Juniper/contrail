#!/usr/bin/env bash
# TODO(daniel): run the same set of linters for all Go files

set -o errexit
set -o nounset
set -o pipefail

gometalinter --disable-all -E gofmt -E golint -E vet -t --deadline 1m --vendor -s gen -e bindata.go ./...

# safesql does not skip excluded directories
# megacheck uses staticcheck which consumes too much RAM
# staticcheck consumes too much RAM
gometalinter --enable-all -D safesql -D megacheck -D staticcheck -t --line-length 120 -j2 --deadline 1m \
	--vendor -s gen -e bindata.go ./cmd/watcher ./watcher/... ./pkg/... ./integration/...
