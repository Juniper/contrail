#!/bin/bash

set -x

PATH="$PATH:/usr/go/bin"
PATH="$(go env GOPATH)/bin/:$PATH"

tee namespace-blue.json << "EOF"
{
  "kind": "Namespace",
  "apiVersion": "v1",
  "metadata": {
    "name": "blue",
    "labels": {
      "name": "blue"
    },
    "annotations": {
      "opencontrail.org/network": "{'project': 'k8s-default', 'domain': 'default-domain', 'name': 'vn-blue'}"
    }
  }
}
EOF
read

sudo kubectl create -f namespace-blue.json
read

tee k8s-pod-ipam.yml << "EOF"
resources:
- kind: network_ipam
  data:
    parent_type: project
    fq_name:
    - default-domain
    - k8s-default
    - k8s-pod-ipam
    ipam_subnet_method: flat-subnet
    ipam_subnets:
      subnets:
      - addr_from_start: true
        alloc_unit: 0
        default_gateway: 13.32.0.1
        dns_server_address: 13.32.0.2
        enable_dhcp: true
        subnet:
          ip_prefix: 13.32.0.0
          ip_prefix_len: 12
EOF
read

contrailcli -c config.yml sync k8s-pod-ipam.yml
read

tee vn-blue.yml << "EOF"
resources:
- kind: virtual_network
  data:
    virtual_network_properties:
      forwarding_mode: l3
    fq_name:
    - default-domain
    - k8s-default
    - vn-blue
    address_allocation_mode: flat-subnet-only
    parent_type: project
    network_ipam_refs:
    - to: ["default-domain", "k8s-default", "k8s-pod-ipam"]
      attr:
        ipam_subnets: []
EOF
read

contrailcli -c config.yml sync vn-blue.yml
read

read -p 'Run pod in second terminal using command: "sudo kubectl run -i --tty busybox-one --image=busybox --namespace blue -- sh"'

sudo kubectl run -i --tty busybox-two --image=busybox --namespace blue -- sh
