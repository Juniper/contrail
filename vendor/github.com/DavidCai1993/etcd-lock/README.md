# etcd-lock
[![Build Status](https://travis-ci.org/DavidCai1993/etcd-lock.svg?branch=master)](https://travis-ci.org/DavidCai1993/etcd-lock)
[![Coverage Status](https://coveralls.io/repos/github/DavidCai1993/etcd-lock/badge.svg?branch=master)](https://coveralls.io/github/DavidCai1993/etcd-lock?branch=master)

Distributed locks powered by [etcd v3](https://github.com/coreos/etcd) for Go.

## Installation

```
go get -u github.com/DavidCai1993/etcd-lock
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/DavidCai1993/etcd-lock

## Usage

```go
import (
  "github.com/DavidCai1993/etcd-lock"
)
```

```go
locker, err := etcdlock.NewLocker(etcdlock.LockerOptions{
  Address:        "127.0.0.1:2379",
  DialOptions:    []grpc.DialOption{grpc.WithInsecure()},
})

if err != nil {
  log.Fatalln(err)
}

// Acquire a lock for a specified recource.
if _, err = locker.Lock(context.Background(), "resource_key", 5*time.Second); err != nil {
  log.Fatalln(err)
}

// This lock will be acquired after 5s, and before that current goroutine
// will be blocked.
anotherLock, err := locker.Lock(context.Background(), "resource_key", 5*time.Second)
if err != nil {
  log.Fatalln(err)
}

// Unlock the lock manually.
if err := anotherLock.Unlock(context.Background()); err != nil {
  log.Fatalln(err)
}
```
