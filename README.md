# Go code base for contrail projects

![CircleCI](https://circleci.com/gh/Juniper/contrail.svg?style=svg&circle-token=b744fe7f84003a898e897e0e4fe335e1e69944fd)
[![Coverage Status](https://coveralls.io/repos/github/Juniper/contrail/badge.svg?t=kKzcsv)](https://coveralls.io/github/Juniper/contrail)

## Important principal

- Apply lint tools
- Go get must simply work
- Follow best practices
  - comply to [Effective Go](https://golang.org/doc/effective_go.html)
  - comply to [Code review comments](https://github.com/golang/go/wiki/CodeReviewComments)
  - keep `make lint` output clean

## Build pre-requisites

The following software is required to build this project:

- Install [git](https://www.atlassian.com/git/tutorials/install-git)
- Install [go](https://golang.org/doc/install)
- Install [mysql](https://dev.mysql.com/doc/en/installing.html)
- Install [fpm](https://github.com/jordansissel/fpm), only required if building packages (described below)
  - Install [ruby](https://www.ruby-lang.org/en/documentation/installation/)
  - Install [rubygems](https://rubygems.org/pages/download)
- Run `make deps` to acquire development dependencies  

## Retrieve the code (using go get)

``` shell
go get github.com/Juniper/contrail
```

## Generate Code

### Setup protoc

see https://github.com/grpc/grpc/blob/master/INSTALL.md

Install protoc for go code generation

### Use make generate

``` shell
make generate
```

Templates are stored in [tools/templates](tools/templates)
[Template configuration](tools/templates/template_config.yaml)
You can add your template on template_config.yaml.

## Schema Files

Note that schema stored here is just a cache for helping development.
Developers should make sure download latest schema from http://github.com/Juniper/contrail-api-client

JSON version stored in public/schema.json

## Testing

In order to run tests you need to configure and start local MySQL and etcd, as described below.

Running tests:

``` shell
make test
```

### MySQL configuration

It is expected that the root password is `contrail123`, you can set this on an existing installation
from the mysql prompt as follows:

``` shell
MariaDB [(none)]> ALTER USER 'root'@'localhost' IDENTIFIED BY 'contrail123';
```

Add following section to `/etc/mysql/my.cnf`: 

``` shell
[mysqld]
log_bin=/var/log/mysql/mysql-bin
server_id=1
```

Restart MySQL to apply changes: `service mysql restart`

Executing the script below, will drop the contrail_test schema if it exists, recreate it and initialise this schema

``` shell
./tools/reset_db.sh
```

### Running etcd

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

Run etcd service with `make run_etcd`.

[etcd-releases]: https://github.com/coreos/etcd/releases/

#### Etcd in Docker

Run etcd in Docker container:

``` shell
docker run -d --name etcd -p 2379:2379 gcr.io/etcd-development/etcd:v3.3.1 etcd \
	--advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379
```

## Commands

Repository holds source code for following CLI applications:
- `contrail` - contains API Server, [Agent](doc/agent.md), [Watcher](doc/watcher.md)
and [API Server command line client][cli] 
- `contrailcli` - contains [API Server command line client][cli]
- `contrailutil` - contains development utilities

Show possible commands of application:

``` shell
contrail -h
```

Show detailed information about specific command:

``` shell
contrail <command> -h
```

[cli]: doc/cli.md

## API Server

API Server is shipped within `contrail` executable.
You can run API server using following command:

``` shell
go run cmd/contrail/main.go server -c packaging/apisrv.yml
```

### Keystone Support

API Server supports Keystone V3 authentication and RBAC.
API Server has minimal Keystone API V3 support for standalone use case.
See a configuration example in tools/test_config.yml

### More

Find out more about API Server:
- [Authentication](doc/authentication.md)
- [Policy](doc/policy.md)
- [REST API](doc/rest_api.md)

## Binary

Deb, RPM and Binaries are stored in release page.

See [releases](https://github.com/Juniper/contrail/releases)

## Packaging

Build the packages:

``` shell
make package
```
