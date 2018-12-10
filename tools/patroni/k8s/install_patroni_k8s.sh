#!/bin/bash

docker build -t patroni_k8s . || { echo "Failed to build docker image" ; exit 1; }

