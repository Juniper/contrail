# Intent Compilation service

Intent Compilation watches for changes of specified API Server resources and
runs specified action for each change.
Change might be creation, update or deletion of a resource.
Action performed on event will be handled by platform specific plugins

Intent Compilation service watching for events from etcd service.

## Requirements

Intent Compilation requires access to service which it pulls data from.
It requires
* API Server
* etcd server with v3 API support

## Configuration

Service reads configuration from YAML file on path specified `--config-file` flag.
Required fields are defined in [source code](../pkg/compilation/config/config.go)
as the `Config` structure.

Example configuration can be found [here](../sample/compilation.yml).

## Running

Start Intent Compilation specifying configuration file path:

	contrail compilation -c <config-file-path>

