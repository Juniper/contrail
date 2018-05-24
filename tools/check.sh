#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

dep ensure -dry-run -no-vendor
