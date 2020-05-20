#!/bin/bash

set -o errexit

ToolsDir=$(cd $(dirname "$0"); pwd -P)
RunDockers="etcd postgres"
Network='contrail'
Project='contrail'

usage()
{
	echo "Usage: $(basename "$0") [-h] [-n NetName] [dockers]"
	echo "Available dockers: $RunDockers"
}

while :; do
	case "$1" in
		'-n') Network="$2"; shift 2;;
		'-h') usage; exit 0;;
		*) break;;
	esac
done

run() {
    SpecialNetworks='bridge none host'
    [[ "$SpecialNetworks" = *"$Network"* ]] || Project="$Network"

    remove_containers
    run_containers "$@"
    await_psql
}

remove_containers() {
    docker rm -f contrail_psql || true
    docker rm -f contrail_etcd || true
    docker volume rm -f contrail_etcd || true
    docker network remove "$Project" || true
}

run_containers() {
    docker network create "$Project" --subnet 10.0.4.0/24 --gateway 10.0.4.1 || true

    [[ -n "$1" ]] && RunDockers="$*"

    for docker in ${RunDockers}; do
        eval "run_docker_${docker}"
    done
}

run_docker_postgres() {
    docker run -d --name contrail_psql \
            --net "$Network" \
            -p 5432:5432 \
            -e POSTGRES_PASSWORD=contrail123 \
            -e POSTGRES_USER=root \
            -e POSTGRES_DB=contrail_test \
            -v "$ToolsDir:/tools" \
            postgres:10 \
            postgres -c wal_level=logical
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

    until docker exec contrail_psql psql -Uroot -d postgres -c 'SHOW server_version' 2>/dev/null
    do
        printf "."
        sleep 1
    done
}

run "$@"
