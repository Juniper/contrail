# Agent service

Agent watches for changes of specified API Server resources and runs specified action for each change.
Change might be creation, update or deletion of a resource.
Action performed on event might be generating files based on provided templates.

Current mechanism of acquiring event data is polling on REST API of Server.
Watching for events from etcd service is a feature during implementation.

## Requirements

Agent requires access to service which it pulls data from.
Depending on chosen operation mode it might be:
* API Server
* etcd with v3 API support

## Configuration

Service reads configuration from YAML file on path specified `--config-file` flag.
Required fields are defined in [source code](../pkg/agent/agent.go) as the `Config` structure.

Example configuration can be found [here](../sample/agent.yml).

## Running

Start Agent specifying configuration file path:

	contrail agent -c <config-file-path>

or You can start agent in server process

	contrail server -c <config-file-path> -a <agent-config-file>
