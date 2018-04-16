#!/usr/bin/env bash

TOP=$(dirname "$0")

echo "Resetting MySQL..."
docker exec contrail_mysql mysql -uroot -pcontrail123 -e "drop database if exists contrail_test;" && echo "Drop db ok"
docker exec contrail_mysql mysql -uroot -pcontrail123 -e "create database contrail_test;" && echo "Create db ok"
docker exec --interactive contrail_mysql mysql -uroot -pcontrail123 contrail_test < $TOP/init_mysql.sql && echo "Init db ok"
