#!/usr/bin/env bash

TOP=/go/src/github.com/Juniper/contrail/tools

echo "Mounts:"
docker inspect -f '{{ range $i, $m := .Mounts }}{{ $m.Source }}:{{ $m.Destination }}{{"\n"}}{{end}}' contrail_mysql

docker exec contrail_mysql mysql -uroot -pcontrail123 -e "drop database if exists contrail_test;"
docker exec contrail_mysql mysql -uroot -pcontrail123 -e "create database contrail_test;"
docker exec --interactive contrail_mysql sh -c "mysql -uroot -pcontrail123 contrail_test < $TOP/gen_init_mysql.sql"
