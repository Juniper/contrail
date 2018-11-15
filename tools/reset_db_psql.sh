#!/bin/bash

#DIR=$(dirname $0)
DIR=/psq
HOST=localhost
PORT=5432
USER=root
PASS=contrail123

#PGPASSWORD=$PASS psql -h $HOST -p $PORT -U $USER postgres -c "drop database contrail_test"
#PGPASSWORD=$PASS psql -h $HOST -p $PORT -U $USER postgres -c "create database contrail_test"
#PGPASSWORD=$PASS psql -h $HOST -p $PORT -U $USER contrail_test -f $DIR/init_psql.sql > /dev/null

docker exec contrail_dbnode_1 psql -h $HOST -p $PORT -U $USER postgres -c "drop database contrail_test"
docker exec contrail_dbnode_1 psql -h $HOST -p $PORT -U $USER postgres -c "create database contrail_test"
docker exec contrail_dbnode_1 psql -h $HOST -p $PORT -U $USER contrail_test -f $DIR/init_psql.sql > /dev/null
