# Deployment for kubernetes (k8s)

## Prerequisites
* Clean VM with Centos7
* Installed k8s + contrail with contrail-ansible-deployer
  - Sample instances.yaml file presented below. Tokens @localip and @routerip must be changed before running contrail-ansible-deployer, passwords should be adjusted.

```
provider_config:
  bms:
    ssh_user: centos
    ssh_private_key: /home/centos/id_rsa
    domainsuffix: local
instances:
  bms1:
    provider: bms
    ip: @localip
    roles:
      config_database:
      config:
      control:
      webui:
      vrouter:
      k8s_master:
      kubemanager:
      k8s_node:

global_configuration:
  CONTAINER_REGISTRY: opencontrailnightly
contrail_configuration:
  CONTRAIL_VERSION: latest
  CLOUD_ORCHESTRATOR: kubernetes
  RABBITMQ_NODE_PORT: 5673
  VROUTER_GATEWAY: @routerip
  PHYSICAL_INTERFACE: eth0
  AUTH_MODE: keystone
  KEYSTONE_AUTH_ADMIN_PASSWORD: contrail123
  KEYSTONE_AUTH_HOST: @localip
  KEYSTONE_AUTH_URL_VERSION: "/v3"
  JVM_EXTRA_OPTS: "-Xms1g -Xmx2g"
  CONFIG_NODEMGR__DEFAULTS__minimum_diskGB: 2
  DATABASE_NODEMGR__DEFAULTS__minimum_diskGB: 2
kolla_config:
  kolla_globals:
    network_interface: "eth0"
    kolla_external_vip_interface: "eth0"
    enable_haproxy: "no"
    openstack_release: "ocata"
  kolla_passwords:
    keystone_admin_password: contrail123
```

## Deployment
* On remote (target) machine clone https://github.com/Juniper/contrail into $HOME/go/src/github.com/Juniper
``` shell
mkdir -p $HOME/go/src/github.com/Juniper
cd !$
git clone https://github.com/Juniper/contrail
```
* Run a deployment script
``` shell
cd contrail
./tools/deploy-for_k8s.sh
```
