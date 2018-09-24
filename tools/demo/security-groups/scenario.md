# Security group demo scenario
* Prepare 3 terminals on a machine with Atom deployed: one for running the scenario script, and two for running Kubernetes pods.
* Go to the `tools/demo/security-groups/bootstrap` directory and run the `bootstrap.sh` script.
  It requires 3 parameters: address of Atom, chosen namespace and directory to deploy the scripts. For example:
``` shell
cd ./tools/demo/security-groups/bootstrap
./bootstrap.sh localhost:8082 sgdemo ../scripts
```
* Go to the deployed scripts directory and run `setup.sh` as root. It will create a kubernetes namespace and a virtual network. For example:
``` shell
cd ../scripts
setup.sh
```
* Go to the deployed scripts directory in the other two terminals and create two pods using `pod.sh` as root. Use different names for the pods. For example:
```shell
cd ./tools/demo/security-groups/scripts
./pod.sh pod1

# Do the same in the other terminal, using e.g. ./pod.sh pod2
```
The script will print the full name of the pods, such as "pod1-75f684d9f9-mndt4". You will need one of them in the next step.
* Obtain the name of the VMI associated to one of the pods using `get-vmi-name.sh`. It requires the name of the pod. For example:
``` shell
# In the deployed scripts directory
./get-vmi-name.sh pod1-75f684d9f9-mndt4
```
It will print the VMI name. You will need it in the next step.
* Run the scenario script `run.sh` as root. It requires the name of the VMI. For example:
``` shell
# In the deployed scripts directory
./run.sh pod1-75f684d9f9-mndt4__5d8703a2-c011-11e8-b1fe-fa163e4988da
```
* The script will allow/disallow traffic from/to the chosen pod. Attach the pods in the other two terminals. Use `ip a` to obtain their IP addresses.
  You can use ping and `echo.sh` with netcat to check if ICMP/TCP is blocked. For example, in one of the pods:
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
