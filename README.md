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

```bash
go get -u github.com/Juniper/contrail
```

Note that `go get -u github.com/Juniper/contrail/cmd/contrailutil` fails because we don't
commit generated code.

### Step3. Install dependencies

```bash
make deps
```

### Step4. Generate code

```bash
make generate
```

### Step5. Install code

```bash
make install
```

### Step6. Install test environment

```bash
# Setup test environment using Docker and setup DB
make testenv reset_db
```

Note that these commands use `docker` command and depending on your docker
configuration they may require root permissions.
See [Docker Documentation](https://docs.docker.com/install/linux/linux-postinstall/#manage-docker-as-a-non-root-user)
for more info.

## Try

- Run processes

    ```bash
    contrail -c sample/contrail.yml run
    ```

    Note that you can overwrite configuration parameters using environment variable with
    prefix "CONTRAIL_"

    For example CONTRAIL_DATABASE_DEBUG is overwriting database.debug value.

    ```bash
    CONTRAIL_DATABASE_DEBUG=true contrail -c sample/contrail.yml run
    ```

    Individual processes can be enabled or disabled using the configuration parameters.

- Run CLI

    ```bash
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
Developers should make sure download latest schema from [contrail-api-client](http://github.com/Juniper/contrail-api-client)

JSON version stored in public/schema.json

Templates for code generation based on this schema are stored in [tools/templates](tools/templates)
[Template configuration](tools/templates/template_config.yaml)
You can add your template on template_config.yaml.

## Testing

```bash
make test
```

You can print out full sql trace too.

```bash
CONTRAIL_DATABASE_DEBUG=true make test
```

## Commands

Repository holds source code for following CLI applications:

- `contrail` - contains API Server, [Agent](doc/agent.md) and [Sync](doc/sync.md) processes and [Cluster](doc/cluster.md) service
- `contrailcli` - contains [API Server command line client](doc/cli.md)
- `contrailschema` - code generator by schema definitions
- `contrailutil` - contains development utilities

Show possible commands of application:

```bash
contrail -h
```

Show detailed information about specific command:

```bash
contrail <command> -h
```

## Keystone Support

API Server supports Keystone V3 authentication and RBAC.
API Server has minimal Keystone API V3 support for standalone use case.
See a configuration [example](https://github.com/Juniper/contrail/blob/master/sample/contrail.yml#L61)

## How to contribute

- Follow [Openstack review process](https://docs.openstack.org/infra/manual/developers.html)
- Use [Tungsten Fabric Gerrit](https://review.opencontrail.org)
- Ensure that `make test lint` passes
- Comply to [Code review guidlines](REVIEW.md)

### Step1

Setup gerrit account. Sign CLA.

### Step2

Install git-review.

```bash
pip install git-review
```

### Step3

Send git review command.

```bash
git review
```

## Document

See [docs](./doc) folder.
