#!/usr/bin/env bash

ContrailPackage="github.com/Juniper/contrail"
TOP=$(cd "$(dirname "$0")" && cd ../ && pwd)

function run_go_tool_fix() {
	local issues
	issues=$(go tool fix --diff ./cmd/ ./pkg/)

	[[ -z "$issues" ]] || (echo "Go tool fix found issues: $issues" && return 1)
}

function run_goimports() {
	local dirty_files
	dirty_files="$(goimports -l -local "$ContrailPackage" ./cmd/ ./pkg/ | grep -v _mock.go)"

	[[ -z "$dirty_files" ]] || (echo "Goimports found issues in files: $dirty_files" | tee -a "$TOP/linter.log" && return 1)
}

run_go_tool_fix || exit 1

# TODO: remove when goimports tool is re-enabled in golangci-lint
run_goimports || exit 1

golangci-lint --config .golangci.yml --verbose run ./... 2>&1 | tee -a "$TOP/linter.log" || exit 1
