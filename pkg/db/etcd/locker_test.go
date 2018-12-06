package etcd

import (
	"context"
	"testing"
	"time"

	etcdlock "github.com/DavidCai1993/etcd-lock"
	"github.com/stretchr/testify/assert"
)

func TestDistributedLocker_DoWithLock(t *testing.T) {
	dl := &DistributedLocker{}

	err := dl.DoWithLock(
		withLock(context.Background(), &etcdlock.Lock{}),
		"",
		time.Second,
		func(ctx context.Context) error {
			if l := getLock(ctx); l == nil {
				return assert.AnError
			}
			return nil
		},
	)

	assert.NoError(t, err)
}
