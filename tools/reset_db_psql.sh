#!/bin/bash

#DIR=$(dirname $0)
#DIR=/psq
#USER=root
#PASS=contrail123
#SUCCESS_MSG="ERROR, SQLSTATE: no results to fetch"

#PGPASSWORD=$PASS psql -h $HOST -p $PORT -U $USER postgres -c "drop database contrail_test"
#PGPASSWORD=$PASS psql -h $HOST -p $PORT -U $USER postgres -c "create database contrail_test"
#PGPASSWORD=$PASS psql -h $HOST -p $PORT -U $USER contrail_test -f $DIR/init_psql.sql > /dev/null
#
#res=$(docker exec contrail_dbnode_1 bash -c "PGPASSWORD=$PASS patronictl query -U $USER -d postgres -c \"drop database contrail_test;\" testcluster")
#if [ "${res:20}" != "$SUCCESS_MSG" ]; then echo "Error while dropping database ${res:20}"; fi
#
#res=$(docker exec contrail_dbnode_1 bash -c "PGPASSWORD=$PASS patronictl query -U $USER -d postgres -c \"create database contrail_test;\" testcluster")
#if [ "${res:20}" != "$SUCCESS_MSG" ]; then echo "Error while creating database ${res:20}"; fi
#
#res=$(docker exec contrail_dbnode_1 bash -c "PGPASSWORD=$PASS patronictl query -U $USER -d contrail_test --file $DIR/init_psql.sql testcluster")
#if [ "${res:20}" != "$SUCCESS_MSG" ]; then echo "Error while initializing database ${res:20}"; fi

TOP=/go/src/github.com/Juniper/contrail/tools

echo "Mounts:"
docker inspect -f '{{ range $i, $m := .Mounts }}{{ $m.Source }}:{{ $m.Destination }}{{"\n"}}{{end}}' contrail_postgres

docker exec contrail_postgres psql -U postgres -c "drop database contrail_test"
docker exec contrail_postgres psql -U postgres -c "create database contrail_test"
docker exec --interactive contrail_postgres psql -U postgres contrail_test -f $TOP/init_psql.sql > /dev/null
