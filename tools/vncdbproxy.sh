#!/bin/bash

usage() {
    echo "Usage: $0 [-n <mode>] [-k] [-r <host:port>] [-z <host:port>] [-c <host:port>]" 1>&2
    echo "-k => Don't remove dockers before running new ones"
    echo "-n => Use specified 'NetworkMode' for docker (default: 'host')"
    echo "-r => Rabbit (default: 'localhost:5672')"
    echo "-z => Zookeeper (default: 'localhost:2181')"
    echo "-c => ConfigDB (default: 'localhost:9160')"
    exit 0
}

Network="host"
RemoveDockers=1
ConfigDB="localhost:9160"
Zookeeper="localhost:2181"
Rabbit="localhost:5672"

while getopts ":n:k:d:r:z:" o; do
    case "${o}" in
        n)
            Network=${OPTARG}
            ;;
        k)
            RemoveDockers=0
            ;;
        c)
            ConfigDB=${OPTARG}
            ;;
        r)
            Rabbit=${OPTARG}
            ;;
        z)
            Zookeeper=${OPTARG}
            ;;
        *)
            usage
            ;;
    esac
done
shift $((OPTIND-1))

[ $RemoveDockers -eq 1 ] && docker rm -f vncdbproxy

TOP=$(cd "$(dirname "$0")" && cd ../ && pwd)

docker build "$TOP/docker/vnc_db_proxy/" -t vncdbproxy

docker run \
    --name vncdbproxy \
    --network ${Network} \
    -p 9082:9082 \
    -d \
    -e CONFIG_API_PORT=9082 \
    -e CONFIG_API_INTROSPECT_PORT=9084 \
    -e LOG_LEVEL=SYS_NOTICE \
    -e log_local=true \
    -e AUTH_MODE=none \
    -e AAA_MODE=cloud-admin \
    -e ZOOKEEPER_SERVERS="${Zookeeper}" \
    -e CONFIGDB_SERVERS="${ConfigDB}" \
    -e RABBITMQ_SERVERS="${Rabbit}" \
    vncdbproxy
