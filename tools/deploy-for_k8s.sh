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

./tools/testenv.sh --patroni-etcd patroni

# Stop services using docker-compose
compose_down kubemanager config control vrouter

# Clear old config-node databases
clear_config_database

# Prepare fresh database in contrail-go
make zero_psql

# Drop contrail related content from etcd
docker exec "$(docker ps -q -f name=k8s_etcd_etcd)" sh -c "ETCDCTL_API=3 etcdctl del /contrail --prefix"

# Update kube_manager docker compose file
sudo ./tools/kube_manager_etcd/update-docker-compose.py

# Update control-node docker compose file
sudo ./tools/control-node_etcd/update-docker-compose.py

# Update schema transformer docker compose file
sudo ./tools/schema_transformer_etcd/update-docker-compose.py

# Update device manager docker compose file
sudo ./tools/device_manager_etcd/update-docker-compose.py

# Load init data to rdbms
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail-k8s.yml

install_config "contrail-k8s"

build_and_run_contrail-go_docker

GoConfigIP='127.0.0.1' # networking mode 'host'
ensure_kubemanager_config_nodes "${GoConfigIP}"

# Start schema transformer
schema_transformer_up

# Start device manager
device_manager_up

# Start services using docker-compose
compose_up control vrouter kubemanager
