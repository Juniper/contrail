#!/usr/bin/env python
import yaml
import os

docker_compose_path = "/etc/contrail/config/docker-compose.yaml"

with open(docker_compose_path) as f:
    docker_compose = yaml.load(f)

svc_monitor = docker_compose["services"]["svcmonitor"]
svc_monitor["image"] = "kaweue/contrail-controller-config-svcmonitor:R6.0-7"
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

environment = svc_monitor.setdefault("environment", [])
for entry in config:
    if entry not in environment:
        environment.append(entry)

if os.environ.get('ORCHESTRATOR') == "k8s":
    etcd_pki_mount = "/etc/kubernetes/pki/etcd:/etc/kubernetes/pki/etcd:ro"
    volumes = svc_monitor.setdefault("volumes", [])
    if etcd_pki_mount not in volumes:
        volumes.append(etcd_pki_mount)

with open(docker_compose_path, "w") as f:
    yaml.dump(docker_compose, f)
