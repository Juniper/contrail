#!/bin/bash
set -x

export PATH=/root/go/bin/:$PATH

docker ps | grep "config_api"
read

kubectl get pods --all-namespaces
read

tee namespace-blue.json << "EOF"
{
  "kind": "Namespace",
  "apiVersion": "v1",
  "metadata": {
    "name": "atom-blue",
    "labels": {
      "name": "atom-blue"
    },
    "annotations": {
      "opencontrail.org/network": "{'project': 'k8s-atom-blue', 'domain': 'default-domain', 'name': 'vn_blue'}"
    }
  }
}
EOF
read

kubectl create -f namespace-blue.json
read

tee vn_blue.yml << "EOF"
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
    fq_name: ["default-domain", "k8s-atom-blue", "vn_blue"]
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

contrailcli -c config.yml sync vn_blue.yml
read


read -p 'Run pod in second terminal using command: "kubectl run -i --tty busybox-one --image=busybox --namespace atom-blue -- sh"'

kubectl run -i --tty busybox-two --image=busybox --namespace atom-blue -- sh
