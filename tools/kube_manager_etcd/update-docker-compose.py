#!/usr/bin/env python
import yaml

docker_compose_path = "/etc/contrail/kubemanager/docker-compose.yaml"

with open(docker_compose_path) as f:
    docker_compose = yaml.load(f)

docker_compose["services"]["kubemanager"]["image"] = "danielfurmancl/contrail-kubernetes-kube-manager:etcd-client-certs"

etcd_pki_mount = "/etc/kubernetes/pki/etcd:/etc/kubernetes/pki/etcd:ro"
volumes = docker_compose["services"]["kubemanager"].setdefault("volumes", [])
if etcd_pki_mount not in volumes:
    volumes.append(etcd_pki_mount)

with open(docker_compose_path, "w") as f:
    yaml.dump(docker_compose, f)
