# [POC] Go code base for Contrail projects - DNM

This repository holds Go implementation for Contrail projects.
The goal of this project is
to realize Go-based implemenation & etcd based such as Kubernetes
in order to improve performance, scale and ease operation.

We are planning to add following sub components.

- API Server ( python based VNC API Server equivalent)
- Sync (ifmap, rabbitMQ realated code equivalent but depends on etcd)
- Agent (SchemaTransformer, Device Manager equivalent)
- Code generation tool (generateDS equivalent)

Currently, this project is
POC stage so any external/internal API or design subject to change up
to community discussion.

## Development setup

### Step1. Install Go and docker

https://golang.org/doc/install

### Step2. Go get contrailutil

``` shell
go get -u github.com/Juniper/contrail/cmd/contrailutil
```

Note that go get -u github.com/Juniper/contrail fails because we don't
commit genreated code.

### Step3 Install dependency and testenv

``` shell
make deps
```

### Step4 Install Test DB

```
# setup testenv using docker
make testenv
# you need wait db process up
make reset_db
```

### Step5 Generate code

``` shell
make generate
```

### Step6 Install code

``` shell
make install
```

### Try

- Run Server
```
contrail -c sample/server.yml server
```

Note that you can overwrite configuraion parameters using environment variable with
prefix "CONTRAIL_"

For example CONTRAIL_DATABASE_DEBUG is overwriting database.debug value.
``` shell
CONTRAIL_DATABASE_DEBUG=true contrail -c sample/server.yml server
```

- Run CLI

```
export CONTRAIL_CONFIG=sample/cli.yml
# Show Schema
contrailcli schema virtual_network
# Create resources
contrailcli sync sample/sample_resource.yml
# List resources
contrailcli list virtual_network --detail
# Delete resources
contrailcli delete sample/sample_resource.yml
```

For more cli command see [CLI Usage](doc/cli.md),

- Run Agent

```
contrail agent -c sample/agent.yml
```

For more agent command see [Agent Usage](doc/agent.md),

## Schema Files

Note that schema stored here is just a cache for helping development.
Developers should make sure download latest schema from http://github.com/Juniper/contrail-api-client

JSON version stored in public/schema.json

Templates for code generation based on this schema are stored in [tools/templates](tools/templates)
[Template configuration](tools/templates/template_config.yaml)
You can add your template on template_config.yaml.

## Testing

``` shell
make test
```

You can print out full sql trace too.

``` shell
CONTRAIL_DATABASE_DEBUG=true make test
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


## Keystone Support

API Server supports Keystone V3 authentication and RBAC.
API Server has minimal Keystone API V3 support for standalone use case.
See a configuration [example](sample/server.yml)

## How to contribute

- Apply lint tools
- Follow best practices
  - comply to [Effective Go](https://golang.org/doc/effective_go.html)
  - comply to [Code review comments](https://github.com/golang/go/wiki/CodeReviewComments)
  - keep `make lint` output clean

We follow openstack way of review. https://docs.openstack.org/infra/manual/developers.html
This is our review system. https://review.opencontrail.org

### Step1.

Setup gerrit account. Sign CLA.

### Step2.

Install git-review

```
pip install git-review
```

### Step3.

Send git review command.
```
git review
```

## Document

see [docs](./docs) folder
