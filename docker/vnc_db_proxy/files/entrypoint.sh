#!/bin/bash

source /common.sh

pre_start_init

cassandra_server_list=$(echo $CONFIGDB_SERVERS | sed 's/,/ /g')

cat > /etc/contrail/contrail-api.conf << EOM
[DEFAULTS]
listen_ip_addr=${LISTEN_IP_ADDR}
listen_port=$CONFIG_API_PORT
http_server_port=${CONFIG_API_INTROSPECT_PORT}
log_file=$LOG_DIR/contrail-api.log
log_level=$LOG_LEVEL
log_local=$LOG_LOCAL
list_optimization_enabled=${CONFIG_API_LIST_OPTIMIZATION_ENABLED:-True}
auth=$AUTH_MODE
aaa_mode=$AAA_MODE
cloud_admin_role=$CLOUD_ADMIN_ROLE
global_read_only_role=$GLOBAL_READ_ONLY_ROLE
cassandra_server_list=$cassandra_server_list
zk_server_ip=$ZOOKEEPER_SERVERS

rabbit_server=$RABBITMQ_SERVERS
$rabbit_config
$kombu_ssl_config

collectors=$COLLECTOR_SERVERS

$sandesh_client_config
EOM

add_ini_params_from_env API /etc/contrail/contrail-api.conf

set_third_party_auth_config
set_vnc_api_lib_ini

exec "$@"