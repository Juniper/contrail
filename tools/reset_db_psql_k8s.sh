#!/usr/bin/env bash

TOOLSDIR=$(dirname $0)
SUCCESS_MSG="ERROR, SQLSTATE: no results to fetch"
PROJECT='contrail'

echo "Resetting psql database"

echo "Dropping old database"
res=$(kubectl exec contrail-0 -- sh -c "PGPASSWORD=contrail123 patronictl query -Uroot -d postgres -c \"drop database if exists contrail_test;\" contrail")
[[ "${res:20}" = "$SUCCESS_MSG" ]] ||  echo "Error while dropping database ${res:20}"

echo "Creating new database"
res=$(kubectl exec contrail-0 -- sh -c "PGPASSWORD=contrail123 patronictl query -Uroot -d postgres -c \"create database contrail_test;\" contrail")
[[ "${res:20}" = "$SUCCESS_MSG" ]] || echo "Error while creating database ${res:20}"

echo "Initializing database"
res=$(kubectl exec contrail-0 -- sh -c "PGPASSWORD=contrail123 patronictl query -Uroot -d contrail_test --file /init_psql.sql testcluster")
[[ "${res:20}" = "$SUCCESS_MSG" ]] || echo "Error while initializing database ${res:20}"

echo "Database initialized"
