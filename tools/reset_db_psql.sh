#!/usr/bin/env bash

TOP=/go/github.com/Juniper/contrail/tools

docker exec contrail_postgres psql -U postgres -c "drop database contrail_test"
docker exec contrail_postgres psql -U postgres -c "create database contrail_test"
docker exec --interactive contrail_postgres sh -c "psql -U postgres contrail_test -f $TOP/init_psql.sql"
