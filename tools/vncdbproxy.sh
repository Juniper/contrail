#!/bin/bash

if [ "$1" == "" ]; then
    echo "Usage: vncdbproxy.sh [-n <NetworkMode>] [CONFIGDBSERVER_IP]"
    exit 1
fi

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
[ $RemoveDockers -eq 1 ] && docker rm -f vncdbproxy

docker build ../docker/vnc_db_proxy/ -t vncdbproxy

docker run \
    --name vncdbproxy \
    --net "$Network"
    -p 8091:8082 \
    -d \
    -e CONFIG_API_PORT=8082 \
    -e CONFIG_API_INTROSPECT_PORT=8084 \
    -e LOG_LEVEL=SYS_NOTICE \
    -e log_local=true \
    -e AUTH_MODE=none \
    -e AAA_MODE=cloud-admin \
    -e ZOOKEEPER_SERVERS="$1:2181" \
    -e CONFIGDB_SERVERS="$1:9161" \
    -e RABBITMQ_SERVERS="$1:5673" \
    vncdbproxy
