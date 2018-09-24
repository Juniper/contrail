#!/bin/bash
set -x

CONTRAIL_ADDR=localhost:8082

tee namespace.json << "EOF"
{
  "kind": "Namespace",
  "apiVersion": "v1",
  "metadata": {
    "name": "atom-pink",
    "labels": {
      "name": "atom-pink"
    },
    "annotations": {
      "opencontrail.org/network": "{'project': 'k8s-atom-pink', 'domain': 'default-domain', 'name': 'vn_pink'}"
    }
  }
}
EOF
read

kubectl create -f namespace.json
read

tee vn_pink.json << "EOF"
{
	"virtual-network": {
		"virtual_network_properties": {
			"forwarding_mode": "l3"
		},
		"fq_name": [
			"default-domain",
			"k8s-atom-pink",
			"vn_pink"
		],
		"address_allocation_mode": "flat-subnet-only",
		"parent_type": "project",
		"network_ipam_refs": [
			{
				"to": [
				  "default-domain",
				  "k8s-default",
				  "k8s-pod-ipam"
				],
				"attr": { "ipam_subnets": [] }
			}
		],
		"fabric_snat": false
	}
}
EOF
read

curl -X POST -H "Content-Type: application/json; charset=UTF-8" -d @vn_pink.json $CONTRAIL_ADDR/virtual-networks
read
