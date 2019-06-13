# Cluster service

Cluster manages clustering of contrail/orchestrator components, currently cluster service
is triggered by agent service.

Agent service does the following:

- Watches for changes in contrail_cluster resource. For each change in resource. Change can be creation, update or deletion of a contrail_cluster resource.
- Creates the contrail-cluster.yml input config file for cluster service.
- Runs cluster service(one-shot) with the input config file(contrail-cluster.yml).

Cluster service does the following:

- Creates the input file for the provisioning tool (ansible/helm).
- Triggers the ansible playbook or helm chart
- Updates the provisioning status/provisioning logs in the contrail_cluster object.

TODO: Run cluster service independent of agent service as a daemon and make it watch for
      events in etcd service to trigger provisioning tool.

## Requirements

Agent service with polling on contrail_cluster resource.
Cluster service requires access to contrail_cluster resource.

## Configuration

Service reads configuration from YAML file on path specified `--config-file` flag.
Required fields are defined in [source code](../pkg/cluster/cluster.go) as the `Config` structure.

Example configuration template can be found [here](../pkg/cluster/configs/contrail-cluster-config.tmpl),
Which will be used by agent to generate the cluster config file.

## Running

Start cluster service in one-shot mode by specifying configuration file path:

```bash
contrailgo deploy -c config-file-path
```
