#!/usr/bin/env bash

TOP=$(dirname "$0")

docker exec contrail_mysql mysql -uroot -pcontrail123 -e "drop database if exists contrail_test;"
docker exec contrail_mysql mysql -uroot -pcontrail123 -e "create database contrail_test;"
docker exec --interactive contrail_mysql mysql -uroot -pcontrail123 contrail_test < $TOP/init_mysql.sql