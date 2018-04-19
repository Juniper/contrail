# Watcher service

Watcher supplies etcd with data received from database replication protocols.
Currently there are drivers for two databases:
- PostgresSQL driver using logical streaming replication protocol based on `pgx` and `pgoutput`
libraries
- MySQL driver reading binary log using `mysqldump` and MySQL replication protocol.

## PostgreSQL
Watcher leverages new [PostgreSQL10 logical streaming](https://www.postgresql.org/docs/10/static/protocol-logical-replication.html) replication protocol to track database events.

Watcher has two phases of operation:
* `dump` - load state from transaction snapshot created during replication slot creation (not implemented yet)
* `sync` - use PostgreSQL logical replication with pgoutput logical decoding
to receive transaction events sent by database

### Requirements

* PostgreSQL 10 and above with following configuration:
  * `wal_level=logical`
* publication created on the server with name `watcherpub` (TODO(Michal) remove when watcher creates publication by itself)
  * `CREATE PUBLICATION watcherpub FOR ALL TABLES`


## MySQL
As in PostgreSQL case watcher has two phases of operation:
* `dump` - use `mysqldump` to dump existing MySQL data from beginning of binlog to latest state
* `sync` - use MySQL replication protocol to block and synchronize new events appended to binlog

Current implementation posses following shortcomings:
* Altering tables is currently not supported, because table information are cached.
* Only one `database.schema` can be specified.
* There is no whitelisting/blacklisting of database tables.
* Tables with multi-column primary keys are currently not supported.
* MySQL enum, set and bit types are not supported.

### Requirements

* `mysqldump` tool available on host machine
* MySQL available on specified URL with following configuration:
  * binlog_format=ROW
  * binlog_row_image=FULL
  * log_bin=<path-to-binary-logs>
  * server_id=<server-id>

### MySQL configuration
To run properly configured MySQL add following section to `/etc/mysql/my.cnf`: 

``` shell
[mysqld]
log_bin=/var/log/mysql/mysql-bin
server_id=1
```

Restart MySQL to apply changes: `service mysql restart`

## Running etcd
Watcher requires etcd with v3 API support to work properly.

In test environment etcd is run in an docker container. Required container is
initialized in [./tools/testenv.sh](../tools/testenv.sh) script which is executed
by calling `make testenv`.

## Configuration

Service reads configuration from YAML file on path specified `--config-file` flag.
Used configuration keys and their defaults can are defined [here](../pkg/watcher/service.go).

Example configuration can be found [here](../sample/watcher.yml).

Available database driver options are: `pgx` and `mysql` named after database
drivers used.


## Running

Start Watcher specifying configuration file path:

	contrail watcher -c <config-file-path>

or you can start watcher in server process:

	contrail server -c <config-file-path> -w <watcher-config-file>
