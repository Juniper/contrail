#!/bin/bash

usage() {
    echo "Usage: $0 [-n <mode>] [-k] [-r <host:port>] [-z <host:port>] [-c <host:port>]" 1>&2
    echo "-k => Don't remove dockers before running new ones"
    echo "-n => Use specified 'NetworkMode' for docker (default: 'bridge')"
    echo "-l => link containers"
    echo "-r => Rabbit (default: 'some-cassandra:5672')"
    echo "-z => Zookeeper (default: 'some-zookeeper:2181')"
    echo "-c => ConfigDB (default: 'some-rabbit:9160')"
    exit 0
}

Network="bridge"
RemoveDockers=1
ConfigDB="some-cassandra:9160"
Zookeeper="some-zookeeper:2181"
Rabbit="some-rabbit:5672"

while getopts ":n:k:d:r:z:l:c:" o; do
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
        l)
            Linked+="--link ${OPTARG} "
            ;;
        *)
            usage
            ;;
    esac
done
shift $((OPTIND-1))
[ $RemoveDockers -eq 1 ] && docker rm -f vncdbproxy
[ ! -z "$Linked" ] && Linked=$(echo "${Linked}" | xargs)
[ "$Network" == 'bridge' ] && [ -z "$Linked" ] && Linked="--link some-cassandra:cassandra --link some-zookeeper:zookeeper --link some-rabbit"

TOP=$(cd "$(dirname "$0")" && cd ../../ && pwd)
echo "sdfdsf ${Linked}"
docker build "$TOP/docker/vnc_db_proxy/" -t vncdbproxy

docker run ${Linked} \
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
