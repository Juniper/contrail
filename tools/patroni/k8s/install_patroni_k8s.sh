#!/bin/bash

[[ -z $(docker images -q patroni_k8s) ]] || { echo "Patroni image for k8s already exists. Skipping building docker image." ; exit 0; }

tmpdir=$(mktemp -d -t patroni_k8s-XXXXXX) || { echo "Failed to create temporary directory" ; exit 1; }
echo "Downloading files"
(cd $tmpdir && curl -LO "https://raw.githubusercontent.com/zalando/patroni/master/kubernetes/Dockerfile" --connect-timeout 60) || { echo "Failed to download Dockerfile" ; exit 1; }
(cd $tmpdir && curl -LO "https://raw.githubusercontent.com/zalando/patroni/master/kubernetes/entrypoint.sh" --connect-timeout 60)|| { echo "Failed to download entrypoint.sh" ; exit 1; }
echo "Building image"
(cd $tmpdir && docker build -t patroni_k8s .) || { echo "Failed to build docker image" ; exit 1; }
rm -rf $tmpdir || { echo "Failed to remove temporary directory" ; exit 1; }

