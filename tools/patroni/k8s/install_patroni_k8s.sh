#!/bin/bash

docker build -t patroni_k8s $(dirname $0) || { echo "Failed to build docker image" ; exit 1; }
