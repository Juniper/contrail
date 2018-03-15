package etcdlock

import (
	"context"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3lock/v3lockpb"
	"github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Exposed errors.
var (
	ErrEmptyKey = errors.New("empty key")
)

const (
	defaultEtcdKeyPrefix = "__etcd_lock/"
	retryCount           = 3
)

// Locker is the client for acquiring distributed locks from etcd. It should be
// created from NewLocker() function.
type Locker struct {
	etcdKeyPrefix string
	leaseCli      etcdserverpb.LeaseClient
	kvCli         etcdserverpb.KVClient
	lockCli       v3lockpb.LockClient
}

// LockerOptions is the options for NewLocker() function.
type LockerOptions struct {
	// The address of etcd(v3) server.
	Address string
	// Options used for `grpc.Dial`.
	DialOptions []grpc.DialOption
	// Prefix of the keys of locks in etcd, by default is "__etcd_lock/"
	EtcdKeyPrefix string
}

// NewLocker creates a Locker according to the given options.
func NewLocker(options LockerOptions) (*Locker, error) {
	conn, err := grpc.Dial(options.Address, options.DialOptions...)
	if err != nil {
		return nil, err
	}

	if options.EtcdKeyPrefix == "" {
		options.EtcdKeyPrefix = defaultEtcdKeyPrefix
	}

	locker := &Locker{
		etcdKeyPrefix: options.EtcdKeyPrefix,
		leaseCli:      etcdserverpb.NewLeaseClient(conn),
		kvCli:         etcdserverpb.NewKVClient(conn),
		lockCli:       v3lockpb.NewLockClient(conn),
	}

	return locker, nil
}

// Lock acquires a distributed lock for the specified resource
// from etcd v3.
func (l *Locker) Lock(ctx context.Context, keyName string, timeout time.Duration) (*Lock, error) {
	if keyName == "" {
		return nil, errors.WithStack(ErrEmptyKey)
	}

	var try int
	for {
		try++
		leaseID, err := l.getLease(ctx, timeout)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		lockRes, err := l.lockCli.Lock(ctx, &v3lockpb.LockRequest{
			Name:  l.assembleKeyName(keyName),
			Lease: leaseID,
		})

		if err != nil {
			// Retry when the etcd server is too busy to handle transactions.
			if try <= retryCount && strings.Contains(err.Error(), "too many requests") {
				time.Sleep(time.Millisecond * time.Duration(500) * time.Duration(try))
				continue
			}
			return nil, errors.WithStack(err)
		}

		return &Lock{locker: l, keyName: lockRes.Key}, nil
	}
}

// IsLocked checks whether the specified resource has already been locked.
func (l *Locker) IsLocked(ctx context.Context, keyName string) (bool, error) {
	if keyName == "" {
		return false, errors.WithStack(ErrEmptyKey)
	}

	key := l.assembleKeyName(keyName)
	end := []byte(clientv3.GetPrefixRangeEnd(string(key)))

	rangeRes, err := l.kvCli.Range(ctx, &etcdserverpb.RangeRequest{
		Key:       key,
		RangeEnd:  end,
		CountOnly: true,
	})

	if err != nil {
		return false, errors.WithStack(err)
	}

	return rangeRes.Count != 0, nil
}

func (l *Locker) unlock(ctx context.Context, keyName []byte) error {
	_, err := l.lockCli.Unlock(ctx, &v3lockpb.UnlockRequest{Key: keyName})

	return errors.WithStack(err)
}

func (l *Locker) getLease(ctx context.Context, timeout time.Duration) (int64, error) {
	leaseRes, err := l.leaseCli.LeaseGrant(ctx, &etcdserverpb.LeaseGrantRequest{
		TTL: int64(timeout.Seconds()),
	})

	if err != nil {
		return 0, errors.WithStack(err)
	}

	return leaseRes.ID, nil
}

func (l *Locker) assembleKeyName(keyName string) []byte {
	return []byte(l.etcdKeyPrefix + keyName)
}
