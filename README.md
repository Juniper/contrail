# Go code base for contrail projects

![CircleCI](https://circleci.com/gh/Juniper/contrail.svg?style=svg&circle-token=b744fe7f84003a898e897e0e4fe335e1e69944fd)
[![Coverage Status](https://coveralls.io/repos/github/Juniper/contrail/badge.svg?t=kKzcsv)](https://coveralls.io/github/Juniper/contrail)

## Important principal

- Apply lint tools
- Go get must simply work
- Follow best practices
  Effective Go: https://golang.org/doc/effective_go.html

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

At this point the tests can be executed

``` shell
make test
```

## API Server

You can run API server using this command.

``` shell
go run cmd/contrail/main.go server -c packaging/apisrv.yml
```

### Keystone Support

API Server supports Keystone V3 authentication and RBAC.
API Server has minimal Keystone API V3 support for standalone use case.
See a configuration example in tools/test_config.yml

## Binary

Deb, RPM and Binaries are stored in release page.

See [releases](https://github.com/Juniper/contrail/releases)

## Commands

- contrail  command for running intent api server/intent compiler etc

``` Shell
Contrail command

Usage:
  contrail [flags]
  contrail [command]

Available Commands:
  help        Help about any command
  server      Start API Server

Flags:
  -c, --config string   Configuration File
  -h, --help            help for contrail

Use "contrail [command] --help" for more information about a command.
```

- contrailutil utility command for help developments

``` shell
Contrail Utility Command

Usage:
  contrailutil [flags]
  contrailutil [command]

Available Commands:
  generate    generate code from schema
  help        Help about any command
  package     make a deb and rpm package

Flags:
  -h, --help   help for contrailutil
```

## Packaging

Build the packages

``` shell
make package
```
