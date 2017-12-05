# Go code base for contrail-common


# Important principal

- Apply lint tools
- Go get must simply work
- Follow best practices
  Effective Go: https://golang.org/doc/effective_go.html

# How to build

``` shell
go get github.com/contrail-common/go
```

# Generate Code

``` shell
export SCHEMA="path to json schema"
make generate
```

# Schema Files

Note that schema stored here is just a cache for helping development. 
Developers should make sure download latest schema from http://github.com/Juniper/contrail-api-client

# Packaging

TBD

# Dependency management

We use glide for dependency management for go code.

https://github.com/Masterminds/glide

see