#!/bin/bash

[[ "$(docker images -q patroni)" == "" ]] || { echo "Patroni image already exists. Skipping pulling docker image." ; exit 0; }

docker pull kaweue/patroni:latest || { echo "Error while pulling repository" ; exit 1; }
docker tag katrybacka/patroni:latest patroni:latest || { echo "Failed to tag repository" ; exit 1; }

