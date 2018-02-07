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

* `mysqldump` tool available on host machine
* MySQL available on specified URL with following configuration:
  * binlog_format=ROW
  * binlog_row_image=FULL
  * log_bin=<path-to-binary-logs>
  * server_id=<server-id>
* etcd available on specified URL with v3 API support

### MySQL configuration
TODO(daniel) Move this section to root `README.md` when tests require this configuration

To run properly configured MySQL add following section to `/etc/mysql/my.cnf`: 

``` shell
[mysqld]
log_bin=/var/log/mysql/mysql-bin
server_id=1
```

Restart MySQL to apply changes: `service mysql restart`

### Running etcd
TODO(daniel) Move this section to root `README.md` when tests require etcd

Etcd can be run with local installation or inside Docker container.

#### Local etcd

Download etcd 3.3.1 from [release page][etcd-releases], extract it and put `etcd` and `etcdctl` binaries
within system PATH, e.g:

``` shell
wget https://github.com/coreos/etcd/releases/download/v3.3.1/etcd-v3.3.1-linux-amd64.tar.gz
tar -zxf etcd-v3.3.1-linux-amd64.tar.gz
sudo mv etcd-v3.3.1-linux-amd64/etcd etcd-v3.3.1-linux-amd64/etcdctl /usr/local/bin/
rm -rf  etcd-v3.3.1-linux-amd64 etcd-v3.3.1-linux-amd64.tar.gz
```

Run etcd service with `./tools/run_etcd` script.

[etcd-releases]: https://github.com/coreos/etcd/releases/

#### Etcd in Docker

Run etcd in Docker container:

``` shell
docker run -d --name etcd -p 2379:2379 gcr.io/etcd-development/etcd:v3.3.1 etcd \
	--advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379
```

## Configuration

Service reads configuration from YAML file on path specified `--config-file` flag.
Required fields are defined in [source code](../pkg/watcher/service.go) as the `Config` structure.

Example configuration can be found [here](../sample/watcher.yml).  

## Running

Start Watcher specifying configuration file path:

	contrail watcher -c <config-file-path>
