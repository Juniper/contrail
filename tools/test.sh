
#!/usr/bin/env bash

TOP=$(cd $(dirname "$0") && cd ../ && pwd)

echo "mode: count" > $TOP/profile.cov

cd $TOP

for dir in $(find . -maxdepth 10 -not -path './vendor/*' -not -path './cmd/*' -not -path '*/.git/*' -not -path '*/test_data/*' -not -path '*/models' -not -path '*/services' -type d);
do
cd $TOP

ls $dir/*.go && result=1 || result=0

if [ $result -eq 1 ]; then
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
fi
done

go tool cover -func $TOP/profile.cov
