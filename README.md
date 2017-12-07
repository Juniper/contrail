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

# Schema Files

Note that schema stored here is just a cache for helping development. 
Developers should make sure download latest schema from http://github.com/Juniper/contrail-api-client

# Testing

You need to run local mysql running with test configuraion.

ID: root 
Password: contrail123 
DataBase: contrail_test 

```
make test
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