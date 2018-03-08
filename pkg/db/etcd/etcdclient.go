/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * ETCD Client interface
 *
 */

package etcdclient

import (
	"time"

	"github.com/coreos/etcd/client"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Callback from WatchRecursive functions
type Callback func(cl *IntentEtcdClient, index uint64, key, newValue string)

// EtcdClient represent client interface
type EtcdClient interface {
	// Get gets a key's value from Etcd
	Get(key string) (string, error)

	// Set sets a key,value in Etcd
	Set(key, value string) error

	// SetWithTTL sets a key,value in Etcd
	SetWithTTL(key, value string, ttl time.Duration) error

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
	etcd client.Client
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
	api := client.NewKeysAPI(etcdClient.etcd)
	response, err := api.Get(context.Background(), key, nil)
	if err != nil {
		if client.IsKeyNotFound(err) {
			// Ignore if key doesnt exist
			return "", nil
		}
		return "", err
	}
	return response.Node.Value, nil
}

// Set sets a value in Etcd
func (etcdClient *IntentEtcdClient) Set(key, value string) error {
	api := client.NewKeysAPI(etcdClient.etcd)
	_, err := api.Set(context.Background(), key, value, nil)
	return err
}

// SetWithTTL sets a value in Etcd with TTL
func (etcdClient *IntentEtcdClient) SetWithTTL(key, value string,
	ttl time.Duration) error {
	api := client.NewKeysAPI(etcdClient.etcd)
	opts := &client.SetOptions{TTL: ttl}
	_, err := api.Set(context.Background(), key, value, opts)
	return err
}

// Update updates a key with a ttl value
func (etcdClient *IntentEtcdClient) Update(key string) error {
	api := client.NewKeysAPI(etcdClient.etcd)
	refreshopts := &client.SetOptions{Refresh: true,
		PrevExist: client.PrevExist}
	_, err := api.Set(context.Background(), key, "", refreshopts)
	return err
}

// Delete deletes a key/value in Etcd
func (etcdClient *IntentEtcdClient) Delete(key string) error {
	api := client.NewKeysAPI(etcdClient.etcd)
	_, err := api.Delete(context.Background(), key, nil)
	if err != nil {
		if client.IsKeyNotFound(err) {
			// Ignore non-existent key
			return nil
		}
	}
	return err
}

// MkDir creates an empty directory in etcd
func (etcdClient *IntentEtcdClient) MkDir(directory string) error {
	api := client.NewKeysAPI(etcdClient.etcd)

	// Check if Directory exists
	res, err := api.Get(context.Background(), directory, nil)
	if err != nil && !client.IsKeyNotFound(err) {
		// Directory exists, ignore error
		return nil
	}

	if err != nil && client.IsKeyNotFound(err) {
		// Directory doesn't exist, create it
		opts := &client.SetOptions{Dir: true, PrevExist: client.PrevIgnore}
		_, err = api.Set(context.Background(), directory, "", opts)
		return err
	}

	// directory exists as a keyname
	if !res.Node.Dir {
		return errors.Wrapf(err, "Cannot overwrite key/value with a directory: %v",
			directory)
	}

	return nil
}

// RmDir deletes an empty directory in etcd
func (etcdClient *IntentEtcdClient) RmDir(directory string) error {
	api := client.NewKeysAPI(etcdClient.etcd)
	_, err := api.Get(context.Background(), directory, nil)
	if err != nil && !client.IsKeyNotFound(err) {
		// Error Occurred
		return err
	}
	if err != nil && client.IsKeyNotFound(err) {
		opts := &client.DeleteOptions{Dir: true}
		_, err = api.Delete(context.Background(), directory, opts)
		return err
	}

	return nil
}

// WatchRecursive Recursively Watches a Directory for changes After an Index
func (etcdClient *IntentEtcdClient) WatchRecursive(directory string,
	callback Callback) error {
	api := client.NewKeysAPI(etcdClient.etcd)
	afterIndex := uint64(0)

	for {
		watcher := api.Watcher(directory, &client.WatcherOptions{Recursive: true,
			AfterIndex: afterIndex})
		resp, err := watcher.Next(context.Background())
		if err != nil {
			if ignoreError(err) {
				continue
			}
			return err
		}

		afterIndex = resp.Index
		callback(etcdClient, resp.Index, resp.Node.Key, resp.Node.Value)
	}
}

func ignoreError(err error) bool {
	switch err := err.(type) {
	default:
		return false
	case *client.Error:
		return err.Code == client.ErrorCodeEventIndexCleared
	}
}
