# API Server command line client

API Server CLI is contained in `contrailcli` executable.
It consists of following commands:

- delete
- list
- rm
- schema
- set
- show
- sync

## Schema command

Schema command shows schema for specified resource.

```bash
contrailcli schema virtual_network
```

[Schema command output](../pkg/cmd/contrailcli/testdata/virtual_network_schema.yml)

## Show command

Show command shows data of specified resource.

```bash
contrailcli show virtual_network first-uuid
```

[Show command output](../pkg/cmd/contrailcli/testdata/virtual_network_showed.yml)

Invoke command with empty schema identifier in order to show possible usages:

```bash
contrailcli show
```

## List command

List command lists data of specified resources.

There are also multiple parameters available, such as filters.
Display them with `contrailcli list -h` command.

```bash
contrailcli list virtual_network
```

[List command output](../pkg/cmd/contrailcli/testdata/virtual_networks_listed.yml)

Invoke command with empty schema identifier in order to show possible usages:

```bash
contrailcli list -c integration/contrailcli.yml
```

## Sync command

Sync command creates or updates resources with data defined in given YAML file.
It creates new resource for every not already existing resource.

```bash
contrailcli sync pkg/cmd/contrailcli/testdata/virtual_networks_update.yml
```

[Input file content](../pkg/cmd/contrailcli/testdata/virtual_networks_update.yml)

## Set command

Set updates properties of specified resource.

```bash
contrailcli set virtual_network first-uuid "external_ipam: true" -c integration/contrailcli.yml
```

[Set command output](../pkg/cmd/contrailcli/testdata/virtual_networks_set_output.yml)

Invoke command with empty schema identifier in order to show possible usages:

```bash
contrailcli set
```

## Rm command

Rm removes a resource with specified UUID.

```bash
contrailcli rm virtual_network second-uuid
```

Rm command output is empty on successful operation.

Invoke command with empty schema identifier in order to show possible usages:

```bash
contrailcli rm
```

### Delete command

Delete removes resources specified in given YAML file.

```bash
contrailcli delete pkg/cmd/contrailcli/testdata/virtual_networks.yml
```

[Input file content](../pkg/cmd/contrailcli/testdata/virtual_networks.yml)

Delete command returns no output.
