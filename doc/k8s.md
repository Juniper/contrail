# Deployment for kubernetes (k8s)

## Prerequisites

- Clean VM with Centos7
- Installed k8s + contrail with contrail-ansible-deployer
  - Sample instances.yaml file presented below. Tokens @localip and @routerip must be changed before running contrail-ansible-deployer, passwords should be adjusted.

```yaml
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

- On remote (target) machine clone [Contrail repository](https://github.com/Juniper/contrail) into $HOME/go/src/github.com/Juniper

  ```bash
  mkdir -p $HOME/go/src/github.com/Juniper
  cd !$
  git clone https://github.com/Juniper/contrail
  ```

- Run a deployment script

  ```bash
  cd contrail
  ./tools/deploy-for_k8s.sh
  ```

## Examples

### Simple ping

Run two pods in different terminals (default namespace can be used):

```bash
kubectl run -i --tty busybox-one --image=busybox -- sh
```

```bash
kubectl run -i --tty busybox-two --image=busybox -- sh
```

You can check their IP addresses by typing
```bash
ip a
```

And then you can try to make them ping each other.

### Intra VN ping

Run the script located in [./tools/demo/intra-vn-ping/run.sh](../tools/demo/intra-vn-ping/run.sh)
on target machine. Last command located in script launches new pod in that terminal.

Next run second pod in another terminal

```bash
kubectl run -i --tty busybox-one --image=busybox --namespace blue -- sh
```

Now you can check their IP addresses by typing

```bash
ip a
```

And then you can try to make them ping each other.

### Logical router ping

Run the script located in [./tools/demo/logical-router-ping/run.sh](../tools/demo/logical-router-ping/run.sh)
on target machine. Last command located in script launches new pod in that terminal.

Next run second pod in another terminal

```bash
sudo kubectl run -i --tty busybox-one --image=busybox --namespace blue -- sh
```

Now you can check their IP addresses and then ping the same way
as described in "Intra VN ping"

### Security groups

- Prepare 3 terminals on a machine with Contrail deployed: one for running the scenario script, and two for running Kubernetes pods.
- Go to the `tools/demo/security-groups` subdirectory of the Contrail repository and run `setup.sh` as root.
  It will create a Kubernetes namespace and a virtual network. For example:

  ```bash
  # In the Contrail repository
  cd ./tools/demo/security-groups/
  sudo ./setup.sh
  ```

- Go to the same directory in the other two terminals and create two pods using `pod.sh` as root. Use different names for the pods. For example:

  ```bash
  cd ./tools/demo/security-groups/
  sudo ./pod.sh pod1

  # Do the same in the other terminal, using e.g. ./pod.sh pod2
  ```

  The script will print the full name of the pods, such as "pod1-75f684d9f9-mndt4". You will need one of them in the next step.

- Run the scenario script `run.sh`. It requires the name of the pod. For example:

  ```bash
  # In the security groups demo directory
  ./run.sh pod1-75f684d9f9-mndt4
  ```

- The script will allow/disallow traffic from/to the chosen pod. Attach the pods in the other two terminals. Use `ip a` to obtain their IP addresses.
  You can use ping and `echo.sh` with netcat to check if ICMP/TCP is blocked. For example, in one of the pods:

  ```bash
  ./echo.sh
  ```

  and in the other:

  ```bash
  # Assuming the IP of the first pod is 10.32.0.1.

  # Should reply if and only if ICMP traffic is allowed:
  ping 10.32.0.1

  # Should print "echo" if and only if TCP traffic is allowed:
  nc 10.32.0.1 8080
  ```

### Floating IP ping

Follow steps described in [./tools/demo/floating-ip-ping/scenario.md](../tools/demo/floating-ip-ping/scenario.md)
