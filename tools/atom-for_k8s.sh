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

build_atom()
{
	dir=$(pwd)
	cd "$RootDir"
	make docker-atomizer
	cd "$dir"
}

set -e
set -x

install_golang()
{
	cd /tmp
	curl -o go.tar.gz https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
	sudo tar --overwrite -C /usr -xzf go.tar.gz
	export PATH="$PATH:/usr/go/bin"
	hash -r
	go env
	# wget and unzip are needed for `make deps` assuming clean system if golang neede to be installed
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
# etcd should be already deployed with kubernetes
"$ThisDir/testenv.sh" postgres

#Stop kubemanager ?
cd /etc/contrail/kubemanager
docker-compose down
cd "$RootDir"

# Dump cassandra from orig config-node
# TODO

# Build run Atom docker and get Atom IP address
build_atom
AtomizerDocker='atom-contrail'
[ "$(docker ps -a -f "name=$AtomizerDocker" --format '{{.ID}}' | wc -l)" -ne 0 ] && docker rm "$AtomizerDocker"
docker run -d --name "$AtomizerDocker" -p "$PORT:$PORT" --net host contrail-atom
docker logs -f "$AtomizerDocker"
#AtomIP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} {{.Name}}' "$AtomizerDocker")
AtomIP='127.0.0.1' # networking mode 'host'

# Convert cassandra data to etcd and feed etcd
# TODO

# Modify k8s config (subst Atom IP address as config-node) and restart if needed
ModifyKubeConfig=1
grep -qE "^CONFIG_NODES\\W*=\\W*$AtomIP" /etc/contrail/common_kubemanager.env && ModifyKubeConfig=0
if [ $ModifyKubeConfig -eq 1 ]; then
	sudo sed "-ibak$(date +%s)" "s/^CONFIG_NODES=.*/CONFIG_NODES=$AtomIP/" /etc/contrail/common_kubemanager.env
	cd /etc/contrail/kubemanager
	docker-compose up -d
fi

