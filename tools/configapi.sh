#!/bin/bash

docker rm -f some-cassandra some-zookeeper some-rabbit some-keystone config-api

docker run --name some-keystone \
    -v `pwd`/keystone/apache2:/etc/apache2/sites-available/ \
    -v `pwd`/keystone/etc:/etc/keystone \
    -v `pwd`/keystone/scripts:/tmp \
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

docker run --name some-cassandra \
    -e CASSANDRA_START_RPC=true \
    -e CASSANDRA_CLUSTER_NAME=ContrailConfigDB \
    -d cassandra:3.11.1

docker run --name some-zookeeper -d zookeeper:latest
docker run --name some-rabbit -d rabbitmq:3.6.10

docker run \
    --name config-api \
    -p 8082:8082 \
    --link some-cassandra:cassandra \
    --link some-zookeeper:zookeeper \
    --link some-rabbit \
    --link some-keystone \
    -d \
    -e CONFIG_API_PORT=8082 \
    -e CONFIG_API_INTROSPECT_PORT=8084 \
    -e LOG_LEVEL=SYS_NOTICE \
    -e log_local=true \
    -e AUTH_MODE=keystone \
    -e AAA_MODE=cloud-admin \
    -e CONFIGDB_SERVERS=some-cassandra:9160 \
    -e ZOOKEEPER_SERVERS=some-zookeeper \
    -e RABBITMQ_SERVERS=some-rabbit \
    -e KEYSTONE_AUTH_ADMIN_USER=admin \
    -e KEYSTONE_AUTH_ADMIN_TENANT=admin \
    -e KEYSTONE_AUTH_ADMIN_PASSWORD=contrail123 \
    -e KEYSTONE_AUTH_USER_DOMAIN_NAME=Default \
    -e KEYSTONE_AUTH_PROJECT_DOMAIN_NAME=Default \
    -e KEYSTONE_AUTH_URL_VERSION=/v3 \
    -e KEYSTONE_AUTH_HOST=some-keystone \
    -e KEYSTONE_AUTH_PROTO=http \
    -e KEYSTONE_AUTH_ADMIN_PORT=35357 \
    -e KEYSTONE_AUTH_PUBLIC_PORT=5000 \
    opencontrailnightly/contrail-controller-config-api

export TOKEN=`docker exec some-keystone openstack token issue | awk '/ id /{print $4}'`

echo 'Contrail VNC API starts running.. try REST API endpoint'
echo 'export TOKEN='$TOKEN
echo 'curl -H "X-Auth-Token: $TOKEN" http://localhost:8082/virtual-networks'