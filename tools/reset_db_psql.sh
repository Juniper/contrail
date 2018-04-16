#!/usr/bin/env bash

TOP=$(dirname "$0")

sleep 1
docker exec contrail_postgres psql -U postgres -c "drop database contrail_test"
sleep 1
docker exec contrail_postgres psql -U postgres -c "create database contrail_test"
sleep 1
docker exec --interactive contrail_postgres psql -U postgres contrail_test < $TOP/init_psql.sql
sleep 1
