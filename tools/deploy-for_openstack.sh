#!/bin/bash

set -o errexit
set -o pipefail
set -o xtrace

ThisDir=$(realpath "$(dirname "$0")")
source "$ThisDir/ensure_docker_group.sh"

ensure_group "$@"

source "$ThisDir/deploy-utils.sh"
source "$ThisDir/ensure_golang_installed.sh"

ensure_golang_installed

ContrailRootDir=$(realpath "$ThisDir/..")
cd "$ContrailRootDir"
make deps
make generate
make build
make install

make docker_k8s

# Ensure patroni installed
./tools/patroni/install_patroni.sh
./tools/testenv.sh -n host patroni etcd

# Stop services using docker-compose
compose_down config control vrouter

# Dump config-node database
contrailutil convert --intype cassandra --in localhost -p 9041 --outtype yaml --out db_dump.yaml

clear_config_database

# Prepare fresh database in contrail-go
make zero_psql

# Update control-node docker compose file
sudo ./tools/control-node_etcd/update-docker-compose.py

# Load init data to rdbms
contrailutil convert --intype yaml --in ./db_dump.yaml --outtype rdbms -c ./docker/contrail_go/etc/contrail-k8s.yml

build_and_run_contrail-go_docker

# Start services using docker-compose
compose_up control vrouter
