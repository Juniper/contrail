#!/usr/bin/env python
import yaml
import os

docker_compose_path = "/etc/contrail/control/docker-compose.yaml"

with open(docker_compose_path) as f:
    docker_compose = yaml.load(f)

docker_compose["services"]["control"]["image"] = "danielfurmancl/contrail-controller-control-control:etcd_sync_2"

if os.environ.get('ORCHESTRATOR') == "k8s":
    etcd_pki_mount = "/etc/kubernetes/pki/etcd:/etc/kubernetes/pki/etcd:ro"
    volumes = docker_compose["services"]["control"].setdefault("volumes", [])
    if etcd_pki_mount not in volumes:
        volumes.append(etcd_pki_mount)

config = ["CONFIGDB_USE_ETCD=true"]
if os.environ.get('ORCHESTRATOR') == "k8s":
    config += [
        "CONFIGDB_ETCD_USE_SSL=true",
        "CONFIGDB_ETCD_KEY_FILE=/etc/kubernetes/pki/etcd/peer.key",
        "CONFIGDB_ETCD_CERT_FILE=/etc/kubernetes/pki/etcd/peer.crt",
        "CONFIGDB_ETCD_CA_CERT_FILE=/etc/kubernetes/pki/etcd/ca.crt",
    ]
else:
    config += ["CONFIGDB_ETCD_USE_SSL=false"]

environment = docker_compose["services"]["control"].setdefault("environment", [])
for entry in config:
    if entry not in environment:
        environment.append(entry)

with open(docker_compose_path, "w") as f:
    yaml.dump(docker_compose, f)
