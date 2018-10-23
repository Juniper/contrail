#!/usr/bin/env python
import yaml

with open("/etc/contrail/kubemanager/docker-compose.yaml") as f:
     docker_compose = yaml.load(f)

docker_compose["services"]["kubemanager"]["image"] = "contrail-kubernetes-kube-manager:etcd"

with open("/etc/contrail/kubemanager/docker-compose.yaml", "w") as f:
    yaml.dump(docker_compose, f)
