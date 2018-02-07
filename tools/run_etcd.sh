#!/usr/bin/env bash
# Run etcd service.

set -o errexit
set -o nounset
set -o pipefail

function main {
	echo "Running etcd"
	etcd --data-dir /tmp/test-etcd &
}

main
