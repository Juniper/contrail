# Documentation index

## Preface

This is a design document for refactoring Contrail using Go.
The goal of this change is to provide:

- Simpler architecture and operation experience
- Higher performance
- Data collections
- Maintainability

## Architecture

This diagram shows the overall architecture.

API Server provides REST API and gRPC API.
Internal logic depending on RDBMS and utilize the power
of the relational model, so that developers can focus on implementing logic.
Sync process replicates RDBMS data to [etcd](https://github.com/etcd-io/etcd) using Replication mechanism.

Intent compiler get updates from etcd, evaluate configuration changes and dependency,
and generate config object from the API resource model.

Existing contrail processes such as Control Node, Device manager,
Kube Managers will be capable of watching for any updates in etcd.

![Architecture](./images/architecture.svg "Architecture")

## Process management

To simplify deployment,
we take a "Single binary" approach for Go-related project. We manage multiple goroutine
as micro-services, and run them based on a switch in configuration.
Each process communicate with each other using Service Interface defined above so that we can
switch internal functional call or gRPC call depending on where the other processes running on.

![Process model](./images/process.svg "Process")

See: [Related source code](../pkg/cmd/contrail/run.go)

## Configuration

We use [Viper](https://github.com/spf13/viper) for configuration management.
YAML is our default configuration format. Every configuration option must be configured via environment variables too with CONTRAIL_ prefix for docker based operation.

## Schema

We manage API Schema using the YAML format.
[schemas directory](../schemas) contains schemas.

In YAML format, the schema as following properties.

- id: unique schema ID
- extends: You can specify a list of abstract schema
- parents: parents resources
- references: many to many relations
- prefix: REST API prefix
- schema: JSON Schema

This is a sample schema.

```yaml
extends:
- base
id: virtual_network
parents:
  project:
    description: Virtual network is collection of end points (interface or ip(s) or
      MAC(s)) that can talk to each other by default. It is collection of subnets
      connected by implicit router which default gateway in each subnet.
    operations: CRUD
    presence: optional
plural: virtual_networks
prefix: /
references:
  network_ipam:
    $ref: types.json#definitions/VnSubnetsType
    description: Reference to network-ipam this network is using. It has list of subnets
      that are to be used as property of the reference.
    operations: CRUD
    presence: required
schema:
  properties:
    external_ipam:
      description: IP address assignment to VM is done statically, outside of (external
        to) Contrail Ipam. vCenter only feature.
      operations: CRUD
      presence: optional
      type: boolean
    fabric_snat:
      default: false
      description: Provide connectivity to underlay network by port mapping
      operations: CRUD
      presence: optional
      type: boolean
  required: []
  type: object
```

## Code generation

We have a makefile target to generate source codes or initial SQL definitions.

```bash
make generate
```

templates are stored in [template directory](../tools/templates)

List of templates are specified in [template configuration](../tools/templates/template_config.yaml)

We use [Pongo2](https://github.com/flosch/pongo2) which support a subset of Jinja2 template.

## Models

[../pkg/models] has Go structs for all resource objects. All processes must
use this model. Note that we should avoid the use of the level objects such as JSON strings /
map[string]interfaces{}.
We have model specific logic here. See more in GoDoc of this package.

## API Server

API Server provides REST API and gRPC for external orchestrators such as UI / OpenStack or
Kubernetes. [Source Code](../pkg/apisrv)
We use [echo framework](https://echo.labstack.com/) for HTTP Web server framework,
and use standard library for gRPC Server.
Internally, we dispatch any API requests for internal services which support
Service Interface described in the next chapter.

![API Server Internal Architecture](./images/api_process.svg "API Process Data")

See: [REST API documentation](rest_api.md)
See: [Authentication documentation](authentication.md)
See: [Access control documentation](policy.md)

## Service Interface & Chain

To decouple logic from specific implementation, we define "Service Interface".
The service interface support gRPC client API plus chain concept.
We apply HTTP middleware concept where we can inject multiple logic later in the process
pipeline.

We need to set "Next" services for each service implementation, and we are expecting
call Next() service on each call.

Next call example

```go
// CreateProject creates a project and ensures a default application policy set for it.
func (sv *ContrailTypeLogicService) CreateProject(
    ctx context.Context, request *services.CreateProjectRequest,
) (response *services.CreateProjectResponse, err error) {

    err = sv.InTransactionDoer.DoInTransaction(
        ctx,
        func(ctx context.Context) error {
            response, err = sv.Next().CreateProject(ctx, request)
            if err != nil {
                return err
            }

            return sv.ensureDefaultApplicationPolicySet(ctx, request.Project)
        })

    return response, err
}
```

## Commands

Repository holds source code for following CLI applications:

- `contrail` - contains core processes, such as: [API Server](rest_api.md), [Agent](agent.md), [Sync](sync.md) and [Cluster service](cluster.md)
- `contrailcli` - [API command line client](cli.md)
- `contrailschema` - code generator using schema definitions
- `contrailutil` - development utilities, such as [Convert](convert.md), `package`, `record_test`

Show possible commands of application:

```bash
contrail -h
```

Show detailed information about specific command:

```bash
contrail <command> -h
```

## Services

### DB Service

### Cache Service

### Contrail Service

### Type Specific logic service

## Intent Compiler

See: [Intent Compiler documentation](./intent_compilation.md)

## Sync Process

See: [Sync process documentation](./sync.md)

## Kubernetes support

See: [Kubernetes support documentation](./k8s.md)
