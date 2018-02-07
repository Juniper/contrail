# Watcher service

Watcher supplies etcd with data from binary log of MySQL database using `mysqldump` and
MySQL replication protocol.

Watcher has two phases of operation:
* `dump` - use `mysqldump` to dump existing MySQL data from beginning of binlog to latest state
* `sync` - use MySQL replication protocol to block and synchronize new events appended to binlog

Current implementation posses following shortcomings:
* Altering tables is currently not supported, because table information are cached.
* Only one `database.schema` can be specified.
* There is no whitelisting/blacklisting of database tables.
* Tables with multi-column primary keys are currently not supported.
* MySQL enum, set and bit types are not supported.

## Requirements

* `Mysqldump` tool available on host machine
* MySQL available on specified URL with following configuration:
  * binlog-format: ROW
  * binlog-row-image: FULL
  * binlog row events version: v1 or v2
* etcd available on specified URL with v3 API support

## Configuration

Service reads configuration from YAML file on path specified `--config-file` flag.
Required fields are defined in [source code](../pkg/watcher/service.go) as the `Config` structure.

Example configuration can be found [here](../integration/watcher.yml).  

## Running

Start Watcher specifying configuration file path:

	contrail watcher -c <config-file-path>
