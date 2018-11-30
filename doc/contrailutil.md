# Contrail Utils

Utility CLI for development and debugging.

```bash
Usage:
  contrailutil [flags]
  contrailutil [command]

Available Commands:
  convert     convert data format
  help        Help about any command
  record_test Record test result 

Flags:
  -c, --config string   Configuration File
  -h, --help            help for contrailutil

Use "contrailutil [command] --help" for more information about a command.
```

## Convert command

Convert command is used for import/export configuration data.
This command can be used for backup, database migration or testing.

```bash
Usage:
  contrailutil convert [flags]

Flags:
  -p, --cassandra_port int      Cassandra port (default 9042)
  -t, --cassandra_timeout int   Cassandra timeout in seconds (default 3600)
  -h, --help                    help for convert
  -i, --in string               Input file or Cassandra host
      --intype string           input type: "cassandra", "cassandra_dump", "yaml" and "rdbms" are supported
  -o, --out string              Output file
      --outtype string          output type: "rdbms", "yaml" and "etcd" are supported

Global Flags:
  -c, --config string   Configuration File
```

### Supported input

- Cassandra (old version of VNC API Server)
- Cassandra JSON Dump (json dump file made by control node utility)
- RDBMS
- File (yaml file)

### Supported output

- RDBMS
- File
- etcd

### Convert example usage

Cassandra -> File (database migration from Cassandra to new DB)

```bash
$ contrailutil convert --intype cassandra --in localhost --outtype yaml --out dbdata.yaml
$ contrailutil convert --intype cassandra --in localhost -p 9041 -t 60 --outtype yaml --out dbdata.yaml
```

File -> etcd (restore to etcd from yaml or json)

```bash
$ contrailutil convert --intype yaml --in tools/init_data.yaml --outtype etcd -c sample/contrail.yml
$ contrailutil convert --intype yaml --in tools/init_data.json --outtype etcd -c sample/contrail.yml
```

RDBMS -> File ( backup in yaml or json )

```bash
$ contrailutil convert --intype rdbms --outtype yaml --out db.yml -c sample/contrail.yml
$ contrailutil convert --intype rdbms --outtype yaml --out db.json -c sample/contrail.yml
```

File -> RDBMS ( initialize database entity by file for restore or testing )

```bash
$ contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail.yml
```

## Test recorder command

Test recorder gives you ability to run test scenario from YAML and save result to file.
Test scenario should cleanup after itself because `contrailutil` does not track changes in DB.
This may cause unique constraint violation errors on second run.

```bash
Usage:
  contrailutil record_test [flags]

Flags:
  -a, --auth_url string   AuthURL
  -e, --endpoint string   Endpoint
  -h, --help              help for record_test
  -i, --input string      Input test scenario path
  -o, --output string     Output test scenario path
  -v, --vars string       test variables

Global Flags:
  -c, --config string   Configuration File
```

### Test recorder example usage

**Before running test record make sure your API server is up and running.**

Run in development environment:

```bash
$ contrail run -c sample/contrail.yml
```

Run and record test scenario:

```bash
$ contrailutil record_test -i scenario.yml -o scenario_record.yml
```

### Example test scenario

```yaml
test_data:
  tag: &tag
    uuid: a0c60f49-3e93-4b87-a8f5-5b5e1fefd082
    fq_name:
    - namespace=k8s-default
    tag_value: k8s-default
    tag_type_name: namespace

clients:
  default:
    id: alice
    password: alice_password
    domain: default
    insecure: true
    scope:
      project:
        name: admin

workflow:
- name: create tag
  request:
    path: http://127.0.0.1:9091/tags
    method: POST
    expected:
    - 200
    data:
      tag: *tag
  expect:
    tag:
      <<: *tag
      display_name: namespace=k8s-default
      tag_id: 0x00ff0002

- name: delete tag
  request:
    path: http://127.0.0.1:9091/tag/a0c60f49-3e93-4b87-a8f5-5b5e1fefd082
    method: DELETE
    expected:
    - 200
  expect: null
```
