/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * ETCD Client interface
 *
 */

// TODO(Michal): Change file name to client since etcd/etcdclient.go is stuttered.
// TODO(Michal): Add some logging to client methods.

package etcd

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/log"
	"github.com/coreos/etcd/clientv3"
	conc "github.com/coreos/etcd/clientv3/concurrency"
	"github.com/sirupsen/logrus"
)

const (
	kvClientRequestTimeout = 60 * time.Second
)

// Callback from Watch functions
type Callback func(ctx context.Context, index int64, oper int32, key string, newValue []byte)

// Txn is a transaction object allowing to perform operations in it.
type Txn interface {
	Get(key string) []byte
	Put(key string, val []byte)
}

// Client is an etcd client using clientv3.
type Client struct {
	Etcd *clientv3.Client
	log  *logrus.Entry
}

// Dial constructs a new EtcdClient
func Dial() (*Client, error) {
	cfg := clientv3.Config{
		Endpoints:   viper.GetStringSlice("etcd.endpoints"),
		Username:    viper.GetString("etcd.username"),
		Password:    viper.GetString("etcd.password"),
		DialTimeout: viper.GetDuration("etcd.dial_timeout"),
	}

	l := log.NewLogger("etcd-client")

	etcd, err := clientv3.New(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "Error connecting to ETCD: %s\n", cfg.Endpoints)
	}
	l.Info("Connected to the ETCD Server")
	return &Client{Etcd: etcd, log: l}, nil
}

// Get gets a value in Etcd
func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	kvHandle := clientv3.NewKV(c.Etcd)
	response, err := kvHandle.Get(ctx, key)
	if err != nil || response.Count == 0 {
		return nil, err
	}
	return response.Kvs[0].Value, nil
}

// Put sets a value in Etcd
func (c *Client) Put(ctx context.Context, key string, value []byte) error {
	kvHandle := clientv3.NewKV(c.Etcd)

	_, err := kvHandle.Txn(ctx).
		If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Then(clientv3.OpPut(key, string(value))).
		Commit()

	return err
}

// Delete deletes a key/value in Etcd
func (c *Client) Delete(ctx context.Context, key string) error {
	kvHandle := clientv3.NewKV(c.Etcd)

	_, err := kvHandle.Txn(ctx).
		If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Else(clientv3.OpDelete(key)).
		Commit()

	return err
}

// WatchAfterIndex Watches a key pattern for changes After an Index
func (c *Client) WatchAfterIndex(ctx context.Context,
	afterIndex int64, keyPattern string, callback Callback) {

	rchan := c.Etcd.Watch(ctx, keyPattern,
		clientv3.WithPrefix(), clientv3.WithRev(afterIndex))
	for wresp := range rchan {
		for _, ev := range wresp.Events {
			afterIndex = wresp.Header.Revision
			if callback == nil {
				continue
			}
			callback(ctx, wresp.Header.Revision, int32(ev.Type),
				string(ev.Kv.Key), ev.Kv.Value)
		}
	}
}

// WatchRecursive Recursively Watches a key pattern for changes
func (c *Client) WatchRecursive(ctx context.Context,
	keyPattern string, callback Callback) {
	afterIndex := int64(0)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			c.WatchAfterIndex(ctx, afterIndex, keyPattern, callback)
		}
	}
}

// WatchRecursiveAfterIndex Recursively Watches a key pattern for changes
//  After an Index
func (c *Client) WatchRecursiveAfterIndex(ctx context.Context,
	afterIndex int64, keyPattern string, callback Callback) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			c.WatchAfterIndex(ctx, afterIndex, keyPattern, callback)
		}
	}
}

// InTransaction wraps clientv3 transaction and wraps conc.STM with own sink.Txn.
func (c *Client) InTransaction(ctx context.Context, do func(Txn) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
	defer cancel()

	_, err := conc.NewSTM(c.Etcd, func(stm conc.STM) error {
		return do(stmTxn{stm, c.log})
	}, conc.WithAbortContext(ctx))
	return err
}

type stmTxn struct {
	conc.STM
	log *logrus.Entry
}

func (s stmTxn) Get(key string) []byte {
	s.log.WithFields(logrus.Fields{"key": key}).Debugf(
		"Get resource from etcd in transaction")
	return []byte(s.STM.Get(key))
}

func (s stmTxn) Put(key string, val []byte) {
	s.log.WithFields(logrus.Fields{"key": key}).Debugf(
		"Put resource in etcd in transaction")
	s.STM.Put(key, string(val))
}
