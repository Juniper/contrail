#!/usr/bin/env python
import yaml

docker_compose_path = "/etc/contrail/control/docker-compose.yaml"

with open(docker_compose_path) as f:
    docker_compose = yaml.load(f)

docker_compose["services"]["control"]["image"] = "katrybacka/contrail-controller-control-control:etcd_sync2"

etcd_pki_mount = "/etc/kubernetes/pki/etcd:/etc/kubernetes/pki/etcd:ro"
volumes = docker_compose["services"]["control"].setdefault("volumes", [])
if etcd_pki_mount not in volumes:
    volumes.append(etcd_pki_mount)

with open(docker_compose_path, "w") as f:
    yaml.dump(docker_compose, f)
