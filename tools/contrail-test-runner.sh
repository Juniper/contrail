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
	set +x
	echo -e "Usage: $(basename "$0") [-h] [-D] [-t <test_spec> ...]\n"
	echo 'This script is intended to be used on a node with contrail is installed '
	echo 'with contrail-ansible-deployer and later tools/deploy-for_k8s.sh is used'
	echo 'This script will run sanity tests defined in https://github.com/Juniper/contrail-test'
	echo
	echo -e "\t-D => Don't check contrail-go dockers before tests"
	echo -e "\t-t => Test cases/tags to be run by contrail-test, specified as contrail-test flags (eg. \`-t -c test_many_pods -T k8s_sanity\`"
	echo -e "\t\ttest_spec is (-c|-T|-f name)"
}

declare -a TestsTypes
declare -a TestsToRun
TestCount=0
GatherTests()
{
	shiftem=1
	TestCount=0
	shift
	while :; do
		case "$1" in
			-c | -T | -f) ttype="$1"; name="$2"; shift 2; shiftem=$((shiftem+2));;
			*) break;;
		esac
		[ -z "$name" ] && { Usage; exit 1; }
		TestsTypes[$TestCount]="$ttype"
		TestsToRun[$TestCount]="$name"
		TestCount=$((TestCount+1))
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
	type -t yq > /dev/null || { sudo yum install -y jq; sudo pip install yq; type -t yq > /dev/null; } || { echo "Command 'yq' not available - aborting"; exit 2; }
	yq -y 'with_entries(select(.key == "provider_config")) | with_entries(.value[] +={ "ssh_user": "root", "ssh_pwd": "contrail123"})' "$InstancesFile" > /tmp/instances-provider.yaml
	yq -y 'with_entries(select(.key != "provider_config"))' "$InstancesFile" > /tmp/instances-rest.yaml
	cp -f /tmp/instances-provider.yaml "$InstancesFile"
	cat /tmp/instances-rest.yaml >> "$InstancesFile"
	modify_sshd=0
	sudo grep '^PasswordAuthentication yes$' /etc/ssh/sshd_config || modify_sshd=1
	sudo grep '^PermitRootLogin yes$' /etc/ssh/sshd_config || modify_sshd=1
	if [ $modify_sshd -eq 1 ]; then
		#shellcheck disable=SC2024
		sudo grep -vE 'PasswordAuthentication|PermitRootLogin' /etc/ssh/sshd_config > /tmp/sshd_config
		cat >> /tmp/sshd_config <<EOF
PasswordAuthentication yes
PermitRootLogin yes
EOF
		sudo cp -f /tmp/sshd_config /etc/ssh/sshd_config
		sudo service sshd restart
	fi
	echo 'contrail123' | sudo passwd --stdin root
}

prepare_test_env()
{
	cd
	type -t wget > /dev/null || { sudo yum install -y wget; type -t wget > /dev/null; } || { echo "Command 'wget' not available - aborting"; exit 2; }
	[ ! -e ./testrunner.sh ] && { wget https://github.com/Juniper/contrail-test/raw/master/testrunner.sh || echo "Could not download testrunner script from https://github.com/Juniper/contrail-test/raw/master/testrunner.sh"; }
	[ ! -e ./testrunner.sh ] && { echo "Missing 'testrunner' script - aborting!"; exit 2; }
	sed -i '/^\s\+tput/ d' ./testrunner.sh
	chmod +x ./testrunner.sh
	docker pull opencontrailnightly/contrail-test-test:latest || { echo 'Fail to pull docker image'; exit 2; }
	ensure_root_access
}

TestStatusMsg=''
run_contrail_test()
{
	echo 'Running testrunner...'
	echo
	if [ 0 -eq ${#TestsTypes[@]} ]; then
		TestsTypes=('-T')
		TestsToRun=('ci_contrail_go_k8s_sanity')
	fi
	status=0
	runs=0
	for n in $(seq 0 $((${#TestsTypes[@]}-1))); do
		./testrunner.sh run -P "$InstancesFile" "${TestsTypes[$n]}" "${TestsToRun[$n]}" -r opencontrailnightly/contrail-test-test:latest
		status=$((status + $?))
		runs=$((runs+1))
	done
	TestStatusMsg="Testing summary: pass $((runs-status)) out of $runs"
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
echo "$TestStatusMsg"
exit "$status"
