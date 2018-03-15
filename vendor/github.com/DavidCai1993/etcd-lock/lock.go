package etcdlock

import (
	"context"
)

// Lock is a distributed lock of specified resource which was acquired from
// etcd v3.
type Lock struct {
	locker  *Locker
	keyName []byte
}

// Unlock unlocks this lock.
func (l *Lock) Unlock(ctx context.Context) error {
	return l.locker.unlock(ctx, l.keyName)
}
