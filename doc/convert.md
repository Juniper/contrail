# Convert command

You can import/export configuration data using contrailutil convert command.
This command can be used for backup, database migration and testing.

Here is a supported input type list.

- Cassandra (old version of VNC API Server)
- Cassandra JSON Dump (json dump file made by control node utility)
- RDBMS
- File (yaml file)

Here is a supported output type list.

- RDBMS
- File
- etcd

``` shell
contrailutil convert --help
This command converts data formats from one to another

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

# Example usage

Cassandra -> File (database migration from Cassandra to new DB)

``` shell
contrailutil convert --intype cassandra --in localhost --outtype yaml --out dbdata.yaml
contrailutil convert --intype cassandra --in localhost -p 9041 -t 60 --outtype yaml --out dbdata.yaml
```

File -> etcd (restore to etcd from yaml or json)

``` shell
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype etcd -c sample/contrail.yml
contrailutil convert --intype yaml --in tools/init_data.json --outtype etcd -c sample/contrail.yml
```

RDBMS -> File ( backup in yaml or json )

``` shell
contrailutil convert --intype rdbms --outtype yaml --out db.yml -c sample/contrail.yml
contrailutil convert --intype rdbms --outtype yaml --out db.json -c sample/contrail.yml
```

File -> RDBMS ( initialize database entity by file for restore or testing )

``` shell
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail.yml
```