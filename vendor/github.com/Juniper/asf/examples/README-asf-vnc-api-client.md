# Introduction

This Document explains how to use asf and it schematool to generate
client library and its dependent models and services library.

Any schema driven application can use schematool and asf to generate
client library based on imput schema and use the client library in
the application.

In ./asf-vnc-api-client example application, the asf's schematool
is used to generate client library for all contrail schemas.

# How to generate client library

### 1. Create application directory

Create your application directory, In this example it is ./asf-vnc-api-client

### 2. Create template config file

Template config file dictates the schematool to use the following,

- template_path -> Which template file to use for code generation
- module -> Which go module contains above template path
- output_dir -> Directory in which the generated code is placed

In this example, template config is placed at ./asf-vnc-api-client/templates/template_config.yaml

### 3. Create tools.go

go module(https://blog.golang.org/using-go-modules) is widely used for managing
dependencies for a go based application/project, go module scans the import statements
in the application/project to figure out dependencies.

code generation tool (schematool) is available in asf go module, So asf is a
dependency for ./asf-vnc-api-client application.

To make go module treat asf schema tool as dependency, tools.go file is
created at ./asf-vnc-api-client/tools.go with proper import

Reference: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

### 4. Initialize go module

```shell
 cd ./asf-vnc-api-client; go mod init
```

Above command will initialize the go.mod file

### 5. Add dependecies to go.mod 

```shell
 cd ./asf-vnc-api-client; go mod tidy
```

Above command will scan the imports (in tools.go) and adds dependencies to go.mod file

### 6. generate code
```shell
 cd ./asf-vnc-api-client; make generate
```

Above command will generate following go pkgs,

- ./asf-vnc-api-client/pkg/client
- ./asf-vnc-api-client/pkg/models
- ./asf-vnc-api-client/pkg/services

Above client pkg can be imported and used in the application

# NOTE

- application should be built only after the generating the libraries
- ./asf-vnc-api-client/Makefile will give the user a better understanding.
