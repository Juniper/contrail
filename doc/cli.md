# API Server command line client

API Server CLI is contained in both `contrail` and `contrailcli` executables.
It consists of following commands:
- schema
- show
- list
- create
- set
- update
- sync
- rm
- delete

Find information about them invoking `help`:

	contrailcli <command> -h

## Configuration

CLI reads configuration from YAML file on path specified `--config-file` flag.
Required fields are defined in [source code](../pkg/agent/agent.go) as the `Config` structure.
Note that Agent-specific fields, such as `watcher` and `tasks` are not required - to be changed in the future.

Example configuration can be found [here](../integration/contrailcli.yml).  

## Running

Invoke CLI command specifying configuration file path, e.g.:

	contrailcli schema virtual_network -c <config-file-path>
