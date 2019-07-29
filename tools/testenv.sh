#!/bin/bash

set -e

SOURCEDIR=$( cd "$(dirname "$0")/../../../../.." ; pwd -P )
TOOLSDIR=$(dirname $0)

RunDockers="etcd patroni"
Network='contrail'
PROJECT='contrail'
PatroniEtcd=0

Usage()
{
	echo "Usage: $(basename "$0") [-h] [-n NetName] [--patroni-etcd] [dockers]"
	echo "Available dockers: $RunDockers"
}

while :; do
	case "$1" in
		'-n') Network="$2"; shift 2;;
		'--patroni-etcd') PatroniEtcd=1; shift 1;;
		'-h') Usage; exit 0;;
		*) break;;
	esac
done

PASSWORD=contrail123
SpecialNetworks='bridge none host'
[[ "$SpecialNetworks" = *"$Network"* ]] || PROJECT=$Network
NETWORKNAME=$PROJECT docker-compose -f "$TOOLSDIR/patroni/docker-compose.yml" -p $PROJECT down -v || true
docker rm -f contrail_etcd || true
docker volume rm -f contrail_etcd || true
docker network remove $PROJECT || true

docker network create $PROJECT --subnet 10.0.4.0/24 --gateway 10.0.4.1 || true

run_docker_patroni()
{
    if [[ "$PatroniEtcd" = 1 ]]; then
        NETWORKNAME=$PROJECT docker-compose -f "$TOOLSDIR/patroni/docker-compose.yml" -p $PROJECT up -d etcd
        ETCDIP=$(docker inspect contrail_patroni_etcd --format='{{ range .NetworkSettings.Networks }}{{.IPAddress}}{{end}}')
    else
        ETCDIP=$(docker inspect contrail_etcd --format='{{ range .NetworkSettings.Networks }}{{.IPAddress}}{{end}}')
    fi
    ETCDIP=$ETCDIP NETWORKNAME=$PROJECT docker-compose -f "$TOOLSDIR/patroni/docker-compose.yml" -p $PROJECT up --scale dbnode=2 -d haproxy dbnode
}

run_docker_etcd()
{
	docker run -d --name contrail_etcd \
		--net "$Network" \
		-p 2379:2379 \
		-e "ETCDCTL_API=3" \
		-v "contrail_etcd:/etcd-data" \
		gcr.io/etcd-development/etcd:v3.3.2 \
		etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379
}

[ ! -z "$1" ] && RunDockers="$*"

for docker in $RunDockers; do
	eval "run_docker_$docker"
done

echo "TestEnv done"
