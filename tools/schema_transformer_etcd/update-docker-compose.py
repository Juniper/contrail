#!/usr/bin/env python
import yaml

docker_compose_path = "/etc/contrail/config/docker-compose.yaml"

with open(docker_compose_path) as f:
    docker_compose = yaml.load(f)

schema = docker_compose["services"]["schema"]
schema["image"] = "mateumann/contrail-controller-config-schema:queens-dev-with-R6.0-1"
environment = schema.setdefault("environment", [])

notification_driver = "NOTIFICATION_DRIVER=etcd"
db_driver = "DB_DRIVER=etcd"

if notification_driver not in environment:
    environment.append(notification_driver)

if db_driver not in environment:
    environment.append(db_driver)

with open(docker_compose_path, "w") as f:
    yaml.dump(docker_compose, f)
