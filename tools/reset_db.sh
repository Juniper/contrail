#!/usr/bin/env bash

TOP=$(dirname "$0")

sleep 1
docker exec contrail_mysql mysql -uroot -pcontrail123 -e "drop database if exists contrail_test;"
sleep 1
docker exec contrail_mysql mysql -uroot -pcontrail123 -e "create database contrail_test;"
sleep 1
docker exec --interactive contrail_mysql mysql -uroot -pcontrail123 contrail_test < $TOP/init_mysql.sql
