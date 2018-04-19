#!/bin/bash

docker rm -f some-cassandra some-zookeeper some-rabbit config-api

docker run --name some-cassandra \
    -e CASSANDRA_START_RPC=true \
    -e CASSANDRA_CLUSTER_NAME=ContrailConfigDB \
    -d cassandra:3.11.1

docker run --name some-zookeeper -d zookeeper:latest
docker run --name some-rabbit -d rabbitmq:3.6.10

docker run \
    --name config-api \
    -p 8082:8082 \
    --link some-cassandra:cassandra \
    --link some-zookeeper:zookeeper \
    --link some-rabbit \
    -d \
    -e CONFIG_API_PORT=8082 \
    -e CONFIG_API_INTROSPECT_PORT=8084 \
    -e LOG_LEVEL=SYS_NOTICE \
    -e log_local=true \
    -e AUTH_MODE=none \
    -e AAA_MODE=none \
    -e CONFIGDB_SERVERS=some-cassandra:9160 \
    -e ZOOKEEPER_SERVERS=some-zookeeper \
    -e RABBITMQ_SERVERS=some-rabbit \
    opencontrailnightly/contrail-controller-config-api

