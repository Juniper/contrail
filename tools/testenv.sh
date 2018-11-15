#!/bin/bash

set -e

SOURCEDIR=$( cd "$(dirname "$0")/../../../../.." ; pwd -P )

RunDockers="mysql etcd"
Network='contrail'

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
[[ "$SpecialNetworks" = *"$Network"* ]] || docker network create contrail || true
docker-compose -f ./tools/patroni/docker-compose.yml -p "contrail" down || true
docker rm -f contrail_mysql contrail_etcd || true

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


PT="src/github.com/Juniper/contrail"
echo "Files inside $PWD" && ls -al
echo "Files inside /go/$PT/tools/patroni" && ls -al /go/$PT/tools/patroni
echo "Files inside $PWD/tools" && ls -al $PWD/tools
echo "Files inside /home/zuul/$PT/tools" && ls -al /home/zuul/$PT/tools
echo "Files inside /home/zuul/" && ls -al /home/zuul/
echo "Files inside /home/zuul/$PT/tools/patroni" && ls -al /home/zuul/$PT/tools/patroni

docker-compose -f ./tools/patroni/docker-compose.yml -p "contrail" up --scale dbnode=2 -d || { echo "Files inside tools/patroni" ; ls -al ./tools/patroni ;
                                                                                               echo "Files inside $PWD" ; ls -al ;
                                                                                               exit 1 ; }

if [ $WaitMysql -eq 1 ]; then
	echo "Waiting for mysql"
	until docker exec contrail_mysql mysql -uroot -p"$PASSWORD" -e "show status" &> /dev/null
	do
		printf "."
		sleep 1
	done
fi

echo "TestEnv done"
