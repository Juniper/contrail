#!/usr/bin/env bash

TOP=$(dirname "$0")

echo "Resetting Postgres..."
docker exec contrail_postgres psql -U postgres -c "drop database if exists contrail_test" && echo "Drop db ok"
docker exec contrail_postgres psql -U postgres -c "create database contrail_test" && echo "Create db ok"
ocker exec --interactive contrail_postgres psql -U postgres contrail_test < $TOP/init_psql.sql && echo "Init db ok"
