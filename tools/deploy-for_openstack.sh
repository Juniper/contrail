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
make docker

# Ensure patroni installed
./tools/patroni/pull_patroni.sh

./tools/testenv.sh etcd patroni

# Stop services using docker-compose
compose_down config control vrouter

# Clear old config-node databases
clear_config_database

# Prepare fresh database in contrail-go
make zero_psql

# Drop etcd content
docker exec "$(docker ps -q -f name=contrail_etcd)" sh -c "ETCDCTL_API=3 etcdctl del / --prefix"

# Ensure keystone is listening on localhost
ensure_keystone_on_localhost

# Update control-node docker compose file
sudo ./tools/control-node_etcd/update-docker-compose.py

# Update schema transformer docker compose file
sudo ./tools/schema_transformer_etcd/update-docker-compose.py

# Update device manager docker compose file
sudo ./tools/device_manager_etcd/update-docker-compose.py

# Update svnmonitor docker compose file
sudo ./tools/svcmonitor_etcd/update-docker-compose.py

# Load init data to rdbms
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail-openstack.yml

install_config "contrail-openstack"

build_and_run_contrail-go_docker

# Start services using docker-compose
compose_up config:nodemgr config:schema config:svcmonitor config:devicemgr control vrouter
