# Intent Compilation service

Intent Compilation watches for changes of specified API Server resources and runs specified action for each change.
Change might be creation, update or deletion of a resource.
Action performed on event will be handled by platform specific plugins

Intent Compilation service watching for events from etcd service.

## Requirements

Intent Compilation requires access to service which it pulls data from.
It requires:

* API Server
* etcd server with v3 API support

## Configuration

Service reads configuration from YAML file on path specified `--config-file` flag.
Required fields are defined in [source code](../pkg/compilation/config/config.go)
as the `Config` structure.

Example configuration can be found [here](../sample/contrail.yml).

## Intents

Intent Compilation Service uses Intent interface:

```go
type Intent interface {
    basemodels.Object
    Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error
    GetObject() basemodels.Object
    GetDependencies() map[string]UUIDSet
    AddDependentIntent(i Intent)
    RemoveDependentIntent(i Intent)
}
```

Intent embeds resource and stores data calculated by Intent Compiler.
When Intent is created/updated/deleted then Dependency Processor is called. It gets all dependencies and evaluates them. Evaluate method is updating database using Api Server (currently using HTTP client, eventually gRPC).

## Dependency Processor

In order to keep all the resources updated we define reaction map which describes relations between Intents. This is sample reaction:

```yaml
  port_tuple:
    self:
      - service_instance
    virtual_machine_interface:
      - service_instance
    service_instance:
      - virtual_machine_interface
```

It means that when event on port_tuple is received, then we need to evaluate dependent resources of kind defined under key "self" (service_instance). However if event is related to service_instance, then if we call evaluate on port_tuple, then all related VMIs are going to be evaluated as well.

When event is received Dependency processor is called and scan resource's references and backreferences based on reaction map. Additionally methods AddDependentIntent and RemoveDependentIntent can be used to perform custom logic, beyond schema relations.
