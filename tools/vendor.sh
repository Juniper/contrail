#!/usr/bin/env bash

set -x
  
rm -rf ./vendor
go mod vendor
# get https://github.com/gogo/protobuf/tree/master/protobuf
git clone https://github.com/gogo/protobuf.git /tmp/protobuf
cp -r /tmp/protobuf/protobuf ./vendor/github.com/gogo/protobuf/
[[ -z $(git status --porcelain -- vendor) ]] || { echo 'vendor files has changed. Please update the vendor files and commit it into repository.'; git diff -- vendor; exit 1; }
