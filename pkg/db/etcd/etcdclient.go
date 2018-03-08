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
	"fmt"

	client "github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Callback from WatchRecursive functions
type Callback func(cl *IntentEtcdClient, index int64, oper int32, key, newValue string)

// EtcdClient represent client interface
type EtcdClient interface {
	// Get gets a key's value from Etcd
	Get(key string) (string, error)

	// Set sets a key,value in Etcd
	Set(key, value string) error

	// Update modifies a key's value in Etcd
	Update(key, value string) error

	// Delete deletes a value in Etcd
	Delete(key, value string) error

	// MkDir creates an empty directory in etcd
	MkDir(directory string) error

	// RmDir deletes an empty directory in etcd
	RmDir(directory string) error

	// Recursively Watches a Directory for changes
	WatchRecursive(directory string, callback Callback) error
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
func (etcdClient *IntentEtcdClient) Get(key string) (string, error) {
	kvHandle := client.NewKV(etcdClient.Etcd)
	response, err := kvHandle.Get(context.Background(), key)
	if err != nil || response.Count == 0 {
		return "", err
	}
	return string(response.Kvs[0].Value[:]), nil
}

// Set sets a value in Etcd
func (etcdClient *IntentEtcdClient) Set(key, value string) error {
	kvHandle := client.NewKV(etcdClient.Etcd)

	_, err := kvHandle.Txn(context.Background()).
		If(client.Compare(client.Version(key), "=", 0)).
		Then(client.OpPut(key, value)).
		Commit()

	return err
}

// Update updates a key with a ttl value
func (etcdClient *IntentEtcdClient) Update(key, value string) error {
	kvHandle := client.NewKV(etcdClient.Etcd)

	_, err := kvHandle.Txn(context.Background()).
		If(client.Compare(client.Version(key), "=", 0)).
		Else(client.OpPut(key, value)).
		Commit()

	return err
}

// Delete deletes a key/value in Etcd
func (etcdClient *IntentEtcdClient) Delete(key string) error {
	kvHandle := client.NewKV(etcdClient.Etcd)

	_, err := kvHandle.Txn(context.Background()).
		If(client.Compare(client.Version(key), "=", 0)).
		Else(client.OpDelete(key)).
		Commit()

	return err
}

// WatchRecursive Recursively Watches a Directory for changes After an Index
func (etcdClient *IntentEtcdClient) WatchRecursive(directory string,
	callback Callback) {

	afterIndex := int64(0)
	for {
		rchan := etcdClient.Etcd.Watch(context.Background(), directory,
			client.WithPrefix(), client.WithRev(afterIndex))
		for wresp := range rchan {
			for _, ev := range wresp.Events {
				fmt.Printf("%d %s %q : %q\n", wresp.Header.Revision, ev.Type, ev.Kv.Key, ev.Kv.Value)
				afterIndex = wresp.Header.Revision
				if callback == nil {
					continue
				}
				callback(etcdClient, wresp.Header.Revision, int32(ev.Type),
					string(ev.Kv.Key[:]), string(ev.Kv.Value[:]))
			}
		}
	}
}
