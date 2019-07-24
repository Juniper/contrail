#!/bin/bash

set -o errexit

ToolsDir=$(dirname "$0")
RunDockers="etcd patroni"
Network='contrail'
Project='contrail'
PatroniEtcd=0

usage()
{
	echo "Usage: $(basename "$0") [-h] [-n NetName] [--patroni-etcd] [dockers]"
	echo "Available dockers: $RunDockers"
}

while :; do
	case "$1" in
		'-n') Network="$2"; shift 2;;
		'--patroni-etcd') PatroniEtcd=1; shift 1;;
		'-h') usage; exit 0;;
		*) break;;
	esac
done

run() {
    SpecialNetworks='bridge none host'
    [[ "$SpecialNetworks" = *"$Network"* ]] || Project=${Network}

    remove_containers
    create_containers "$@"
    await_psql
}

remove_containers() {
    NETWORKNAME="$Project" docker-compose -f "$ToolsDir/patroni/docker-compose.yml" -p "$Project" down -v || true
    docker rm -f contrail_etcd || true
    docker volume rm -f contrail_etcd || true
    docker network remove "$Project" || true
}

create_containers() {
    docker network create "$Project" --subnet 10.0.4.0/24 --gateway 10.0.4.1 || true

    [[ -n "$1" ]] && RunDockers="$*"

    for docker in ${RunDockers}; do
        eval "run_docker_${docker}"
    done
}

run_docker_patroni() {
    if [[ "$PatroniEtcd" = 1 ]]; then
        NETWORKNAME="$Project" docker-compose -f "$ToolsDir/patroni/docker-compose.yml" -p "$Project" up -d etcd
        ETCDIP=$(docker inspect contrail_patroni_etcd --format='{{ range .NetworkSettings.Networks }}{{.IPAddress}}{{end}}')
    else
        ETCDIP=$(docker inspect contrail_etcd --format='{{ range .NetworkSettings.Networks }}{{.IPAddress}}{{end}}')
    fi

    ETCDIP=${ETCDIP} NETWORKNAME="$Project" docker-compose -f "$ToolsDir/patroni/docker-compose.yml" -p "$Project" up \
        --scale dbnode=2 -d haproxy dbnode
}

run_docker_etcd() {
	docker run -d --name contrail_etcd \
		--net "$Network" \
		-p 2379:2379 \
		-e "ETCDCTL_API=3" \
		-v "contrail_etcd:/etcd-data" \
		gcr.io/etcd-development/etcd:v3.3.2 \
		etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379
}

await_psql() {
    echo "Awaiting PostgreSQL"

    until docker exec contrail_haproxy patronictl list | grep 'Leader.*running'
    do
        printf "."
        sleep 1
    done
}

run "$@"
