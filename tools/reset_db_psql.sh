#!/usr/bin/env bash

TOP=/go/github.com/Juniper/contrail/tools

echo "Mounts:"
docker inspect -f '{{ range $i, $m := .Mounts }}{{ $m.Source }}:{{ $m.Destination }}{{"\n"}}{{end}}' contrail_postgres

docker exec contrail_postgres psql -U postgres -c "drop database contrail_test"
docker exec contrail_postgres psql -U postgres -c "create database contrail_test"
docker exec --interactive contrail_postgres psql -U postgres contrail_test -f $TOP/init_psql.sql
