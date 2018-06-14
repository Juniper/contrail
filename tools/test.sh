
#!/usr/bin/env bash

TOP=$(cd $(dirname "$0") && cd ../ && pwd)

echo "mode: count" > $TOP/profile.cov

for dir in $(go list -f '{{if .TestGoFiles}}{{.Dir}}{{end}}' ./... | \
                 grep -v -e 'pkg/cmd' -e 'pkg/services')
do
    cd $TOP

    cd $dir
    go test -parallel 1 -covermode=atomic -coverprofile=profile.tmp .
    result=$?
    if [ $result -ne 0 ]; then
        echo "failed"
        exit $result
    fi

    if [ -f profile.tmp ]
    then
        cat profile.tmp | tail -n +2 >> $TOP/profile.cov
        rm profile.tmp
    fi
done

go tool cover -func $TOP/profile.cov
