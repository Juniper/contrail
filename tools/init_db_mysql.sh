#!/usr/bin/env bash

TOP=/go/src/github.com/Juniper/contrail/tools

docker exec --interactive contrail_mysql sh -c "mysql -uroot -pcontrail123 contrail_test < $TOP/init_mysql.sql"
go run cmd/contrailutil/main.go convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail.yml
