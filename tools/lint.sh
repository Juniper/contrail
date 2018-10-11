#!/usr/bin/env bash

function run_go_tool_fix() {
	local issues
	issues=$(go tool fix --diff ./cmd/ ./extension/ ./pkg/)

	echo "Go tool fix found issues: $issues"
	[[ -z "$issues" ]] || exit 1
}

run_go_tool_fix
golangci-lint --config .golangci.yml run ./... 2>&1 || exit 1
