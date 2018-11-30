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
"$ContrailRootDir/tools/testenv.sh" -n host postgres etcd

# Stop original config-node, control-node and vrouter
docker-compose -f /etc/contrail/config/docker-compose.yaml down
docker-compose -f /etc/contrail/control/docker-compose.yaml down
docker-compose -f /etc/contrail/vrouter/docker-compose.yaml down

# Clear old config-node databases
docker-compose -f /etc/contrail/config_database/docker-compose.yaml down -v
docker-compose -f /etc/contrail/config_database/docker-compose.yaml up -d

# Prepare fresh database in contrail-go
make zero_psql

# Run vnc-db-proxy
./tools/vncdbproxy/vncdbproxy.sh -n host -z localhost:2181 -c localhost:9161 -r localhost:5673

# Wait for vnc-db-proxy
until $(curl --output /dev/null --silent --head --fail http://localhost:9082); do
    printf '.'
    sleep 5
done

# Load init data to new and legacy databases
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c docker/contrail_go/etc/contrail-k8s.yml
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype http -u http://127.0.0.1:9082 || true

# Build and run contrail-go2 docker
build_docker
ContrailGoDocker='contrail-go-config-node'
[ "$(docker ps -a -f "name=$ContrailGoDocker" --format '{{.ID}}' | wc -l)" -ne 0 ] && docker rm -f "$ContrailGoDocker"
docker run -d --name "$ContrailGoDocker" --net host contrail-go-config

# Start control-node, vrouter
docker-compose -f /etc/contrail/control/docker-compose.yaml up -d
docker-compose -f /etc/contrail/vrouter/docker-compose.yaml up -d
