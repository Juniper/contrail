#!/bin/bash

build_and_run_contrail-go_docker()
{
    local ContrailGoDocker='contrail-go-config-node'
    [ "$(docker ps -a -f "name=$ContrailGoDocker" --format '{{.ID}}' | wc -l)" -ne 0 ] && docker rm -f "$ContrailGoDocker"
    docker run -d --name "$ContrailGoDocker" --net host --volume /etc/kubernetes/pki/etcd:/etc/kubernetes/pki/etcd:ro \
        contrail-go-config
}

# Modify k8s config (subst contrail-go-config IP address as config-node) and restart if needed
ensure_kubemanager_config_nodes()
{
    local GoConfigIP="$1"
    local ModifyKubeConfig=1
    grep -qE "^CONFIG_NODES\\W*=\\W*$GoConfigIP" /etc/contrail/common_kubemanager.env && ModifyKubeConfig=0
    if [ $ModifyKubeConfig -eq 1 ]; then
        sudo sed "-ibak$(date +%s)" "s/^CONFIG_NODES=.*/CONFIG_NODES=$GoConfigIP/" /etc/contrail/common_kubemanager.env
    fi
}

ensure_keystone_on_localhost()
{
    grep -q -x -F "Listen 127.0.0.1:5000" /etc/kolla/keystone/wsgi-keystone.conf || \
        sudo sh -c 'echo "Listen 127.0.0.1:5000" >> /etc/kolla/keystone/wsgi-keystone.conf'
    docker restart keystone
}

schema_transformer_up()
{
    docker-compose -f "/etc/contrail/config/docker-compose.yaml" up -d schema
}

device_manager_up()
{
    docker-compose -f "/etc/contrail/config/docker-compose.yaml" up -d devicemgr
}

compose_up()
{
    for docker_dir in "$@"
    do
        docker-compose -f "/etc/contrail/${docker_dir}/docker-compose.yaml" up -d
    done
}

compose_down()
{
    for docker_dir in "$@"
    do
        docker-compose -f "/etc/contrail/${docker_dir}/docker-compose.yaml" down
    done
}

clear_config_database()
{
    docker-compose -f /etc/contrail/config_database/docker-compose.yaml down -v
    docker-compose -f /etc/contrail/config_database/docker-compose.yaml up -d zookeeper
}

