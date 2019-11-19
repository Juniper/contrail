#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

TOP=$(cd "$(dirname "$0")" && cd ../ && pwd)

FILTER="$1"
COVERPROFILE=${2:--coverprofile=profile.tmp}
COVERMODE='-covermode=atomic'
[ "$COVERPROFILE" = "none" ] && { COVERPROFILE=''; COVERMODE=''; }
[ ! -z "$COVERPROFILE" ] && echo "mode: count" > "$TOP/profile.cov"

# test_directories lists directories that contain _test.go files
# either inside Go package (.TestGoFiles) or outside Go package (.XTestGoFiles, e.g. in "foo_test" package).
function test_directories {
	cd "$TOP"
	go list -f '{{if (or .TestGoFiles .XTestGoFiles)}}{{.Dir}} {{.Name}}{{end}}' ./... | awk '
	BEGIN{
		split("'$FILTER'", tmp, ",")
		for(i in tmp)
			filter[tmp[i]]=1
	}
	{
		if(!length(filter) || $2 in filter)
			print $1
	}'
}

for dir in $(test_directories)
do
	cd "$dir"
	go test ${COVERMODE} ${COVERPROFILE} . | tee -a "$TOP/unittests.log"
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

[ -z "$COVERPROFILE" ] && exit 0
go tool cover -func "$TOP/profile.cov" > "$TOP/coverage.log"
