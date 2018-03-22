#!/usr/bin/env bash

TOP=$(dirname "$0")

mysql -uroot -pcontrail123 -h 127.0.0.1 -e "drop database if exists contrail_test;"
mysql -uroot -pcontrail123 -h 127.0.0.1  -e "create database contrail_test;"
mysql -uroot -pcontrail123 -h 127.0.0.1 contrail_test < $TOP/init_mysql.sql