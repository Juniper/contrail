#!/bin/bash

DIR=$(dirname $0)
USER=root
PASS=contrail123
SUCCESS_MSG="ERROR, SQLSTATE: no results to fetch"

res=$(docker-compose -f $DIR/patroni/docker-compose.yml -p "contrail" exec -T dbnode bash -c "PGPASSWORD=$PASS patronictl query -U $USER -d postgres -c \"drop database contrail_test;\" testcluster")
if [ "${res:20}" != "$SUCCESS_MSG" ]; then echo "Error while dropping database ${res:20}"; fi

res=$(docker-compose -f $DIR/patroni/docker-compose.yml -p "contrail" exec -T dbnode bash -c "PGPASSWORD=$PASS patronictl query -U $USER -d postgres -c \"create database contrail_test;\" testcluster")
if [ "${res:20}" != "$SUCCESS_MSG" ]; then echo "Error while creating database ${res:20}"; fi

res=$(docker-compose -f $DIR/patroni/docker-compose.yml -p "contrail" exec -T dbnode bash -c "PGPASSWORD=$PASS patronictl query -U $USER -d contrail_test --file /tools/init_psql.sql testcluster")
if [ "${res:20}" != "$SUCCESS_MSG" ]; then echo "Error while initializing database ${res:20}"; fi
