#!/usr/bin/env bash

TOP=$(dirname "$0")
mysql -uroot -pcontrail123 -e "drop database if exists contrail_test;"
mysql -uroot -pcontrail123 -e "create database contrail_test;"
mysql -uroot -pcontrail123 contrail_test < $TOP/init.sql