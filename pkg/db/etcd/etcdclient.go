/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * ETCD Client interface
 *
 */

package etcdclient

import (
	"context"
	"time"

	"github.com/DavidCai1993/etcd-lock"
	"github.com/pkg/errors"

	client "github.com/coreos/etcd/clientv3"
	log "github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

// Callback from Watch functions
type Callback func(ctx context.Context, cl *IntentEtcdClient, index int64,
	oper int32, key, newValue string)

// EtcdClient represent client interface
type EtcdClient interface {
	// Get gets a key's value from Etcd
	Get(ctx context.Context, key string) (string, error)

	// Set sets a key,value in Etcd
	Set(ctx context.Context, key, value string) error

	// Update modifies a key's value in Etcd
	Update(ctx context.Context, key, value string) error

	// Delete deletes a value in Etcd
	Delete(ctx context.Context, key, value string) error

	// Watches a key pattern for changes after an index
	WatchAfterIndex(ctx context.Context, afterIndex int64,
		keyPattern string, callback Callback) error

	// Recursively Watches a key pattern for changes
	WatchRecursive(ctx context.Context, keyPattern string, callback Callback)

	// Recursively Watches a key pattern for changes after an index
	WatchRecursiveAfterIndex(ctx context.Context, afterIndex int64,
		keyPattern string, callback Callback)

	// CreateLock creates a lock for a key
	CreateLock(server string) (*etcdlock.Locker, error)

	// AcquireLock acquires a lock
	AcquireLock(ctx context.Context, locker *etcdlock.Locker, key string,
		ttl int) (*etcdlock.Lock, error)

	// ReleaseLock releases the acquired lock
	ReleaseLock(ctx context.Context) error
}

// IntentEtcdClient implements EtcdClient
type IntentEtcdClient struct {
	Etcd *client.Client
}

// Dial constructs a new EtcdClient
func Dial(etcdURI []string) (*IntentEtcdClient, error) {
	cfg := client.Config{
		Endpoints: etcdURI,
		//Transport: DefaultTransport,
	}

	etcd, err := client.New(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "Error connecting to ETCD: %s\n", etcdURI)
	}
	log.Info("Connected to the ETCD Server")
	return &IntentEtcdClient{etcd}, nil
}

// Get gets a value in Etcd
func (etcdClient *IntentEtcdClient) Get(ctx context.Context, key string) (string, error) {
	kvHandle := client.NewKV(etcdClient.Etcd)
	response, err := kvHandle.Get(ctx, key)
	if err != nil || response.Count == 0 {
		return "", err
	}
	return string(response.Kvs[0].Value[:]), nil
}

// Set sets a value in Etcd
func (etcdClient *IntentEtcdClient) Set(ctx context.Context, key, value string) error {
	kvHandle := client.NewKV(etcdClient.Etcd)

	_, err := kvHandle.Txn(ctx).
		If(client.Compare(client.Version(key), "=", 0)).
		Then(client.OpPut(key, value)).
		Commit()

	return err
}

// Update updates a key with a ttl value
func (etcdClient *IntentEtcdClient) Update(ctx context.Context, key, value string) error {
	kvHandle := client.NewKV(etcdClient.Etcd)

	_, err := kvHandle.Txn(ctx).
		If(client.Compare(client.Version(key), "=", 0)).
		Else(client.OpPut(key, value)).
		Commit()

	return err
}

// Delete deletes a key/value in Etcd
func (etcdClient *IntentEtcdClient) Delete(ctx context.Context, key string) error {
	kvHandle := client.NewKV(etcdClient.Etcd)

	_, err := kvHandle.Txn(ctx).
		If(client.Compare(client.Version(key), "=", 0)).
		Else(client.OpDelete(key)).
		Commit()

	return err
}

// WatchAfterIndex Watches a key pattern for changes After an Index
func (etcdClient *IntentEtcdClient) WatchAfterIndex(ctx context.Context,
	afterIndex int64, keyPattern string, callback Callback) {

	rchan := etcdClient.Etcd.Watch(ctx, keyPattern,
		client.WithPrefix(), client.WithRev(afterIndex))
	for wresp := range rchan {
		for _, ev := range wresp.Events {
			afterIndex = wresp.Header.Revision
			if callback == nil {
				continue
			}
			callback(ctx, etcdClient, wresp.Header.Revision, int32(ev.Type),
				string(ev.Kv.Key[:]), string(ev.Kv.Value[:]))
		}
	}
}

// WatchRecursive Recursively Watches a key pattern for changes
func (etcdClient *IntentEtcdClient) WatchRecursive(ctx context.Context,
	keyPattern string, callback Callback) {
	afterIndex := int64(0)
	for {
		etcdClient.WatchAfterIndex(ctx, afterIndex, keyPattern, callback)
	}
}

// WatchRecursiveAfterIndex Recursively Watches a key pattern for changes
//  After an Index
func (etcdClient *IntentEtcdClient) WatchRecursiveAfterIndex(ctx context.Context,
	afterIndex int64, keyPattern string, callback Callback) {
	for {
		etcdClient.WatchAfterIndex(ctx, afterIndex, keyPattern, callback)
	}
}

// CreateLock creates a lock for a key
func (etcdClient *IntentEtcdClient) CreateLock(server string) (*etcdlock.Locker, error) {
	return etcdlock.NewLocker(etcdlock.LockerOptions{
		Address:     server,
		DialOptions: []grpc.DialOption{grpc.WithInsecure()},
	})
}

// AcquireLock acquires a lock on a key
// The Context will define if the Lock has a timeout or Blocking
func (etcdClient *IntentEtcdClient) AcquireLock(ctx context.Context,
	locker *etcdlock.Locker, key string, ttl int) (*etcdlock.Lock, error) {
	lock, err := locker.Lock(ctx, key,
		time.Duration(ttl)*time.Second)
	if err != nil {
		log.Error("Cannot acquire Lock", err)
		return nil, err
	}
	return lock, nil
}

// ReleaseLock releases the acquired lock
func (etcdClient *IntentEtcdClient) ReleaseLock(ctx context.Context,
	locker *etcdlock.Lock) error {
	if err := locker.Unlock(ctx); err != nil {
		log.Error("Cannot Release lock", err)
		return err
	}
	return nil
}
