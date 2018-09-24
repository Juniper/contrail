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
## Security group demo scenario
* Prepare 3 terminals on a machine with Atom deployed: one for running the scenario script, and two for running Kubernetes pods.
* Go to the tools/security-group-demo/bootstrap directory and run the `bootstrap.sh` script.
  It requires 3 parameters: address of Atom, chosen namespace and directory to deploy the scripts. For example:
``` shell
cd ./tools/security-group-demo/bootstrap
./bootstrap.sh localhost:8082 sgdemo ../demo
```
* Go to the deployed scripts directory and run `setup.sh`. It will create a kubernetes namespace and a virtual network. For example:
``` shell
cd ../demo
setup.sh
```
* Go to the deployed scripts directory in the other two terminals and create two pods using `pod.sh`. Use different names for the pods. For example:
```shell
cd ./tools/security-group-demo/demo
./pod.sh pod1

# Do the same in the other terminal, using e.g. ./pod.sh pod2
```
The script will print the full name of the pods, such as "pod1-75f684d9f9-mndt4". You will need one of them in the next step.
* Obtain the name of the VMI associated to one of the pods using `get-vmi-name.sh`. It requires the name of the pod. For example:
``` shell
# In the deployed scripts directory
./get-vmi-name.sh pod1-75f684d9f9-mndt4
```
You will need the VMI name in the next step.
* Run the scenario script using `run.sh`. It requires the name of the VMI. For example:
``` shell
# In the deployed scripts directory
./run.sh pod1-75f684d9f9-mndt4__5d8703a2-c011-11e8-b1fe-fa163e4988da
```
* The script will allow/disallow traffic from/to the chosen pod. Attach the pods in the other two terminals. Use "ip a" to obtain their IP addresses.
  You can use ping and ./echo.sh with netcat to check if ICMP/TCP is blocked. For example, in one of the pods:
``` shell
./echo.sh
```
and in the other:
``` shell
# Assuming the IP of the first pod is 10.32.0.1.
ping 10.32.0.1
# Should reply "echo" if and only if TCP traffic is allowed:
nc 10.32.0.1 8080
```
