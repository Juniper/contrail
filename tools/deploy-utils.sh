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

# compose is running docker-compose up or down on list of dirs
# or dir:service pairs. For each bare dir specified it will work on
# all services defined in relevant docker-compose.yaml file. 
# For each dir:service pair it will work on specified service only.
compose()
{
    case "$1" in
        --up ) action="up -d";;
        --down ) action="down";;
        * ) echo "Usage: compose (--up|--down) (<dir>|<dir>:<service>)..." && return 1;;
    esac
    shift

    for docker_dir_img in "$@"
    do
        case "$docker_dir_img" in
            *:* ) docker-compose -f "/etc/contrail/${docker_dir_img%:*}/docker-compose.yaml" $action ${docker_dir_img#*:};;
            * ) docker-compose -f "/etc/contrail/${docker_dir_img}/docker-compose.yaml" $action;;
        esac
    done

}

clear_config_database()
{
    docker-compose -f /etc/contrail/config_database/docker-compose.yaml down -v
    docker-compose -f /etc/contrail/config_database/docker-compose.yaml up -d zookeeper
}

