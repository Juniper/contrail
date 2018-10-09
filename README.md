# [POC] Go code base for Contrail projects


This repository holds Go implementation for Contrail projects.
The goal of this project is
to realize Go-based implementation & etcd based such as Kubernetes
in order to improve performance, scale and ease operation.

We are planning to add following sub components.

- API Server ( python based VNC API Server equivalent)
- Sync (ifmap, rabbitMQ related code equivalent but depends on etcd)
- Agent (SchemaTransformer, Device Manager equivalent)
- Code generation tool (generateDS equivalent)

Currently, this project is
POC stage so any external/internal API or design subject to change up
to community discussion.

## Development setup

### Step1. Install Go and docker

- [golang.org/doc/install](https://golang.org/doc/install)
- [docs.docker.com/install](https://docs.docker.com/install/)

### Step2. Go get contrail

``` shell
go get -u github.com/Juniper/contrail
```

Note that `go get -u github.com/Juniper/contrail/cmd/contrailutil` fails because we don't
commit generated code.

### Step3. Install dependencies

``` shell
make deps
```

### Step4. Generate code

``` shell
make generate
```


### Step5. Install code

``` shell
make install
```

### Step6. Install test environment

``` shell
# setup testenv using docker
make testenv
# you need wait db process up
make reset_db
```

Note that these commands use `docker` command and depending on your docker
configuration they may require root permissions.
See [Docker Documentation](https://docs.docker.com/install/linux/linux-postinstall/#manage-docker-as-a-non-root-user)
for more info.

## Try

- Run processes

    ```
    contrail -c sample/contrail.yml run 
    ```

    Note that you can overwrite configuration parameters using environment variable with
    prefix "CONTRAIL_"

    For example CONTRAIL_DATABASE_DEBUG is overwriting database.debug value.

    ``` shell
    CONTRAIL_DATABASE_DEBUG=true contrail -c sample/contrail.yml run
    ```

    Individual processes can be enabled or disabled using the configuration parameters.

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
- `contrail` - contains API Server, [Agent](doc/agent.md) and [Sync](doc/sync.md)
processes and [Cluster](doc/cluster.md) service
- `contrailcli` - contains [API Server command line client][cli]
- `contrailschema` - code generator by schema definitions
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
See a configuration [example](https://github.com/Juniper/contrail/blob/master/sample/contrail.yml#L61)

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

Install git-review.

```
pip install git-review
```

### Step3.

Send git review command.
```
git review
```

## Document

See [docs](./doc) folder.
