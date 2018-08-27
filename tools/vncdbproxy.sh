#!/bin/bash

TOP=$(cd "$(dirname "$0")" && cd ../ && pwd)

Usage()
{
	echo "Usage: vncdbproxy.sh [-k] [-n <NetworkMode>] [CONFIGDBSERVER_IP]"
}

ServicesIpAddress='127.0.0.1'
RemoveDockers=1
Network='bridge' # This is default for `docker run` if no `--net` param is specified
while :; do
	case "$1" in
		'-n') Network="$2"; shift 2;;
		'-k') RemoveDockers=0; shift;;
		'-h') Usage; exit 0;;
		*) break;;
	esac
done

if [ "$Network" = 'bridge' ]; then
	[ "$1" = '' ] && { Usage; exit 1; }
	ServicesIpAddress="$1"
fi

[ $RemoveDockers -eq 1 ] && docker rm -f vncdbproxy
docker build "$TOP/docker/vnc_db_proxy/" -t vncdbproxy

docker run \
    --name vncdbproxy \
    --net "$Network" \
    -p 8091:8082 \
    -d \
    -e CONFIG_API_PORT=8082 \
    -e CONFIG_API_INTROSPECT_PORT=8084 \
    -e LOG_LEVEL=SYS_NOTICE \
    -e log_local=true \
    -e AUTH_MODE=none \
    -e AAA_MODE=cloud-admin \
    -e ZOOKEEPER_SERVERS="$ServicesIpAddress:2181" \
    -e CONFIGDB_SERVERS="$ServicesIpAddress:9161" \
    -e RABBITMQ_SERVERS="$ServicesIpAddress:5673" \
    vncdbproxy
