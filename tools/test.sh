#!/usr/bin/env bash

TOP=$(cd "$(dirname "$0")" && cd ../ && pwd)

COVERPROFILE=${1:--coverprofile=profile.tmp}
COVERMODE='-covermode=atomic'
[ "$COVERPROFILE" = "none" ] && { COVERPROFILE=''; COVERMODE=''; }
[ ! -z "$COVERPROFILE" ] && echo "mode: count" > "$TOP/profile.cov"

for dir in $(go list -f '{{if .TestGoFiles}}{{.Dir}}{{end}}' ./... | \
	grep -v -e 'pkg/cmd' -e 'pkg/services')
do
	cd "$TOP"

	cd "$dir"
	go test -parallel 1 ${COVERMODE} ${COVERPROFILE} .
	result=$?
	if [ $result -ne 0 ]; then
		rm -f ./profile.tmp
		echo "failed"
		exit $result
	fi

	if [ -f profile.tmp ]
	then
		tail -n +2 profile.tmp >> "$TOP/profile.cov"
		rm -f profile.tmp
	fi
done

go tool cover -func "$TOP/profile.cov"
