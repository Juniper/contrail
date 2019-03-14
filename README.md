# [POC] Go code base for Contrail projects

This repository holds Go implementation for Contrail projects.
The goal of this project is
to realize Go-based implementation & etcd based such as Kubernetes
in order to improve performance, scale and ease operation.

We are planning to add following sub components.

- API Server (Python-based VNC API Server equivalent)
- Sync (IF-MAP, RabbitMQ related code equivalent but depends on etcd)
- Agent (SchemaTransformer, Device Manager equivalent)
- Code generation tool (generateDS equivalent)

Currently, this project is
POC stage so any external/internal API or design subject to change up
to community discussion.

## Development setup

### Step1. Install Go, Docker and Docker Compose

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Step2. Go get Contrail

```bash
go get -u github.com/Juniper/contrail
```

Note that `go get -u github.com/Juniper/contrail/cmd/contrailutil` fails because we don't
commit generated code.

### Step3. Install dependencies

```bash
# move to repo
cd $HOME/go/src/github.com/Juniper/contrail
# make sure put GOBIN to path
export PATH=$PATH:$HOME/go/bin
# or if you defined GOPATH
# cd $GOPATH/src/github.com/Juniper/contrail
make deps
```

### Step4. Generate source code

```bash
make generate
```

### Step5. Install Contrail binaries

```bash
make install
```

### Step6. Setup test environment

```bash
make testenv reset_db
```

Note that these commands use `docker` command and depending on your Docker configuration they may require root permissions.
See: [Docker Documentation](https://docs.docker.com/install/linux/linux-postinstall/#manage-docker-as-a-non-root-user)

## First run

- Run Contrail process

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

- Run CLI commands

    ```bash
    export CONTRAIL_CONFIG=sample/cli.yml

    # Show schema
    contrailcli schema virtual_network

    # Sync resources
    contrailcli sync sample/sample_resource.yml

    # List resources
    contrailcli list virtual_network --detail

    # Delete resources
    contrailcli delete sample/sample_resource.yml
    ```

    See: [CLI Usage](doc/cli.md)

## Testing

Run all tests with coverage:

```bash
make test
```

## How to contribute

- Follow [Openstack review process](https://docs.openstack.org/infra/manual/developers.html)
- Use [Tungsten Fabric Gerrit](https://review.opencontrail.org)
- Comply to [Code review guidelines](REVIEW.md)

### Step1

Setup Gerrit account. Sign CLA.

### Step2

Install [git-review tool](https://docs.openstack.org/infra/git-review/installation.html).

```bash
pip install git-review
```

### Step3

Send review to Gerrit:

```bash
git review
```

## Documentation

See: [Documentation index](./doc/index.md)

# Dummy change

- 1
