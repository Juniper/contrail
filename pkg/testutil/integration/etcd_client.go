package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	etcdEndpoint = "localhost:2379"
	dialTimeout  = 60 * time.Second
)

// EtcdClient is etcd client extending etcd.clientv3 with test functionality and using etcd v3 API.
type EtcdClient struct {
	*clientv3.Client
	log *logrus.Entry
}

// NewEtcdClient is a constructor of etcd client.
func NewEtcdClient(t *testing.T) *EtcdClient {
	l := pkglog.NewLogger("etcd-client")
	l.WithFields(logrus.Fields{"endpoint": etcdEndpoint, "dial-timeout": dialTimeout}).Debug("Connecting")
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdEndpoint},
		DialTimeout: dialTimeout,
	})
	require.NoError(t, err, "connecting etcd failed")

	return &EtcdClient{
		Client: c,
		log:    l,
	}
}

// GetAllWithPrefix returns etcd Get response for all keys starting from given prefix.
func (e *EtcdClient) GetAllWithPrefix(t *testing.T, prefix string) *clientv3.GetResponse {
	r, err := e.Get(context.Background(), prefix, clientv3.WithPrefix())
	assert.NoError(t, err, "getting resource from etcd failed")
	e.log.WithFields(logrus.Fields{"prefix": prefix, "response": r}).Debug("Received etcd Get response")
	return r
}

// CheckKeyDoesNotExist checks that there is no value on given key.
func (e *EtcdClient) CheckKeyDoesNotExist(t *testing.T, key string) {
	gr := e.GetAllWithPrefix(t, key)
	assert.Equal(t, int64(0), gr.Count, fmt.Sprintf("key %v should be empty", key))
}

// Close closes connection to etcd.
func (e *EtcdClient) Close(t *testing.T) {
	err := e.Client.Close()
	assert.NoError(t, err, "closing etcd connection failed")
}
