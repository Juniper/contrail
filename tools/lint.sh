#!/usr/bin/env bash

ContrailPackage="github.com/Juniper/contrail"

function run_go_tool_fix() {
	local issues
	issues=$(go tool fix --diff ./cmd/ ./extension/ ./pkg/)

	[[ -z "$issues" ]] || (echo "Go tool fix found issues: $issues" && exit 1)
}

function run_goimports() {
	local dirty_files
	dirty_files="$(goimports -l -local "$ContrailPackage" ./cmd/ ./extension/ ./pkg/ | grep -v _mock.go)"

	[[ -z "$dirty_files" ]] || (echo "Goimports found issues in files: $dirty_files" && exit 1)
}

run_go_tool_fix

# TODO: remove when goimports tool is re-enabled in golangci-lint
run_goimports

golangci-lint --config .golangci.yml run ./... 2>&1 || exit 1
