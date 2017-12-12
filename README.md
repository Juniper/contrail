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

# Packaging

TBD

# Dependency management

We use golang standard dep tool for dependency management.
(see https://github.com/golang/dep)
brew install dep

We use glide for dependency management for go code.

https://github.com/Masterminds/glide

see