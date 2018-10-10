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
	echo -e "Usage: $(basename "$0") [-h] [-D] [-t <test_spec> ...]\n"
	echo 'This script is intended to be used on a node with contrail is installed '
	echo 'with contrail-ansible-deployer and later tools/deploy-for_k8s.sh is used'
	echo 'This script will run sanity tests defined in https://github.com/Juniper/contrail-test'
	echo
	echo -e "\t-D => Don't check contrail-go dockers before tests"
	echo -e "\t-t => Test cases/tags to be run by contrail-test, specified as contrail-test flags (eg. \`-t -c test_many_pods -T k8s_sanity\`"
	echo -e "\t\ttest_spec is (-c|-t|-f name)"
}

declare -a TestsTypes
declare -a TestsToRun
GatherTests()
{
	shiftem=1
	n=0
	shift
	while :; do
		case "$1" in
			-c | -T | -f) ttype="$1"; name="$2"; shift 2; shiftem=$((shiftem+2));;
			*) Usage; exit 1;;
		esac
		[ -z "$name" ] && { Usage; exit 1; }
		TestsTypes[$n]="$ttype"
		TestsToRun[$n]="$name"
		n=$((n+1))
	done
	return $shiftem
}

CheckDockers=1
while :; do
	case "$1" in
		'-D') CheckDockers=0; shift;;
		'-t') GatherTests "$@"; shift $?;;
		'-h') Usage; exit 0;;
		*) break;;
	esac
done
[ $# -ne 0 ] && { Usage; exit 1; }

RealPath()
{
	pushd "$1" &> /dev/null
	pwd
	popd &> /dev/null
}

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
	yq -y 'with_entries(select(.key == "provider_config")) | with_entries(.value[] +={ "ssh_user": "root", "ssh_pwd": "contrail123"})' "$InstancesFile" > /tmp/instances-provider.yaml
	yq -y 'with_entries(select(.key != "provider_config"))' "$InstancesFile" > /tmp/instances-rest.yaml
	cp -f /tmp/instances-provider.yaml "$InstancesFile"
	cat /tmp/instances-rest.yaml >> "$InstancesFile"
	modify_sshd=0
	grep '^PasswordAuthentication yes$' /etc/ssh/sshd_config || modify_sshd=1
	grep '^PermitRootLogin yes$' /etc/ssh/sshd_config || modify_sshd=1
	if [ $modify_sshd -eq 1 ]; then
		cat >> /etc/ssh/sshd_config <<EOF
PasswordAuthentication yes
PermitRootLogin yes
EOF
	sudo service sshd restart
	fi
}

prepare_test_env()
{
	cd
	[ ! -e ./testrunner.sh ] && wget https://github.com/Juniper/contrail-test/raw/master/testrunner.sh
	sed -i '/^\s\+tput/ d' ./testrunner.sh
	chmod +x ./testrunner.sh
	docker pull opencontrailnightly/contrail-test-test:latest || { echo 'Fail to pull docker image'; exit 2; }
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
	if [ 0 -eq ${#TestsTypes[@]} ]; then
		TestsTypes=('-c' '-c' '-c')
		TestsToRun=('test_namespace_1' 'test_ping_between_two_pods' 'test_many_pods')
	fi
	status=0
	for n in $(seq 0 $((${#TestsTypes[@]}-1))); do
		./testrunner.sh run -P "$InstancesFile" "${TestsTypes[$n]}" "${TestsToRun[$n]}" -r opencontrailnightly/contrail-test-test:latest
		status=$((status + $?))
	done
	return $status
}

finalize_test_run()
{
	cd ~/contrail-test-runs || return 0
	find . -maxdepth 1 -type d -group docker -exec sudo chown -R "$ThisUser:$ThisUser" "{}" \;
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
status=$?
finalize_test_run
exit "$status"
