#!/usr/bin/env python
import yaml

docker_compose_path = "/etc/contrail/kubemanager/docker-compose.yaml"

with open(docker_compose_path) as f:
    docker_compose = yaml.load(f)

docker_compose["services"]["kubemanager"]["image"] = "contrail-kubernetes-kube-manager:etcd"

with open(docker_compose_path, "w") as f:
    yaml.dump(docker_compose, f)
