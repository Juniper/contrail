package etcd

import (
	"context"
	"time"

	"github.com/DavidCai1993/etcd-lock"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type DistributedLocker struct {
	locker *etcdlock.Locker
}

// NewDistributedLocker creates locker connected to the first etcd node from viper configuration.
func NewDistributedLocker() (*DistributedLocker, error) {
	endpoints := viper.GetStringSlice("etcd.endpoints")
	if len(endpoints) < 1 {
		return nil, errors.New("no etcd endpoints in config")
	}

	l, err := etcdlock.NewLocker(etcdlock.LockerOptions{
		Address: endpoints[0],
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Error connecting to ETCD: %s\n", endpoints[0])
	}
	return &DistributedLocker{locker: l}, nil
}

// DoWithLock executes provided callback inside lock secured section.
func (l *DistributedLocker) DoWithLock(
	ctx context.Context,
	key string,
	lockTTL time.Duration,
	do func(ctx context.Context) error,
) (err error) {
	if lock := getLock(ctx); lock != nil {
		return do(ctx)
	}

	lock, err := l.locker.Lock(ctx, key, lockTTL)
	if err != nil {
		return errors.Wrap(err, "cannot acquire lock")
	}

	defer func() {
		if unlockErr := lock.Unlock(ctx); unlockErr != nil {
			err = errors.Wrap(unlockErr, "cannot release lock")
		}
	}()

	return do(withLock(ctx, lock))
}

var lockKey interface{} = "etcd-lock"

func getLock(ctx context.Context) *etcdlock.Lock {
	iLock := ctx.Value(lockKey)
	l, _ := iLock.(*etcdlock.Lock)
	return l
}

func withLock(ctx context.Context, l *etcdlock.Lock) context.Context {
	return context.WithValue(ctx, lockKey, l)
}
