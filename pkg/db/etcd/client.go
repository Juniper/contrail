/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * etcd Client interface
 *
 */

// TODO(Michal): Add some logging to client methods.

package etcd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	conc "github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"

	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/logutil"
)

const (
	kvClientRequestTimeout = 60 * time.Second
)

// Client is an etcd client using clientv3.
type Client struct {
	ETCD *clientv3.Client
	log  *logrus.Entry
}

// Config holds Client configuration.
type Config struct {
	*clientv3.Client // optional clientv3.Client
	clientv3.Config  // config for new clientv3.Client to create
	TLSConfig        TLSConfig
	ServiceName      string
}

// TLSConfig holds Client TLS configuration.
type TLSConfig struct {
	Enabled         bool
	CertificatePath string
	KeyPath         string
	TrustedCAPath   string
}

// NewClientByViper creates etcd client based on global Viper configuration.
func NewClientByViper(serviceName string) (*Client, error) {
	return NewClient(&Config{
		Config: clientv3.Config{
			Endpoints:   viper.GetStringSlice(constants.ETCDEndpointsVK),
			Username:    viper.GetString(constants.ETCDUsernameVK),
			Password:    viper.GetString(constants.ETCDPasswordVK),
			DialTimeout: viper.GetDuration(constants.ETCDDialTimeoutVK),
		},
		TLSConfig: TLSConfig{
			Enabled:         viper.GetBool(constants.ETCDTLSEnabledVK),
			CertificatePath: viper.GetString(constants.ETCDTLSCertificatePathVK),
			KeyPath:         viper.GetString(constants.ETCDTLSKeyPathVK),
			TrustedCAPath:   viper.GetString(constants.ETCDTLSTrustedCAPathVK),
		},
		ServiceName: serviceName,
	})
}

// NewClient creates new etcd Client with given clientv3.Client.
// It creates new clientv3.Client if it is not passed by parameter.
func NewClient(c *Config) (*Client, error) {
	clientv3.SetLogger(grpclog.NewLoggerV2(ioutil.Discard, os.Stdout, os.Stdout))

	var etcd *clientv3.Client
	if c.Client != nil {
		etcd = c.Client
	} else {
		var err error
		etcd, err = newETCDClient(c)
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		ETCD: etcd,
		log:  logutil.NewLogger(fmt.Sprint(c.ServiceName, "-etcd-client")),
	}, nil
}

func newETCDClient(c *Config) (*clientv3.Client, error) {
	if c.TLSConfig.Enabled {
		var err error
		c.TLS, err = transport.TLSInfo{
			CertFile:      c.TLSConfig.CertificatePath,
			KeyFile:       c.TLSConfig.KeyPath,
			TrustedCAFile: c.TLSConfig.TrustedCAPath,
		}.ClientConfig()
		if err != nil {
			return nil, errors.Wrapf(err, "invalid TLS config")
		}
	}

	etcd, err := clientv3.New(c.Config)
	if err != nil {
		return nil, errors.Wrapf(err, "connecting to etcd failed")
	}

	return etcd, nil
}

// Get gets a value in etcd.
func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	kvHandle := clientv3.NewKV(c.ETCD)
	response, err := kvHandle.Get(ctx, key)
	if err != nil || response.Count == 0 {
		return nil, err
	}
	return response.Kvs[0].Value, nil
}

// Put puts value in etcd no matter if it was there or not.
func (c *Client) Put(ctx context.Context, key string, value []byte) error {
	kvHandle := clientv3.NewKV(c.ETCD)

	_, err := kvHandle.Put(ctx, key, string(value))

	return err
}

// Create puts value in etcd if following key didn't exist.
func (c *Client) Create(ctx context.Context, key string, value []byte) error {
	kvHandle := clientv3.NewKV(c.ETCD)

	_, err := kvHandle.Txn(ctx).
		If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Then(clientv3.OpPut(key, string(value))).
		Commit()

	return err
}

// Update puts value in etcd if key existed before.
func (c *Client) Update(ctx context.Context, key, value string) error {
	kvHandle := clientv3.NewKV(c.ETCD)

	_, err := kvHandle.Txn(ctx).
		If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Else(clientv3.OpPut(key, value)).
		Commit()

	return err
}

// Delete deletes a key/value in etcd.
func (c *Client) Delete(ctx context.Context, key string) error {
	kvHandle := clientv3.NewKV(c.ETCD)

	_, err := kvHandle.Txn(ctx).
		If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Else(clientv3.OpDelete(key)).
		Commit()

	return err
}

// WatchRecursive watches a key pattern for changes After an Index and returns channel with messages.
func (c *Client) WatchRecursive(
	ctx context.Context, keyPattern string, afterIndex int64,
) chan Message {
	return c.Watch(ctx, keyPattern, clientv3.WithPrefix(), clientv3.WithRev(afterIndex))

}

// Watch watches a key and returns channel with messages.
func (c *Client) Watch(
	ctx context.Context, key string, opts ...clientv3.OpOption,
) chan Message {
	resultChan := make(chan Message)
	rchan := c.ETCD.Watch(ctx, key, opts...)

	go func() {
		for wresp := range rchan {
			for _, ev := range wresp.Events {
				resultChan <- NewMessage(ev)
			}
		}
		close(resultChan)
	}()

	return resultChan
}

// DoInTransaction wraps clientv3 transaction and wraps conc.STM with own Txn.
func (c *Client) DoInTransaction(ctx context.Context, do func(context.Context) error) error {
	if txn := GetTxn(ctx); txn != nil {
		// Transaction already in context
		return do(ctx)
	}
	// New transaction required

	ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
	defer cancel()

	_, err := conc.NewSTM(c.ETCD, func(stm conc.STM) error {
		return do(WithTxn(ctx, stmTxn{stm, c.log}))
	}, conc.WithAbortContext(ctx))
	return err
}

// Close closes client.
func (c *Client) Close() error {
	return c.ETCD.Close()
}
