#!/bin/bash

kubectl create -f $(dirname $0)/patroni_k8s.yaml
