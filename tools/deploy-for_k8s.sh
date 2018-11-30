#!/bin/bash

set -o errexit
set -o pipefail
set -o xtrace

ThisDir=$(realpath "$(dirname "$0")")
source "$ThisDir/ensure_docker_group.sh"

ensure_group "$@"

source "$ThisDir/deploy-utils.sh"

ensure_golang_installed

ContrailRootDir=$(realpath "$ThisDir/..")
cd "$ContrailRootDir"
make docker_k8s
make deps
make generate
make build
make install

# etcd should be already deployed with kubernetes
"$ThisDir/patroni/install_patroni.sh"
"$ThisDir/testenv.sh" -n bridge patroni

stop-dockers kubemanager config control vrouter

# Clear old config-node databases
clear_config_database

# Prepare fresh database in contrail-go
make zero_psql

# Drop contrail related content from etcd
docker exec $(docker ps -q -f name=k8s_etcd_etcd) sh -c "ETCDCTL_API=3 etcdctl del /contrail --prefix"

# Build patched kube_manager with ETCD support
docker build "$ContrailRootDir/docker/kube_manager_etcd/" -t contrail-kubernetes-kube-manager:etcd

# Update kube_manager docker compose file
sudo ./tools/kube_manager_etcd/update-docker-compose.py

# Update control-node docker compose file
sudo ./tools/control-node_etcd/update-docker-compose.py

# Load init data to rdbms
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c docker/contrail_go/etc/contrail-k8s.yml

build_and_run_contrail-go_docker

GoConfigIP='127.0.0.1' # networking mode 'host'
ensure_kubemanager_config_nodes "${GoConfigIP}"

start-dockers control vrouter kubemanager
