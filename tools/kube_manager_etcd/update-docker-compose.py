#!/usr/bin/env python
import yaml

docker_compose_path = "/etc/contrail/kubemanager/docker-compose.yaml"

with open(docker_compose_path) as f:
    docker_compose = yaml.load(f)

docker_compose["services"]["kubemanager"]["image"] = "danielfurmancl/contrail-kubernetes-kube-manager:61dc46749"

environment = docker_compose["services"]["kubemanager"].setdefault("environment", [])
notification_driver = "NOTIFICATION_DRIVER=etcd"
if notification_driver not in environment:
    environment.append(notification_driver)

db_driver = "DB_DRIVER=etcd"
if db_driver not in environment:
    environment.append(db_driver)

etcd_pki_mount = "/etc/kubernetes/pki/etcd:/etc/kubernetes/pki/etcd:ro"
volumes = docker_compose["services"]["kubemanager"].setdefault("volumes", [])
if etcd_pki_mount not in volumes:
    volumes.append(etcd_pki_mount)

with open(docker_compose_path, "w") as f:
    yaml.dump(docker_compose, f)
