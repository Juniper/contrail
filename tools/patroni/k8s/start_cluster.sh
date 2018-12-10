#!/bin/bash

NAMESPACE="pat-space"
CLUSTER_NAME="contrail"

[[ -z $(kubectl get ns | grep NAMESPACE) ]] || kubectl create ns $NAMESPACE || { echo "Failed to create namespace $NAMESPACE" ; exit 1; }

[[ -z "$1" ]] || CLUSTER_NAME="$1"

sed "s/CLUSTER_NAME/$CLUSTER_NAME/g" patroni_k8s.yaml.tmpl > patroni_k8s.yaml

kubectl create -f ./patroni_k8s.yaml --namespace $NAMESPACE

