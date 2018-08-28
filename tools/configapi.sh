#!/bin/bash

RunDockers="keystone cassandra zookeeper rabbitmq config_api config_schema"

Usage()
{
	echo "Usage: $(basename "$0") [-k] [-n <mode>] [dockers]"
	echo "-k => Don't remove dockers before running new ones"
	echo "-n => Use specified 'NetworkMode' for docker"
	echo "Default dockers: $RunDockers"
	echo "Available dockers: $(grep -E '^run_docker_.*\(\)$' "$0" | sed 's/^run_docker_//; s/()$//;' | tr '\n' ' ')"
}

remove_dockers()
{
	docker rm -f \
		some-cassandra \
		some-zookeeper \
		some-rabbit \
		some-keystone \
		config-api \
		schema-transformer
		some-redis \
		config-ui
}

run_docker_keystone()
{
	docker run --name some-keystone \
		--net "$Network" \
		-v "$PWD/keystone/apache2:/etc/apache2/sites-available/" \
		-v "$PWD/keystone/etc:/etc/keystone" \
		-v "$PWD/keystone/scripts:/tmp" \
		-e OS_USERNAME=admin \
		-e OS_PASSWORD=contrail123 \
		-e OS_PROJECT_NAME=admin \
		-e OS_PROJECT_DOMAIN_ID=default \
		-e OS_USER_DOMAIN_ID=default \
		-e OS_AUTH_URL=http://localhost:5000 \
		-e OS_IDENTITY_API_VERSION=3 \
		-p 5000:5000 \
		-d \
		openstackhelm/keystone:newton \
		bash /tmp/start.sh

	sleep 10

	docker exec some-keystone bash /tmp/init.sh
}

run_docker_cassandra()
{
	docker run --name some-cassandra \
		--net "$Network" \
		-p 9160:9160 \
		-p 9042:9042 \
		-e CASSANDRA_START_RPC=true \
		-e CASSANDRA_CLUSTER_NAME=ContrailConfigDB \
		-d cassandra:3.11.1
}

run_docker_redis()
{
	docker run --name some-redis \
		--net "$Network" \
		-d \
		redis:4.0.2
}

run_docker_zookeeper()
{
	docker run --net "$Network" --name some-zookeeper -d zookeeper:latest
}

run_docker_rabbitmq()
{
	docker run  --net "$Network" --name some-rabbit -p 5672:5672 -d rabbitmq:3.6.10
}

run_docker_config_api()
{
	local link_params="--link some-cassandra:cassandra --link some-zookeeper:zookeeper --link some-rabbit --link some-keystone"
	[ $NoLink -eq 1 ] && link_params=''
	docker run $link_params \
		--net "$Network" \
		--name config-api \
		-p 8082:8082 \
		-d \
		-e CONFIG_API_PORT=8082 \
		-e CONFIG_API_INTROSPECT_PORT=8084 \
		-e LOG_LEVEL=SYS_NOTICE \
		-e log_local=true \
		-e AUTH_MODE=keystone \
		-e AAA_MODE=cloud-admin \
		-e CONFIGDB_SERVERS=localhost:9160 \
		-e ZOOKEEPER_SERVERS=localhost \
		-e RABBITMQ_SERVERS=localhost \
		-e KEYSTONE_AUTH_ADMIN_USER=admin \
		-e KEYSTONE_AUTH_ADMIN_TENANT=admin \
		-e KEYSTONE_AUTH_ADMIN_PASSWORD=contrail123 \
		-e KEYSTONE_AUTH_USER_DOMAIN_NAME=Default \
		-e KEYSTONE_AUTH_PROJECT_DOMAIN_NAME=Default \
		-e KEYSTONE_AUTH_URL_VERSION=/v3 \
		-e KEYSTONE_AUTH_HOST=localhost \
		-e KEYSTONE_AUTH_PROTO=http \
		-e KEYSTONE_AUTH_ADMIN_PORT=35357 \
		-e KEYSTONE_AUTH_PUBLIC_PORT=5000 \
		opencontrailnightly/contrail-controller-config-api
}

