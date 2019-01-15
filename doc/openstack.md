# Deployment for openstack

## Prerequisites

- Clean VM with Centos7
- Installed openstack + contrail with contrail-ansible-deployer
  - Sample instances.yaml file presented below. Tokens @localip and @routerip must be changed before running contrail-ansible-deployer, passwords should be adjusted.

```yaml
provider_config:
  bms:
    ntpserver: "pool.ntp.org"
    nameserver: "8.8.8.8"
    ssh_user: "centos"
    ssh_private_key: "/home/centos/id_rsa"
    domainsuffix: "novalocal"

instances:
  kolla-aio:
    provider: "bms"
    ip: @localip
    roles:
      config_database:
      config:
      control:
      analytics_database:
      analytics:
      webui:
      openstack:
      vrouter:
      openstack_compute:

global_configuration:
  CONTAINER_REGISTRY: "opencontrailnightly"

contrail_configuration:
  CONTRAIL_VERSION: "latest"
  CLOUD_ORCHESTRATOR: "openstack"
  OPENSTACK_VERSION: "ocata"
  CONTROLLER_NODES: @localip
  LOG_LEVEL: "SYS_DEBUG"
  PHYSICAL_INTERFACE: "eth0"
  VROUTER_GATEWAY: @routerip
  AUTH_MODE: "keystone"
  AAA_MODE: None
  KEYSTONE_AUTH_ADMIN_PASSWORD: "contrail123"
  KEYSTONE_AUTH_HOST: @localip
  KEYSTONE_AUTH_URL_VERSION: "/v3"
  RABBITMQ_NODE_PORT: 5673
  JVM_EXTRA_OPTS: "-Xms1g -Xmx2g"
  CONFIG_NODEMGR__DEFAULTS__minimum_diskGB: 2
  DATABASE_NODEMGR__DEFAULTS__minimum_diskGB: 2

kolla_config:
  kolla_globals:
    network_interface: "eth0"
    api_interface: "eth0"
    neutron_external_interface: "eth0"
    kolla_external_vip_interface: "eth0"
    kolla_internal_vip_address: @localip
    contrail_api_interface_address: @localip
    enable_haproxy: "no"
    enable_barbican: "no"
    enable_ironic: "no"
    enable_ironic_notifications: "no"
    openstack_service_workers: 1
    openstack_release: "ocata"
    kolla_base_distro: "centos"
  kolla_passwords:
    keystone_admin_password: "contrail123"
  customize:
    nova.conf: |
      [libvirt]
      virt_type=qemu
      cpu_mode=none
    neutron.conf: |
      [DEFAULT]
      service_plugins=""

orchestrator_configuration:
  internal_vip: @localip
  keystone:
    version: "v3"
    password: "contrail123"
```

## Deployment

- On remote (target) machine clone [Contrail repository](https://github.com/Juniper/contrail) into $HOME/go/src/github.com/Juniper

  ```bash
  mkdir -p $HOME/go/src/github.com/Juniper
  cd !$
  git clone https://github.com/Juniper/contrail
  ```

- Run a deployment script

  ```bash
  cd contrail
  ./tools/deploy-for_openstack.sh
  ```

## Examples

### Simple ping

Log in to Openstack WebUI(horizon) http://@local.ip

For project admin:

- Upload tiny OS image, for example: CirrOS.

- Create flavor for your instances. For example:
  ```
  VCPUs: 1
  RAM (MB): 256
  Root Disk (GB): 1
  ```

- Create network.

- Create subnet, for example: 10.0.1.0/24.

- Launch two instances using prepared image and flavor in created network.

- Open Web console's of both instances and log in to machines.

- Retrieve IP addresses of both instances using command:
  ```bash
  ip a
  ```

- Try to make them ping each other.
