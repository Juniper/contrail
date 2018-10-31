# Command `generate` in `contrailschema`

Tool `contrailschema` provides `generate` command that is able to generate source code for REST API, database access, resource management etc. based on schema stored in *.yaml* files.

## Usage

```bash
contrailschema generate [Flags]

Flags:
  -h, --help                    help for generate
      --openapi-output string   OpenAPI Output path
      --output-dir string       output dir (default "./")
  -p, --package-path string     Package name (default "github.com/Juniper/contrail")
      --proto-package string    Protoc package base (default "github.com.Juniper.contrail")
      --schema-output string    Schema Output path
  -s, --schemas string          Schema Directory
  -t, --templates string        Template Configuration
```

## Example usage
Reference example for contrailschema generate usage is in [`Makefile`](../Makefile)

```bash
go run cmd/contrailschema/main.go generate --schemas schemas --templates tools/templates/template_config.yaml --schema-output public/schema.json --openapi-output public/openapi.json
```

## Schema Directory
`generate` command reads schemas from **Schema Directory** recursively. All the files with *.yaml* extension are read and used to generate code.

There is also special directory name *overrides* that is not read in a same way as other directories. This special directory  is used to read information that will override any information from other files. It is required to inject correct data into schema because schema files are also generated elsewhere.
