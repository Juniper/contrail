# Convert command

You can import/export configuraion data using contrailutil convert command.
This command can be used for backup, database migraion and testing.

Here is a supported input type list.

- Cassandra (old version of VNC API Server)
- Cassandra JSON Dump (json dump file made by control node utility)
- File (yaml file)

Here is a supported output type list.

- RDBMS
- File

``` shell
contrailutil convert --help
This command converts data formats from one to another

Usage:
  contrailutil convert [flags]

Flags:
  -h, --help             help for convert
  -i, --in string        Input file or Cassandra host
      --intype string    input type: cassandra,cassandra_dump and yaml are supported
  -o, --out string       Output file
      --outtype string   output type: rdbms and yaml are supported

Global Flags:
  -c, --config string   Configuration File
```

# Example usage

File -> RDBMS ( iniliaize database entiry by file for restore or testing )

``` shell
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail.yml
```

Cassandra -> File (database migraion from cassandra to new DB)

``` shell
contrailutil convert --intype cassandra --in localhost --outtype yaml --out dbdata.yaml
```