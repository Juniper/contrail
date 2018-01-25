#!/usr/bin/env bash
# Run integration tests.

set -o nounset
set -o pipefail

function main {
	echo "Running integration tests"
	go test -tags=integration ./integration/...
}

main
