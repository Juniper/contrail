# Convert command

You can import/export configuraion data using contrailutil convert command.

Here is a supported source list.

- Cassandra (old version of VNC API Server)
- Cassandra JSON Dump (json dump file made by control node utility)
- File (yaml file)

Here is a supported export list.

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

File -> RDBMS

``` shell
contrailutil convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail.yml
```

Cassandra -> File

``` shell
contrailutil convert --intype cassandra --in localhost --outtype yaml --out dbdata.yaml 
```