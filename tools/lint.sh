#!/usr/bin/env bash

function run_go_tool_fix() {
	local issues
	issues=$(go tool fix --diff ./cmd/ ./extension/ ./pkg/)

	echo "$issues"
	[[ -z "$issues" ]] || exit 1
}

function run_goimports() {
	local dirty_files
	dirty_files="$(goimports -l -local github.com/Juniper/contrail ./cmd/ ./extension/ ./pkg/ | grep -v gen_)"

	echo "$dirty_files"
	[[ -z "$dirty_files" ]] || exit 1
}

run_go_tool_fix

# TODO: remove when goimports is re-enabled in golangci-lint
run_goimports

golangci-lint --config .golangci.yml run ./... 2>&1 || exit 1
