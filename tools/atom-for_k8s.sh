#!/bin/sh

ThisDir=$(cd $(dirname "$0"); pwd -P)
SOURCEDIR=$( cd "$(dirname "$0")/../../../../.." ; pwd -P )

build_atom()
{
	dir=$(pwd)
	cd "$SOURCEDIR"
	make docker-atomizer
	cd "$dir"
}

set -e

"$ThisDir/configapi.sh"
"$ThisDir/testenv.sh" postgres etcd
build_atom
AtomIP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} {{.Name}}' contrail-atom)
sed "-ibak$(date +%s)" "s/^CONFIG_NODES=.*/CONFIG_NODES=$AtomIP/" /etc/contrail/common_kubemanager.env
