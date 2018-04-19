# API Server command line client

API Server CLI is contained in both `contrail` and `contrailcli` executables.
It consists of following commands:
- schema
- show
- list
- set
- sync
- rm
- delete

Find information about them invoking `help`:

	contrailcli <command> -h

## Configuration

CLI reads configuration from YAML file on path specified `-c` or `--config` flag.
You can also set environment variable `CONTRAIL_CONFIG` for path.
Example configuration can be found [here](../sample/cli.yml).  

```
export CONTRAIL_CONFIG=./sample/cli.yml
```

## Running

Usage examples for available commands are presented below.

### Schema command

Schema command shows schema for specified resource.

	contrailcli schema virtual_network 

[Schema command output](../pkg/cmd/contrailcli/testdata/virtual_network_schema.yml)
### Show command

Show command shows data of specified resource.

	contrailcli show virtual_network first-uuid 

[Show command output](../pkg/cmd/contrailcli/testdata/virtual_network_showed.yml)

Invoke command with empty schema identifier in order to show possible usages.

	contrailcli show 

Show command output (shortened):
```
Show command possible usages:

contrail show base $UUID
contrail show has_node $UUID
contrail show has_status $UUID
contrail show access_control_list $UUID
contrail show address_group $UUID
contrail show alarm $UUID
contrail show alias_ip_pool $UUID
contrail show alias_ip $UUID
contrail show analytics_node $UUID
contrail show api_access_list $UUID
contrail show application_policy_set $UUID
contrail show bgp_as_a_service $UUID
contrail show bgp_router $UUID
contrail show bgpvpn $UUID
contrail show bridge_domain $UUID
...
```

### List command

List command lists data of specified resources.

There are also multiple parameters available, such as filters.
Display them with `contrailcli list -h` command.

	contrailcli list virtual_network

[List command output](../pkg/cmd/contrailcli/testdata/virtual_networks_listed.yml)

Invoke command with empty schema identifier in order to show possible usages.

	contrailcli list -c integration/contrailcli.yml

List command output (shortened):
```
List command possible usages:

contrail list base
contrail list has_node
contrail list has_status
contrail list access_control_list
contrail list address_group
contrail list alarm
contrail list alias_ip_pool
contrail list alias_ip
contrail list analytics_node
contrail list api_access_list
contrail list application_policy_set
contrail list bgp_as_a_service
contrail list bgp_router
contrail list bgpvpn
contrail list bridge_domain
...
```


### Sync command

Sync command create or update resources with data defined in given YAML file.
It creates new resource for every not already existing resource.

	contrailcli sync pkg/cmd/contrailcli/testdata/virtual_networks_update.yml

[Input file content](../pkg/cmd/contrailcli/testdata/virtual_networks_update.yml)

### Set command

Set updates properties of specified resource.

	contrailcli set virtual_network first-uuid "external_ipam: true" -c integration/contrailcli.yml

[Set command output](../pkg/cmd/contrailcli/testdata/virtual_networks_set_output.yml)

Invoke command with empty schema identifier in order to show possible usages.

	contrailcli set

### Rm command

Rm removes a resource with specified UUID.

	contrailcli rm virtual_network second-uuid
	
Rm command output is empty on successful operation.

Invoke command with empty schema identifier in order to show possible usages.

	contrailcli rm 

### Delete command

Delete removes resources specified in given YAML file.

	contrailcli delete pkg/cmd/contrailcli/testdata/virtual_networks.yml
	
[Input file content](../pkg/cmd/contrailcli/testdata/virtual_networks.yml)

Delete command returns no output.
