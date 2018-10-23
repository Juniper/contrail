#!/usr/bin/env bash

TOP=/go/src/github.com/Juniper/contrail/tools

docker exec --interactive contrail_postgres psql -U postgres contrail_test -f $TOP/init_psql.sql
go run cmd/contrailutil/main.go convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail_postgres.yml
