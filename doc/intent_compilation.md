# Intent Compilation service

Intent Compiler watches for changes of higher-level resources in API Server
and creates/updates/deletes lower-level resources for them.

It watches for events from `etcd`.
Sometimes, it fetches some extra information it needs by calling API Server directly.
All modifications of resources are done by calling API Server (currently using HTTP client, eventually gRPC).

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

## Request handling

When an event from `etcd` arrives, it is sent to a job queue,
which has multiple workers for processing events in parallel.

When a worker takes the event, it calls `compilation/logic.Service` on it.

The logic service associates the event with an `Intent` in its cache
(by creating an `Intent` if there is none
or by replacing/deleting the existing `Intent` if one exists)
and calls the `DependencyProcessor` on the `Intent`.

## Intent cache

Intent Compiler has its own cache, which it uses for:
- maintaining the state of resources previously received from etcd. This is necessary for comparing the new contents of resources with the old ones.
- maintaining the relationships between resources. This is necessary because, in events for a given resource, `etcd`:
  - does not provide children and backrefs;
  - provides UUIDs of refs, but does not provide their contents;
  - does not provide contents of other undeclared kinds of related resources (e.g. referred security groups);
- maintaining additional information about resources, calculated from their or related resources' contents.

## Intents

An `Intent` is a wrapper over a resource
that also contains information calculated from it and other related resources.

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

When the logic service creates/updates/deletes an `Intent`, it calls the `DependencyProcessor`.
The `DependencyProcessor` calls `Evaluate()` on the `Intent` and, if necessary, on dependent `Intent`s.

The main Intent Compilation logic (creating/updating/deleting lower-level resources)
is usually implemented in `Evaluate()`.

## Dependency Processor

When event is received Dependency processor is called and scan resource's references and backreferences based on reaction map. Additionally methods `AddDependentIntent` and `RemoveDependentIntent` can be used to perform custom logic, beyond schema relations.

In order to keep all the resources in sync with each other, we define a reaction map which describes relations between Intents. This is a sample reaction:

```yaml
  port_tuple:
    self:
      - service_instance
    virtual_machine_interface:
      - service_instance
    service_instance:
      - virtual_machine_interface
```

It means that when an event on `port_tuple` is received, then we need to evaluate dependent resources of kind defined under key `"self"` (`service_instance`).
However, if evaluation of the resource is caused by a `service_instance` resource, then, for the `port_tuple` event, all related `virtual_machine_interface`s are going to be evaluated as well.

## Implementing a new kind of resource

- Add an `Intent` type for the resource.
- Add Create/Update/Delete methods for the resource in `compilation/logic.Service`. These should deal with creating/updating/deleting the `Intent` in the cache and call `EvaluateDependencies`.
- Add unit tests that mock `WriteService`, put the other `Intent`s into the cache, call the logic service, and check
  - the `Intent` in the cache;
  - whether the necessary calls to the `WriteService` were made.
- Add an `Evaluate` method for the `Intent` type.
- Add integration tests for the added functionality.
