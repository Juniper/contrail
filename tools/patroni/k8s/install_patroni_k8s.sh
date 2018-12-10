#!/bin/bash

[[ -z $(docker images -q patroni_k8s) ]] || { echo "Patroni image for k8s already exists. Skipping building docker image." ; exit 0; }

docker build -t patroni_k8s . || { echo "Failed to build docker image" ; exit 1; }

