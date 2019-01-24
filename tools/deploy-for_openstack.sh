#!/bin/bash

set -o errexit
set -o pipefail
set -o xtrace

RealPath()
{
    pushd "$1" &> /dev/null
    pwd
    popd &> /dev/null
}

ThisDir=$(RealPath "$(dirname "$0")")
#shellcheck disable=SC1090
source "$ThisDir/ensure_docker_group.sh"

ensure_group "$@"

source "$ThisDir/deploy-utils.sh"
source "$ThisDir/ensure_golang_installed.sh"

ensure_golang_installed

ContrailRootDir=$(RealPath "$ThisDir/..")
cd "$ContrailRootDir"
make deps
make generate
make build
make install

make docker_config_api

# Ensure patroni installed
./tools/patroni/pull_patroni.sh

./tools/testenv.sh etcd patroni

# Stop services using docker-compose
compose --down config control vrouter

# Clear old config-node databases
clear_config_database

# Prepare fresh database in contrail-go
make zero_psql

# Update control-node docker compose file
sudo ./tools/control-node_etcd/update-docker-compose.py

# Update schema transformer docker compose file
sudo ./tools/schema_transformer_etcd/update-docker-compose.py

# Update svnmonitor docker compose file
sudo ./tools/svcmonitor_etcd/update-docker-compose.py

# Load init data to rdbms
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail-config_api.yml

build_and_run_contrail-go_docker

# Start services using docker-compose
compose --up control vrouter config:schema config:svcmonitor
