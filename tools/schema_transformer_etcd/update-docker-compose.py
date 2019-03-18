#!/usr/bin/env python
import yaml
import os

docker_compose_path = "/etc/contrail/config/docker-compose.yaml"

with open(docker_compose_path) as f:
    docker_compose = yaml.load(f)

schema = docker_compose["services"]["schema"]
schema["image"] = "kaweue/contrail-controller-config-schema:R6.0-9"
config = [
    "NOTIFICATION_DRIVER=etcd",
    "DB_DRIVER=etcd",
]

if os.environ.get('ORCHESTRATOR') == "k8s":
    config += [
        "ETCD_USE_SSL=true",
        "ETCD_SSL_KEYFILE=/etc/kubernetes/pki/etcd/peer.key",
        "ETCD_SSL_CERTFILE=/etc/kubernetes/pki/etcd/peer.crt",
        "ETCD_SSL_CA_CERT=/etc/kubernetes/pki/etcd/ca.crt",
    ]

environment = schema.setdefault("environment", [])
for entry in config:
    if entry not in environment:
        environment.append(entry)

if os.environ.get('ORCHESTRATOR') == "k8s":
    etcd_pki_mount = "/etc/kubernetes/pki/etcd:/etc/kubernetes/pki/etcd:ro"
    volumes = schema.setdefault("volumes", [])
    if etcd_pki_mount not in volumes:
        volumes.append(etcd_pki_mount)

with open(docker_compose_path, "w") as f:
    yaml.dump(docker_compose, f)