run_docker_config_schema()
{
	link_params="--link some-cassandra:cassandra --link some-zookeeper:zookeeper --link some-rabbit --link some-keystone --link config-api"
	[ $NoLink -eq 1 ] && link_params=''
	docker run $link_params \
		--net "$Network" \
	    --name schema-transformer \
	    -d \
	    -e CONFIG_API_PORT=8082 \
	    -e CONFIG_API_INTROSPECT_PORT=8084 \
	    -e LOG_LEVEL=SYS_NOTICE \
	    -e log_local=true \
	    -e AUTH_MODE=keystone \
	    -e AAA_MODE=cloud-admin \
	    -e CONFIGDB_SERVERS=localhost:9160 \
	    -e ZOOKEEPER_SERVERS=localhost \
	    -e RABBITMQ_SERVERS=localhost \
	    -e CONFIG_NODES=config-api \
	    -e KEYSTONE_AUTH_ADMIN_USER=admin \
	    -e KEYSTONE_AUTH_ADMIN_TENANT=admin \
	    -e KEYSTONE_AUTH_ADMIN_PASSWORD=contrail123 \
	    -e KEYSTONE_AUTH_USER_DOMAIN_NAME=Default \
	    -e KEYSTONE_AUTH_PROJECT_DOMAIN_NAME=Default \
	    -e KEYSTONE_AUTH_URL_VERSION=/v3 \
	    -e KEYSTONE_AUTH_HOST=localhost \
	    -e KEYSTONE_AUTH_PROTO=http \
	    -e KEYSTONE_AUTH_ADMIN_PORT=35357 \
	    -e KEYSTONE_AUTH_PUBLIC_PORT=5000 \
	    opencontrailnightly/contrail-controller-config-schema
}

run_docker_webui()
{
	link_params="--link some-cassandra:cassandra --link some-keystone --link some-redis --link config-api"
	[ $NoLink -eq 1 ] && link_params=''
	docker run $link_params \
		--net "$Network" \
		--name config-ui \
		-v "$PWD/webui:/etc/contrail" \
		-p 8143:8143 \
		-p 8080:8080 \
		-it \
		-e LOG_LEVEL=SYS_NOTICE \
		-e log_local=true \
		-e AUTH_MODE=keystone \
		-e CLOUD_ORCHESTRATOR=openstack \
		-e CONFIGDB_SERVERS=localhost:9160 \
		-e ZOOKEEPER_SERVERS=localhost \
		-e RABBITMQ_SERVERS=localhost \
		-e AAA_MODE=cloud-admin \
		-e CONFIG_NODES=config-api \
		-e KEYSTONE_AUTH_ADMIN_USER=admin \
		-e KEYSTONE_AUTH_ADMIN_TENANT=admin \
		-e KEYSTONE_AUTH_ADMIN_PASSWORD=contrail123 \
		-e KEYSTONE_AUTH_USER_DOMAIN_NAME=Default \
		-e KEYSTONE_AUTH_PROJECT_DOMAIN_NAME=Default \
		-e KEYSTONE_AUTH_URL_VERSION=/v3 \
		-e KEYSTONE_AUTH_HOST=localhost \
		-e KEYSTONE_AUTH_PROTO=http \
		-e KEYSTONE_AUTH_ADMIN_PORT=35357 \
		-e KEYSTONE_AUTH_PUBLIC_PORT=5000 \
		--entrypoint "bash" \
		opencontrailnightly/contrail-controller-webui-web
}

RemoveDockers=1
Network='bridge' # This is default for `docker run` if no `--net` param is specified
while :; do
	case "$1" in
		'-n') Network="$2"; shift 2;;
		'-k') RemoveDockers=0; shift;;
		'-h') Usage; exit 0;;
		*) break;;
	esac
done
[ $RemoveDockers -eq 1 ] && remove_dockers

[ ! -z "$1" ] && RunDockers="$*"
NoLink=0
[ "$Network" != 'bridge' ] && NoLink=1

for docker in $RunDockers; do 
	eval "run_docker_$docker"
done

docker exec some-keystone openstack token issue

TOKEN=$(docker exec some-keystone openstack token issue | awk '/ id /{print $4}')
export TOKEN

echo 'Contrail VNC API starts running.. try REST API endpoint'
echo "export TOKEN=$TOKEN"
#shellcheck disable=2016
echo 'curl -H "X-Auth-Token: $TOKEN" http://localhost:8082/virtual-networks'
