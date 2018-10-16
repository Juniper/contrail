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
  - does not provide the contents of other undeclared kinds of related resources (e.g. referred security groups);
- maintaining additional information about resources, calculated from their or related resources' contents.

A resource's dependencies are calculated when it is added to the cache.
By default, references, backreferences, parents, and children are recorded as dependencies.
Additionally, methods `AddDependentIntent` and `RemoveDependentIntent` can be used
to add custom dependencies that do not correspond to relations specified in the schema (e.g. referred security groups).

Backreferences and children are resolved when a resource that has a parent in the cache or refers to a resource in the cache is added.

The idea is to have all the resources of the kinds relevant to Intent Compiler in the cache -
a copy of `etcd` contents for those kinds of resources.

Note that Intent Compiler does not fetch the current state of resources when it (re)starts (it starts with an empty cache).
Additionally:
- When an update/delete event for a resource that's not in the cache is received,
  the IC does not know if the lower-level resources for it have already been created.
- When a resource is put into the cache and its parent or a resource it refers to is not in the cache,
  the referred resource won't have a child/backreference to it even after getting added to the cache.
This means that, in general, the Intent Compiler does not work after a restart.

## Intents

An `Intent` is a wrapper over a resource
that additionally stores information calculated from it and other related resources.

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

## Dependency Processor

When the logic service creates/updates/deletes an `Intent`, it calls the `DependencyProcessor`.
The `DependencyProcessor` calls `Evaluate()` on the `Intent` and, if necessary, on dependent `Intent`s.

`Evaluate()` usually contains the main Intent Compilation logic for the resource (creating/updating/deleting lower-level resources).

The `DependencyProcessor` decides whether to `Evaluate` a dependency based on a static reaction map,
which describes relations between Intents. This is a sample reaction:

```yaml
  port_tuple:
    self:
      - service_instance
    virtual_machine_interface:
      - service_instance
    service_instance:
      - virtual_machine_interface
```

It means that, when an event on a `port_tuple` is received, it needs to evaluate the kinds of dependent resources defined under `"self"` (in this case, all related `service_instance`s).
However, if evaluation of the resource is caused by a `service_instance` resource, then, for the `port_tuple` event, all related `virtual_machine_interface`s are going to be evaluated instead.

## Implementing a new kind of resource

- Add an `Intent` type for the resource.
- Add `compilation/logic.Service` methods for the resource. These should deal with creating/updating/deleting the `Intent` in the cache and calling `EvaluateDependencies`.
- Add unit tests that:
  - mock `WriteService` (and `ReadService`);
  - put related `Intent`s into the cache;
  - call the logic service;
  - check the resulting `Intent` in the cache;
  - check whether the necessary calls to the `WriteService` were made.
- Add an `Evaluate` method for the `Intent` type. It should ensure that the lower-level resources are in sync with the resource in the `Intent`.
- Add integration tests for the introduced functionality.
