#!/bin/bash
set -x

PATH="$PATH:/usr/go/bin"
PATH="$(go env GOPATH)/bin/:$PATH"

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

sudo kubectl create -f namespace.json
read

tee vn_pink.yml << "EOF"
resources:
- kind: virtual_network
  data:
    virtual_network_properties:
      forwarding_mode: l3
      allow_transit:
      network_id:
      max_flow_rate:
      mirror_destination: false
      vxlan_network_identifier:
      max_flows:
      rpf:
    fq_name: ["default-domain", "k8s-atom-pink", "vn_pink"]
    address_allocation_mode: flat-subnet-only
    parent_type: project
    network_ipam_refs:
    - to: ["default-domain", "k8s-default", "k8s-pod-ipam"]
      attr:
        ipam_subnets: []
        host_routes:
    fabric_snat: false
EOF
read

contrailcli -c config.yml sync vn_pink.yml
read
