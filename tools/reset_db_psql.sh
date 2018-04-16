#!/usr/bin/env bash

TOP=$(dirname "$0")

echo "Drop db"
docker exec contrail_postgres psql -U postgres -c "drop database if exists contrail_test" && echo "ok"

echo "Create db"
docker exec contrail_postgres psql -U postgres -c "create database contrail_test" && echo "ok"

echo "Fill db"
docker exec --interactive contrail_postgres psql -U postgres contrail_test < $TOP/init_psql.sql && echo "ok"
