#!/usr/bin/env bash

TOP=$(dirname "$0")

docker exec contrail_postgres psql -U postgres -c "drop database if exists contrail_test"
docker exec contrail_postgres psql -U postgres -c "create database contrail_test"
docker exec --interactive contrail_postgres psql -U postgres contrail_test < $TOP/init_psql.sql
