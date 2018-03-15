package etcdlock_test

import (
	"context"
	"log"
	"time"

	"github.com/DavidCai1993/etcd-lock"
	"google.golang.org/grpc"
)

func Example() {
	locker, err := etcdlock.NewLocker(etcdlock.LockerOptions{
		Address:     "127.0.0.1:2379",
		DialOptions: []grpc.DialOption{grpc.WithInsecure()},
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
	anotherLock, err := locker.Lock(context.Background(), "resource_key", 3*time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	// Unlock the lock manually.
	if err := anotherLock.Unlock(context.Background()); err != nil {
		log.Fatalln(err)
	}
}

func ExampleLocker_Lock() {
	locker, err := etcdlock.NewLocker(etcdlock.LockerOptions{
		Address:     "127.0.0.1:2379",
		DialOptions: []grpc.DialOption{grpc.WithInsecure()},
	})

	if err != nil {
		log.Fatalln(err)
	}

	// This lock will be expired in 3 seconds.
	if _, err := locker.Lock(context.Background(), "resource_key", 3*time.Second); err != nil {
		log.Fatalln(err)
	}
}
