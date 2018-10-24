#!/usr/bin/env bash

ContrailPackage="github.com/Juniper/contrail"

function run_go_tool_fix() {
	local issues
	issues=$(go tool fix --diff ./cmd/ ./extension/ ./pkg/)

	[[ -z "$issues" ]] || (echo "Go tool fix found issues: $issues" && return 1)
}

function run_goimports() {
	local issues
	issues="$(goimports -d -local "$ContrailPackage" ./cmd/ ./extension/ ./pkg/ | grep -v _mock.go)"

	[[ -z "$issues" ]] || (echo "Goimports found issues: $issues" && return 1)
}

run_go_tool_fix || exit 1

# TODO: remove when goimports tool is re-enabled in golangci-lint
run_goimports || exit 1

golangci-lint --config .golangci.yml run ./... 2>&1 || exit 1
