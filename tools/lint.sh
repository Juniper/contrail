#!/usr/bin/env bash

set -o pipefail

TOP=$(cd "$(dirname "$0")" && cd ../ && pwd)

run() {
    run_go_tool_fix || return 1
    run_golangci_lint || return 1
    run_golint || return 1
}

run_go_tool_fix() {
    local issues
    issues=$(go tool fix --diff ./cmd/ ./pkg/)

    [[ -z "$issues" ]] || (echo "Go tool fix found issues: $issues" && return 1)
}

run_golangci_lint() {
    GL_DEBUG=env golangci-lint \
        --config .golangci.yml --verbose \
        run ./... 2>&1 | tee -a "$TOP/linter.log" || return 1
}

run_golint() {
    golint -set_exit_status ./... || return 1
}

run
