#!/bin/bash

create_group()
{
	sudo groupadd "$1"
	sudo usermod -aG "$1" "$USER"
}
ensure_group()
{
	local expected_group='docker'
	groups | grep -q "$expected_group" || create_group "$expected_group"
	if [ "$(id -gn)" != "$expected_group" ]; then
		exec sg "$expected_group" -c "$0 $*"
	fi
}
ensure_group "$@"

RealPath()
{
	pushd "$1" &> /dev/null
	pwd
	popd &> /dev/null
}

ThisDir=$(RealPath "$(dirname "$0")")
RootDir=$(RealPath "$ThisDir/..")
PORT=8082

build_docker()
{
	dir=$(pwd)
	cd "$RootDir"
	make docker_k8s
	cd "$dir"
}

set -e
set -x

install_golang()
{
	cd /tmp
	curl -o go.tar.gz https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
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
[ "$GOPATH/src/github.com/Juniper/contrail" = "$RootDir" ] || { echo "This repo should be clonned into GOPATH == $GOPATH"; exit 2; }
echo "$PATH" | grep -q "$GOPATH/bin" || export PATH="$PATH:$GOPATH/bin"

cd "$RootDir"
make deps
make generate
make build
make install
# etcd should be already deployed with kubernetes
"$ThisDir/testenv.sh" -n host postgres

# Stop kubemanager, original config-node, control-node and vrouter
docker-compose -f /etc/contrail/kubemanager/docker-compose.yaml down
docker-compose -f /etc/contrail/config/docker-compose.yaml down
docker-compose -f /etc/contrail/control/docker-compose.yaml down
docker-compose -f /etc/contrail/vrouter/docker-compose.yaml down

# Clear old config-node databases
docker-compose -f /etc/contrail/config_database/docker-compose.yaml stop cassandra rabbitmq zookeeper
docker-compose -f /etc/contrail/config_database/docker-compose.yaml rm -v
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
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype etcd -c docker/contrail_go/etc/contrail-k8s.yml
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype http -u http://127.0.0.1:9082 || true

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
