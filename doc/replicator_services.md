# Replicator services

Replicators are a temporary layer of backward compatibility. Used until all Contrail microservices are upgraded to use etcd instead of Cassandra and RabbitMQ.

**Source code:** [run.go](../pkg/cmd/contrail/run.go), [cassandra.go](../pkg/db/cassandra/cassandra.go)

## How it works

Replicators watch recursively on events in etcd database and pipes them out to RabbitMQ and Cassandra.

Config extract from [contrail.yml](../sample/contrail.yml):

```yaml
replication:
  cassandra:
    enabled: true
  amqp:
    enabled: true

cassandra:
  host: localhost
  port: 9042
  timeout: 3600s
  connect_timeout: 600ms

amqp:
  url: amqp://guest:guest@localhost:5672/
```
