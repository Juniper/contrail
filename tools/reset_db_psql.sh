#!/usr/bin/env bash

TOOLSDIR=$(dirname $0)
SUCCESS_MSG="ERROR, SQLSTATE: no results to fetch"

echo "Resetting psql database"

res=$(docker-compose -f $TOOLSDIR/patroni/docker-compose.yml -p "contrail" exec -T dbnode bash -c "PGPASSWORD=contrail123 patronictl query -Uroot -d postgres -c \"drop database contrail_test;\" testcluster")
if [ "${res:20}" != "$SUCCESS_MSG" ]; then echo "Error while dropping database ${res:20}"; fi

res=$(docker-compose -f $TOOLSDIR/patroni/docker-compose.yml -p "contrail" exec -T dbnode bash -c "PGPASSWORD=contrail123 patronictl query -Uroot -d postgres -c \"create database contrail_test;\" testcluster")
if [ "${res:20}" != "$SUCCESS_MSG" ]; then echo "Error while creating database ${res:20}"; fi

res=$(docker-compose -f $TOOLSDIR/patroni/docker-compose.yml -p "contrail" exec -T dbnode bash -c "PGPASSWORD=contrail123 patronictl query -Uroot -d contrail_test --file /tools/init_psql.sql testcluster")
if [ "${res:20}" != "$SUCCESS_MSG" ]; then echo "Error while initializing database ${res:20}"; fi
