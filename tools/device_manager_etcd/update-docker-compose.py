#!/usr/bin/env python
import yaml

docker_compose_path = "/etc/contrail/config/docker-compose.yaml"

with open(docker_compose_path) as f:
    docker_compose = yaml.load(f)

schema = docker_compose["services"]["devicemgr"]
schema["image"] = "danielfurmancl/contrail-controller-config-devicemgr:1e7204fa1"

environment = schema.setdefault("environment", [])
for entry in [
        "NOTIFICATION_DRIVER=etcd",
        "DB_DRIVER=etcd",
        "ETCD_USE_SSL=true",
        "ETCD_SSL_KEYFILE=/etc/kubernetes/pki/etcd/peer.key",
        "ETCD_SSL_CERTFILE=/etc/kubernetes/pki/etcd/peer.crt",
        "ETCD_SSL_CA_CERT=/etc/kubernetes/pki/etcd/ca.crt",
]:
    if entry not in environment:
        environment.append(entry)

etcd_pki_mount = "/etc/kubernetes/pki/etcd:/etc/kubernetes/pki/etcd:ro"
volumes = schema.setdefault("volumes", [])
if etcd_pki_mount not in volumes:
    volumes.append(etcd_pki_mount)

with open(docker_compose_path, "w") as f:
    yaml.dump(docker_compose, f)
