# Go code base for contrail projects

# Important principal

- Apply lint tools
- Go get must simply work
- Follow best practices
  Effective Go: https://golang.org/doc/effective_go.html

# How to build

``` shell
go get github.com/Juniper/contrail
```

# Generate Code

``` shell
make generate
```

Templates are stored in tools/templates
You can add your template on template_config.yaml

# Schema Files

Note that schema stored here is just a cache for helping development.
Developers should make sure download latest schema from http://github.com/Juniper/contrail-api-client

JSON version stored in public/schema.json

# Testing

You need to run local mysql running with test configuraion.

ID: root
Password: contrail123
DataBase: contrail_test

Init DB before test
```
./tool/reset_db.sh
```

```
make test
```

# API Server

You can run API server using this command.

```
go run cmd/contrail/main.go server -c tools/test_config.yml
```

# Binary

Dep, RPM and Binaries are stored in release page.

(see https://github.com/Juniper/contrail/releases)

# commands

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
  -c, --config string   Configuraion File
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
  package     make a dep and rpm package

Flags:
  -h, --help   help for contrailutil
```

# Packaging

(1) Install FPM (https://github.com/jordansissel/fpm)

```
gem install --no-ri --no-rdoc fpm
```

(2) make package

```
make package
```

# Dependency management

We use golang standard dep tool for dependency management.
(see https://github.com/golang/dep)

```
brew install dep
```
