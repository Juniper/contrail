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
- Install [dep](https://github.com/golang/dep)
- Install [mysql](https://dev.mysql.com/doc/en/installing.html)
- Install [fpm](https://github.com/jordansissel/fpm), only required if building packages (described below)
  - Install [ruby](https://www.ruby-lang.org/en/documentation/installation/)
  - Install [rubygems](https://rubygems.org/pages/download)

## Retrieve the code (using go get)

``` shell
go get github.com/Juniper/contrail
```

## Generate Code

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

You need to run a local mysql instance running with test configuration.

It is expected that the root password is 'contrail123', you can set this on an existing installation
from the mysql prompt as follows:

``` shell
MariaDB [(none)]> ALTER USER 'root'@'localhost' IDENTIFIED BY 'contrail123';
```

Executing the script below, will drop the contrail_test schema if it exists, recreate it and initialise this schema

``` shell
./tools/reset_db.sh
```

At this point the tests can be executed:

``` shell
make test
```

Run integration tests:

``` shell
make integration
```

## Commands

Repository holds source code for following CLI applications:
- `contrail` - contains API Server, [Agent](doc/agent.md) and [API Server command line client][cli] 
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
