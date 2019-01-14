#!/usr/bin/env bash

TOOLSDIR=$(dirname $0)
SUCCESS_MSG="ERROR, SQLSTATE: no results to fetch"
PROJECT='contrail'

echo "Resetting psql database"

echo "Dropping old database"
res=$(NETWORKNAME=$PROJECT docker-compose -f $TOOLSDIR/patroni/docker-compose.yml -p $PROJECT exec -T dbnode bash -c "PGPASSWORD=contrail123 patronictl query -Uroot -d postgres -c \"drop database if exists contrail_test;\" testcluster")
[[ "${res:20}" = "$SUCCESS_MSG" ]] ||  echo "Error while dropping database ${res:20}"

echo "Creating new database"
res=$(NETWORKNAME=$PROJECT docker-compose -f $TOOLSDIR/patroni/docker-compose.yml -p $PROJECT exec -T dbnode bash -c "PGPASSWORD=contrail123 patronictl query -Uroot -d postgres -c \"create database contrail_test;\" testcluster")
[[ "${res:20}" = "$SUCCESS_MSG" ]] || echo "Error while creating database ${res:20}"

echo "Initializing database"
res=$(NETWORKNAME=$PROJECT docker-compose -f $TOOLSDIR/patroni/docker-compose.yml -p $PROJECT exec -T dbnode bash -c "PGPASSWORD=contrail123 patronictl query -Uroot -d contrail_test --file /tools/gen_init_psql.sql testcluster")
[[ "${res:20}" = "$SUCCESS_MSG" ]] || echo "Error while initializing database ${res:20}"

echo "Database initialized"
