#!/bin/bash
set +x

PATH="$PATH:/usr/go/bin"
PATH=$(go env GOPATH)/bin/:$PATH

tee namespace1.json << "EOF"
{
  "kind": "Namespace",
  "apiVersion": "v1",
  "metadata": {
    "name": "namespace1",
    "labels": {
      "name": "namespace1"
    },
    "annotations": {
      "opencontrail.org/network": "{'project': 'k8s-default', 'domain': 'default-domain', 'name': 'vn1_lrp'}"
    }
  }
}
EOF
read

sudo kubectl create -f namespace1.json
read

tee namespace2.json << "EOF"
{
  "kind": "Namespace",
  "apiVersion": "v1",
  "metadata": {
    "name": "namespace2",
    "labels": {
      "name": "namespace2"
    },
    "annotations": {
      "opencontrail.org/network": "{'project': 'k8s-default', 'domain': 'default-domain', 'name': 'vn2_lrp'}"
    }
  }
}
EOF
read

sudo kubectl create -f namespace2.json
read

tee ipam1.yml << "EOF"
resources:
- kind: network_ipam
  data:
    parent_type: project
    fq_name:
    - default-domain
    - k8s-default
    - net_ipam_blue_lrp
    ipam_subnet_method: flat-subnet
    ipam_subnets:
      subnets:
      - addr_from_start: true
        alloc_unit: 0
        allocation_pools: []
        default_gateway: 13.32.0.1
        dhcp_option_list:
        dns_nameservers: []
        dns_server_address: 13.32.0.2
        enable_dhcp: true
        subnet:
          ip_prefix: 13.32.0.0
          ip_prefix_len: 12
EOF
read

contrailcli -c config.yml sync ipam1.yml
read

tee ipam2.yml << "EOF"
resources:
- kind: network_ipam
  data:
    parent_type: project
    fq_name:
    - default-domain
    - k8s-default
    - net_ipam_red_lrp
    ipam_subnet_method: flat-subnet
    ipam_subnets:
      subnets:
      - addr_from_start: true
        alloc_unit: 0
        allocation_pools: []
        default_gateway: 14.32.0.1
        dhcp_option_list:
        dns_nameservers: []
        dns_server_address: 14.32.0.2
        enable_dhcp: true
        subnet:
          ip_prefix: 14.32.0.0
          ip_prefix_len: 12
EOF
read

contrailcli -c config.yml sync ipam2.yml
read

tee vn1.yml << "EOF"
resources:
- kind: virtual_network
  data:
    address_allocation_mode: flat-subnet-only
    parent_type: project
    fabric_snat: false
    virtual_network_properties:
      forwarding_mode: l3
      allow_transit: false
      rpf: enable
      mirror_destination: false
    flood_unknown_unicast: false
    layer2_control_word: false
    network_ipam_refs:
    - to:
      - default-domain
      - k8s-default
      - net_ipam_blue_lrp
      attr:
        ipam_subnets: []
    fq_name:
    - default-domain
    - k8s-default
    - vn1_lrp
EOF
read

contrailcli -c config.yml sync vn1.yml
read

tee vn2.yml << "EOF"
resources:
- kind: virtual_network
  data:
    address_allocation_mode: flat-subnet-only
    parent_type: project
    fabric_snat: false
    virtual_network_properties:
      forwarding_mode: l3
      allow_transit: false
      rpf: enable
      mirror_destination: false
    flood_unknown_unicast: false
    layer2_control_word: false
    network_ipam_refs:
    - to:
      - default-domain
      - k8s-default
      - net_ipam_red_lrp
      attr:
        ipam_subnets: []
    fq_name:
    - default-domain
    - k8s-default
    - vn2_lrp
EOF
read

contrailcli -c config.yml sync vn2.yml
read

tee vmi1.yml << "EOF"
resources:
- kind: virtual_machine_interface
  data:
    parent_type: project
    fq_name:
    - default-domain
    - k8s-default
    - vmi_blue_lrp
    virtual_network_refs:
    - to:
      - default-domain
      - k8s-default
      - vn1_lrp
    virtual_machine_interface_device_owner: network:router_interface
    display_name: vmi_blue
EOF
read

contrailcli -c config.yml sync vmi1.yml
read

tee vmi2.yml << "EOF"
resources:
- kind: virtual_machine_interface
  data:
    parent_type: project
    fq_name:
    - default-domain
    - k8s-default
    - vmi_red_lrp
    virtual_network_refs:
    - to:
      - default-domain
      - k8s-default
      - vn2_lrp
    virtual_machine_interface_device_owner: network:router_interface
    display_name: vmi_red
EOF
read

contrailcli -c config.yml sync vmi2.yml
read

tee iip1.yml << "EOF"
resources:
- kind: instance_ip
  data:
    fq_name:
    - instance_ip_blue_lrp
    display_name: instance_ip_blue
    instance_ip_address: ''
    virtual_machine_interface_refs:
    - to:
      - default-domain
      - k8s-default
      - vmi_blue_lrp
    virtual_network_refs:
    - to:
      - default-domain
      - k8s-default
      - vn1_lrp
EOF
read

contrailcli -c config.yml sync iip1.yml
read

tee iip2.yml << "EOF"
resources:
- kind: instance_ip
  data:
    fq_name:
    - instance_ip_red_lrp
    display_name: instance_ip_red
    instance_ip_address: ''
    virtual_machine_interface_refs:
    - to:
      - default-domain
      - k8s-default
      - vmi_red_lrp
    virtual_network_refs:
    - to:
      - default-domain
      - k8s-default
      - vn2_lrp
EOF
read

contrailcli -c config.yml sync iip2.yml
read

tee lr.yml << "EOF"
resources:
- kind: logical_router
  data:
    name: logical_router_1
    fq_name:
    - default-domain
    - k8s-default
    - logical_router_1_lrp
    parent_type: project
    virtual_machine_interface_refs:
    - to:
      - default-domain
      - k8s-default
      - vmi_blue_lrp
    - to:
      - default-domain
      - k8s-default
      - vmi_red_lrp
    virtual_network_refs: []
    id_perms:
      enable: true
    configured_route_target_list:
      route_target: []
EOF
read

contrailcli -c config.yml sync lr.yml
read

read -p 'Run pod in second terminal using command: "sudo kubectl run -i --tty busybox-one --image=busybox --namespace namespace2 -- sh"'

sudo kubectl run -i --tty busybox-two --image=busybox --namespace namespace1 -- sh