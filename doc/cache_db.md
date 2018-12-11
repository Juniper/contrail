# Cache Database

**Source code:** [cachedb.tmpl](../tools/templates/contrail/cachedb.tmpl), [cache.go](../pkg/db/cache/cache.go)

Cache Database is an in memory linked list of `service.Event` nodes. Does not persist after restart and always must be recreated. Cache is configurable in `config.yaml`.

The purpose of Cache Database is to store resources in memory and send push notifications to clients on every change in resources.

## Configuration

[Example config](../sample/contrail.yml)

```yaml
cache:
  enabled: true
  timeout: 10s
  max_history: 100000  # how long revision deleted event preserved.
  cassandra:
    enabled: true  # listen on events from cassandra
  etcd:
    enabled: true  # listen on events from etcd
```

## How it works

Cache Database listen on events from cassandra and etcd database or events created by API server.

On new event, if resource already exists, old value is retrieve with all related resources like `back references` and `children`. Then new node is being created with updated resource. Old relationships are evaluated and new node replace old one.

If resource does not exist in cache, will be added as new node with calculated dependencies. On delete, if resource exists in cache the node will be removed.

At last Cache Database sends push notification about newly created node through websocket to listening clients.
