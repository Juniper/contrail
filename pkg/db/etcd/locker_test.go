package etcd

import (
	"context"
	"testing"
	"time"

	etcdlock "github.com/DavidCai1993/etcd-lock"
	"github.com/stretchr/testify/assert"
)

func TestDistributedLocker_DoWithLock(t *testing.T) {
	l := &DistributedLocker{}
	lock := &etcdlock.Lock{}
	ctx := withLock(context.Background(), lock)

	err := l.DoWithLock(ctx, "", time.Second, func(ctx context.Context) error {
		if t := getLock(ctx); t == nil {
			return assert.AnError
		}
		return nil
	})

	assert.NoError(t, err)
}
