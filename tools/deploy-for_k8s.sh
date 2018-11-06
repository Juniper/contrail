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
. "$ThisDir/ensure_docker_group.sh"

ensure_group "$@"

ContrailRootDir=$(RealPath "$ThisDir/..")

build_docker()
{
	dir=$(pwd)
	cd "$ContrailRootDir"
	make docker_k8s
	cd "$dir"
}

install_golang()
{
	cd /tmp
	curl -o go.tar.gz https://dl.google.com/go/go1.11.1.linux-amd64.tar.gz
	sudo tar --overwrite -C /usr -xzf go.tar.gz
	sudo yum install -y wget unzip
	export PATH="$PATH:/usr/go/bin"
	hash -r
	go env
}
if [ -d /usr/go/bin ]; then
	echo "$PATH" | grep -q /usr/go/bin || export PATH="$PATH:/usr/go/bin"
fi
go env || install_golang
[ -z "$GOPATH" ] && export GOPATH="$HOME/go"
echo "$PATH" | grep -q "$GOPATH/bin" || export PATH="$PATH:$GOPATH/bin"

cd "$ContrailRootDir"
make deps
make generate
make build
make install
# etcd should be already deployed with kubernetes
"$ContrailRootDir/tools/testenv.sh" -n host postgres

# Stop kubemanager, original config-node, control-node and vrouter
docker-compose -f /etc/contrail/kubemanager/docker-compose.yaml down
docker-compose -f /etc/contrail/config/docker-compose.yaml down
docker-compose -f /etc/contrail/control/docker-compose.yaml down
docker-compose -f /etc/contrail/vrouter/docker-compose.yaml down

# Clear old config-node databases
docker-compose -f /etc/contrail/config_database/docker-compose.yaml down -v
docker-compose -f /etc/contrail/config_database/docker-compose.yaml up -d zookeeper

# Prepare fresh database in contrail-go
make zero_psql

# Build patched kube_manager with ETCD support
docker build "$ContrailRootDir/docker/kube_manager_etcd/" -t contrail-kubernetes-kube-manager:etcd

# Update kube_manager docker compose file
sudo ./tools/kube_manager_etcd/update-docker-compose.py

# Update control-node docker compose file
sudo ./tools/control-node_etcd/update-docker-compose.py

# Load init data to rdbms
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c docker/contrail_go/etc/contrail-k8s.yml

# Build and run contrail-go2 docker
build_docker
ContrailGoDocker='contrail-go-config-node'
[ "$(docker ps -a -f "name=$ContrailGoDocker" --format '{{.ID}}' | wc -l)" -ne 0 ] && docker rm -f "$ContrailGoDocker"
docker run -d --name "$ContrailGoDocker" --net host contrail-go-config
GoConfigIP='127.0.0.1' # networking mode 'host'

# Modify k8s config (subst contrail-go-config IP address as config-node) and restart if needed
ModifyKubeConfig=1
grep -qE "^CONFIG_NODES\\W*=\\W*$GoConfigIP" /etc/contrail/common_kubemanager.env && ModifyKubeConfig=0
if [ $ModifyKubeConfig -eq 1 ]; then
	sudo sed "-ibak$(date +%s)" "s/^CONFIG_NODES=.*/CONFIG_NODES=$GoConfigIP/" /etc/contrail/common_kubemanager.env
fi

# Start control-node, vrouter and kubemanager
docker-compose -f /etc/contrail/control/docker-compose.yaml up -d
docker-compose -f /etc/contrail/vrouter/docker-compose.yaml up -d
docker-compose -f /etc/contrail/kubemanager/docker-compose.yaml up -d
