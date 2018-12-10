#!/bin/bash

set -e

SOURCEDIR=$( cd "$(dirname "$0")/../../../../.." ; pwd -P )
TOOLSDIR=$(dirname $0)

RunDockers="mysql etcd patroni"
Network='contrail'
PROJECT='contrail'

Usage()
{
	echo "Usage: $(basename "$0") [-h] [-n NetName] [dockers]"
	echo "Available dockers: $RunDockers"
}

while :; do
	case "$1" in
		'-n') Network="$2"; shift 2;;
		'-h') Usage; exit 0;;
		*) break;;
	esac
done

PASSWORD=contrail123
SpecialNetworks='bridge none host'
docker network create $PROJECT --subnet 10.0.4.0/24 --gateway 10.0.4.1 || true
HOSTIP=$HOSTGATE NETWORKNAME=$PROJECT docker-compose -f "$TOOLSDIR/patroni/docker-compose.yml" -p $PROJECT down || true
docker rm -f contrail_mysql contrail_etcd || true

run_docker_patroni()
{
    if [[ "$Network" = "host" ]]; then
        HOSTIP=$HOSTGATE NETWORKNAME=$PROJECT docker-compose -f "$TOOLSDIR/patroni/docker-compose.yml" -p $PROJECT up -d etcd
        HOSTGATE=$(docker inspect contrail_etcd --format='{{ range .NetworkSettings.Networks }}{{.IPAddress}}{{end}}')
    else
        HOSTGATE=$(docker network inspect $PROJECT --format='{{ range .IPAM.Config }}{{ .Gateway }}{{ end }}')
    fi
    HOSTIP=$HOSTGATE NETWORKNAME=$PROJECT docker-compose -f "$TOOLSDIR/patroni/docker-compose.yml" -p $PROJECT up --scale dbnode=2 -d haproxy dbnode
}

run_docker_mysql()
{
	docker run -d --name contrail_mysql \
		--net "$Network" \
		-v "$SOURCEDIR:/go" \
		-p 3306:3306 \
		-e "MYSQL_ROOT_PASSWORD=$PASSWORD" \
		circleci/mysql:5.7
}

run_docker_etcd()
{
	docker run -d --name contrail_etcd \
		--net "$Network" \
		-p 2379:2379 \
		-e "ETCDCTL_API=3" \
		gcr.io/etcd-development/etcd:v3.3.2 \
		etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379
}

[ ! -z "$1" ] && RunDockers="$*"

WaitMysql=0
for docker in $RunDockers; do
	eval "run_docker_$docker"
	[ "$docker" = mysql ] && WaitMysql=1
done

if [ $WaitMysql -eq 1 ]; then
	echo "Waiting for mysql"
	until docker exec contrail_mysql mysql -uroot -p"$PASSWORD" -e "show status" &> /dev/null
	do
		printf "."
		sleep 1
	done
fi

echo "TestEnv done"
