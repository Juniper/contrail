#!/bin/bash

set -x

RealPath()
{
	pushd "$1" &> /dev/null
	pwd
	popd &> /dev/null
}

ThisDir=$(RealPath "$(dirname "$0")")
. "$ThisDir/ensure_docker_group.sh"

ensure_group "$@"

Usage()
{
	echo -e "Usage: $(basename "$0") [-D]\n"
	echo 'This script is intended to be used on a node with contrail is installed '
	echo 'with contrail-ansible-deployer and later tools/deploy-for_k8s.sh is used'
	echo 'This script will run sanity tests defined in https://github.com/Juniper/contrail-test'
	echo
	echo -e "\t-A => Don't check contrail-go dockers before tests"
}

CheckDockers=1
while :; do
	case "$1" in
		'-D') CheckDockers=0; shift;;
		'-h') Usage; exit 0;;
		*) break;;
	esac
done
[ $# -ne 0 ] && { Usage; exit 1; }

InstancesFile="$HOME/contrail-ansible-deployer/config/instances.yaml"
ThisUser=$(id -nu)

check_dockers()
{
	[ $CheckDockers -eq 0 ] && { echo "WARNING: Skipping contrail-go dockers checking!"; return 0; }
	dockers='contrail-go-config-node contrail_postgres'
	for d in $dockers; do
		docker ps | grep -qE "Up.*$d\$" || { echo "Expected docker $d to be running"; exit 2; }
	done
}

ensure_root_access()
{
	key="$HOME/.ssh/id_rsa"
	[ ! -e "$key" ] && ssh-keygen -N '' -f "$key"
	auth_entry="$ThisUser@$(hostname)"
	sudo grep -q "$auth_entry" /root/.ssh/authorized_keys || sudo sh -c "cat \"$HOME/.ssh/id_rsa.pub\" >> /root/.ssh/authorized_keys"
}

prepare_test_env()
{
	cd
	[ ! -e ./testrunner.sh ] && wget https://github.com/Juniper/contrail-test/raw/master/testrunner.sh
	sed -i '/^\s\+tput/ d' ./testrunner.sh
	chmod +x ./testrunner.sh
	[ ! -e contrail_test_input.yaml.sample ] && {
		wget https://github.com/Juniper/contrail-test/raw/master/contrail_test_input.yaml.sample && \
		cp contrail_test_input.yaml.sample contrail_test_input.yaml; }
	docker pull opencontrailnightly/contrail-test-test:latest
	ensure_root_access
	grep -q 'orchestrator: kubernetes' "$InstancesFile" || cat >> "$InstancesFile" <<EOF
deployment:
  orchestrator: kubernetes
EOF

}

run_contrail_test()
{
	echo 'Running testrunner...'
	echo
	./testrunner.sh run -P "$InstancesFile" -c test_many_pods opencontrailnightly/contrail-test-test:latest
}

finalize_test_run()
{
	cd ~/contrail-test-runs || return 0
	sudo chown -R "$ThisUser:$ThisUser" ./
	for dir in $(find ./ -maxdepth 1 -type d | cut -f 2 -d '/'); do
		echo "$dir" | grep -qE '^[0-9_]+$' || continue # skip non-report directories (if any)
		archive="$dir.tgz"
		[ -e "$archive" ] && continue
		tar -zcf "$archive" "$dir"
		ls -al "$archive"
	done
}

check_dockers
prepare_test_env
run_contrail_test
finalize_test_run
