#!/bin/bash

set -e

SOURCEDIR=$( cd "$(dirname "$0")/../../../../.." ; pwd -P )

Network='contrail'
while :; do
	case "$1" in
		'-n') Network="$2"; shift 2;;
		*) break;;
	esac
done

PASSWORD=contrail123
SpecialNetworks='bridge none host'
[[ "$SpecialNetworks" = *"$Network"* ]] || docker network create contrail || true
docker rm -f contrail_postgres contrail_mysql contrail_etcd || true

run_docker_postgres()
{
	docker run -d --name contrail_postgres \
		--net "$Network" \
		-v "$SOURCEDIR:/go" \
		-p 5432:5432 \
		-e "POSTGRES_USER=root" \
		-e "POSTGRES_PASSWORD=$PASSWORD" \
		circleci/postgres:10.3-alpine -c 'wal_level=logical'
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
		gcr.io/etcd-development/etcd:v3.3.2 \
		etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379
}

RunDockers="postgres mysql etcd"
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
