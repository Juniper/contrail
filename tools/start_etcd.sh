#!/usr/bin/env bash
# Start etcd service.

set -o errexit
set -o nounset
set -o pipefail

function main {

}

main



#DB_USER=${DB_USER:-root}
#DB_PASSWORD=${DB_PASSWORD:-contrail123}
#DB_HOSTNAME=${DB_HOSTNAME:-localhost}
#DB_PORT=${DB_PORT:-3307}

#function main {
#	echo "Running containers"
#	docker-compose up -d
#
#	await_mysql
#	./integration/init_db.sh
#
#	echo "Running integration tests"
#	if ! go test -tags=integration ./integration/...
#	then
#		teardown 1
#	fi
#
#	teardown 0
#}
#
#function await_mysql {
#	local readonly MYSQLADMIN="mysqladmin --user ${DB_USER} --password=${DB_PASSWORD} --protocol tcp \
#		--host ${DB_HOSTNAME} --port ${DB_PORT}"
#
#	echo "Awaiting MySQL"
#	until ${MYSQLADMIN} ping &> /dev/null; do
#	  echo "MySQL is unavailable - waiting"
#	  sleep 1
#	done
#}
#
#function teardown {
#	local readonly EXIT_CODE=$1
#
#	# Docker containers on CircleCI cannot be removed
#	# More information: https://circleci.com/docs/1.0/docker-btrfs-error/
#	if [ -v "CI" ]; then
#		echo "Stopping containers"
#		docker-compose stop
#	else
#		echo "Stopping and removing containers"
#		docker-compose down --volumes
#	fi
#
#	exit ${EXIT_CODE}
#}
#
#main
