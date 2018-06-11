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

	"github.com/Juniper/contrail/pkg/log"
	"github.com/coreos/etcd/clientv3"
	conc "github.com/coreos/etcd/clientv3/concurrency"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	kvClientRequestTimeout = 60 * time.Second
)

// WatchCallback is callback for Watch functions.
type WatchCallback func(ctx context.Context, index int64, oper int32, key string, newValue []byte)

// Client is an etcd client using clientv3.
type Client struct {
	Etcd *clientv3.Client
	log  *logrus.Entry
}

// DialByConfig connects to the etcd db based on viper configuration.
func DialByConfig() (*clientv3.Client, error) {
	cfg := clientv3.Config{
		Endpoints:   viper.GetStringSlice("etcd.endpoints"),
		Username:    viper.GetString("etcd.username"),
		Password:    viper.GetString("etcd.password"),
		DialTimeout: viper.GetDuration("etcd.dial_timeout"),
	}

	etcd, err := clientv3.New(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "Error connecting to ETCD: %s\n", cfg.Endpoints)
	}
	return etcd, nil
}

func NewClient(c *clientv3.Client) *Client {
	return &Client{Etcd: c, log: log.NewLogger("etcd-client")}
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

// Put puts value in etcd no matter if it was there or not.
func (c *Client) Put(ctx context.Context, key string, value []byte) error {
	kvHandle := clientv3.NewKV(c.Etcd)

	_, err := kvHandle.Put(ctx, key, string(value))

	return err
}

// Create puts value in etcd if following key didn't exist.
func (c *Client) Create(ctx context.Context, key string, value []byte) error {
	kvHandle := clientv3.NewKV(c.Etcd)

	_, err := kvHandle.Txn(ctx).
		If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Then(clientv3.OpPut(key, string(value))).
		Commit()

	return err
}

// Update puts value in etcd if key existed before.
func (c *Client) Update(ctx context.Context, key, value string) error {
	kvHandle := clientv3.NewKV(c.Etcd)

	_, err := kvHandle.Txn(ctx).
		If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Else(clientv3.OpPut(key, value)).
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

type Event struct {
	Revision int64
	Type     int32
	Key      string
	Value    []byte
}

// WatchRecursive Watches a key pattern for changes After an Index
func (c *Client) WatchRecursive(
	ctx context.Context, keyPattern string, afterIndex int64,
) chan Event {
	resultChan := make(chan Event)
	rchan := c.Etcd.Watch(ctx, keyPattern,
		clientv3.WithPrefix(), clientv3.WithRev(afterIndex))

	go func() {
		for wresp := range rchan {
			for _, ev := range wresp.Events {
				resultChan <- Event{
					Revision: wresp.Header.Revision,
					Type:     int32(ev.Type),
					Key:      string(ev.Kv.Key),
					Value:    ev.Kv.Value,
				}
			}
		}
	}()

	return resultChan
}

// WatchRecursive Recursively Watches a key pattern for changes
//func (c *Client) WatchRecursive(ctx context.Context,
//keyPattern string) {
//afterIndex := int64(0)
//for {
//select {
//case <-ctx.Done():
//return
//default:
//c.WatchAfterIndex(ctx, afterIndex, keyPattern)
//}
//}
//}

//// WatchRecursiveAfterIndex Recursively Watches a key pattern for changes
//// After an Index
//func (c *Client) WatchRecursiveAfterIndex(ctx context.Context,
//afterIndex int64, keyPattern string) {
//for {
//select {
//case <-ctx.Done():
//return
//default:
//c.WatchAfterIndex(ctx, afterIndex, keyPattern)
//}
//}
//}

// InTransaction wraps clientv3 transaction and wraps conc.STM with own Txn.
func (c *Client) InTransaction(ctx context.Context, do func(context.Context) error) error {
	if txn := GetTxn(ctx); txn != nil {
		// Transaction already in context
		return do(ctx)
	}
	// New transaction required

	ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
	defer cancel()

	_, err := conc.NewSTM(c.Etcd, func(stm conc.STM) error {
		return do(WithTxn(ctx, stmTxn{stm, c.log}))
	}, conc.WithAbortContext(ctx))
	return err
}
