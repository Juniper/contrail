#!/bin/sh

ThisDir=$(cd "$(dirname "$0")"; pwd -P)
SOURCEDIR=$( cd "$(dirname "$0")/../../../../.." ; pwd -P )
PORT=8082

build_atom()
{
	dir=$(pwd)
	cd "$SOURCEDIR"
	make docker-atomizer
	cd "$dir"
}

set -e
set -x

install_golang()
{
	cd /tmp
	curl -o go.tar.gz https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
	sudo tar -C /usr -xzf go.tar.gz
	go env
	cd "$SOURCEDIR"
	make deps
}

go env || install_golang
[ -n "$GOPATH" ] && export GOPATH="$HOME/go"
[ "$GOPATH/src/github.com/Juniper/contrail" = "$SOURCEDIR" ] || { echo "This repo should be clonned into GOPATH == $GOPATH"; exit 2; }

cd "$SOURCEDIR"
make generate
make build
"$ThisDir/testenv.sh" postgres etcd

#Stop kubemanager ?
# TODO

# Dump cassandra from orig config-node
# TODO

# Build run Atom docker and get Atom IP address
build_atom
AtomizerDocker='atom-contrail'
[ "$(docker ps -a -f "name=$AtomizerDocker" --format '{{.ID}}' | wc -l)" -ne 0 ] && docker rm "$AtomizerDocker"
docker run -d --name "$AtomizerDocker" -p "$PORT:$PORT" \
	--link contrail_etcd:etcd \
	--link contrail_postgres:postgres \
	contrail-atomizer
docker logs -f "$AtomizerDocker"
AtomIP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} {{.Name}}' "$AtomizerDocker")

# Modify k8s config (subst Atom IP address as config-node) and restart if needed
ModifyKubeConfig=1
grep -qE "^CONFIG_NODES\\W*=\\W*$AtomIP" && ModifyKubeConfig=0
if [ $ModifyKubeConfig -eq 1 ]; then
	sed "-ibak$(date +%s)" "s/^CONFIG_NODES=.*/CONFIG_NODES=$AtomIP/" /etc/contrail/common_kubemanager.env
	cd /etc/contrail/kubemanager
	docker-compose down
	docker-compose up
fi

# Convert cassandra data to etcd and feed etcd
# TODO

# Start kubemanager

