#!/usr/bin/env bash

TOP=/go/github.com/Juniper/contrail/tools

docker exec contrail_mysql mysql -uroot -pcontrail123 -e "drop database if exists contrail_test;"
docker exec contrail_mysql mysql -uroot -pcontrail123 -e "create database contrail_test;"
docker exec --interactive contrail_mysql sh -c "mysql -uroot -pcontrail123 contrail_test < $TOP/init_mysql.sql"
