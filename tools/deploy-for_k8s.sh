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
cp "$ContrailRootDir/tools/init_psql.sql" "$ContrailRootDir/tools/patroni/k8s/"
"$ContrailRootDir/tools/patroni/k8s/install_patroni_k8s.sh"
sudo "$ContrailRootDir/tools/patroni/k8s/start_cluster.sh"

# Stop services using docker-compose
compose_down kubemanager config control vrouter

# Clear old config-node databases
clear_config_database

# Prepare fresh database in contrail-go
"$ContrailRootDir/tools/reset_db_psql_k8s.sh"

# Drop contrail related content from etcd
docker exec "$(docker ps -q -f name=k8s_etcd_etcd)" sh -c "ETCDCTL_API=3 etcdctl del /contrail --prefix"

# Build patched kube_manager with etcd support
docker build "$ContrailRootDir/docker/kube_manager_etcd/" -t contrail-kubernetes-kube-manager:etcd

# Update kube_manager docker compose file
sudo ./tools/kube_manager_etcd/update-docker-compose.py

# Update control-node docker compose file
sudo ./tools/control-node_etcd/update-docker-compose.py

# Get IP address of cluster
CLIP=$(sudo kubectl describe svc -l application=contrail-postgres,cluster-name=contrail | grep IP: | sed "s/IP:[ ]*//g")

# TODO: substitute address in config

# Load init data to rdbms
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail-config_api.yml

build_and_run_contrail-go_docker

GoConfigIP='127.0.0.1' # networking mode 'host'
ensure_kubemanager_config_nodes "${GoConfigIP}"

# Start services using docker-compose
compose_up control vrouter kubemanager
