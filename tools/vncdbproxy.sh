#!/bin/bash

if [ "$1" == "" ]; then
    echo "Usage: vncdbproxy.sh <CONFIGDBSERVER_IP>"
    exit 1
fi

TOP=$(cd "$(dirname "$0")" && cd ../ && pwd)

docker rm -f vncdbproxy

docker build "$TOP/docker/vnc_db_proxy/" -t vncdbproxy

docker run \
    --name vncdbproxy \
    -p 8091:8082 \
    -d \
    -e CONFIG_API_PORT=8082 \
    -e CONFIG_API_INTROSPECT_PORT=8084 \
    -e LOG_LEVEL=SYS_NOTICE \
    -e log_local=true \
    -e AUTH_MODE=none \
    -e AAA_MODE=cloud-admin \
    -e ZOOKEEPER_SERVERS="$1":2181 \
    -e CONFIGDB_SERVERS="$1":9161 \
    -e RABBITMQ_SERVERS="$1":5673 \
    vncdbproxy
