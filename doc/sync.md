# Sync service

Sync supplies etcd with data received from database via PostgreSQL logical streaming replication protocol.
Sync requires etcd with v3 API support to work properly. Single Sync process can replicate data from one database.

## Service configuration

Configuration keys and their defaults are defined [here](../pkg/sync/service.go).

## PostgreSQL

Sync leverages new [PostgreSQL10 logical streaming](https://www.postgresql.org/docs/10/static/protocol-logical-replication.html) replication protocol to track database events.

Sync has two phases of operation:

- `dump` - load state from transaction snapshot created during replication slot creation
- `sync` - use PostgreSQL logical replication with `pgoutput` logical decoding to receive transaction events sent by database

### PostgreSQL requirements

- PostgreSQL 10 and above with following configuration:
  - `wal_level=logical`

### Automatic failover

Sync is run in retry loop, which means that in case of DB failover it tries to reconnect to a new master.
If Sync connects to a DB node that is currently in recovery node, then it retries, because logical replication is not working in such case.

To demonstrate Sync behavior in case of DB failover we can use following scenario:

1. Run ATOM with Postgres configuration.
   ```
   $ contrail run -c sample/contrail.yml
   ```

2. Wait until dump phase passes. After it's done Sync will log messages like:
   ```
   INFO[0053] Sending standby status    logger=postgres-watcher receivedLSN=0/BF98FE8 savedLSN=0/BF98FE8
   ```

3. In another terminal inspect patroni cluster state. In testenv you can use:
   ```
   $ docker exec contrail_haproxy patronictl list

   +-------------+--------------+----------+--------+---------+-----------+
   |   Cluster   |    Member    |   Host   |  Role  |  State  | Lag in MB |
   +-------------+--------------+----------+--------+---------+-----------+
   | testcluster | 9e012d8f9829 | 10.0.4.4 |        | running |         0 |
   | testcluster | f91edc7c34a8 | 10.0.4.5 | Leader | running |         0 |
   +-------------+--------------+----------+--------+---------+-----------+
   ```

   From it's outputs we can identify leader node which is named `f91edc7c34a8`.

4. Restart leader node triggering automatic failover. Then Sync will try to reconnect until recovery mode passes.

5. After few tries, Sync successfully reconnects, dumps the database and starts sending standby statuses like in step 2.
