package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pkglog "github.com/Juniper/contrail/pkg/log"
)

// Integration test settings.
const (
	EtcdEndpoint       = "localhost:2379"
	EtcdDialTimeout    = 60 * time.Second
	EtcdRequestTimeout = 60 * time.Second
)

// EtcdClient is etcd client extending etcd.clientv3 with test functionality and using etcd v3 API.
type EtcdClient struct {
	*clientv3.Client
	log *logrus.Entry
}

// NewEtcdClient is a constructor of etcd client.
// After usage Close() needs to be called to close underlying connections.
func NewEtcdClient(t *testing.T) *EtcdClient {
	l := pkglog.NewLogger("etcd-client")
	l.WithFields(logrus.Fields{"endpoint": EtcdEndpoint, "dial-timeout": EtcdDialTimeout}).Debug("Connecting")
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{EtcdEndpoint},
		DialTimeout: EtcdDialTimeout,
	})
	require.NoError(t, err, "connecting etcd failed")

	return &EtcdClient{
		Client: c,
		log:    l,
	}
}

// GetKey gets etcd key.
func (e *EtcdClient) GetKey(t *testing.T, key string, opts ...clientv3.OpOption) *clientv3.GetResponse {
	ctx, cancel := context.WithTimeout(context.Background(), EtcdRequestTimeout)
	defer cancel()

	r, err := e.Get(ctx, key, opts...)
	assert.NoError(t, err, fmt.Sprintf("getting etcd resource failed\n response: %+v", r))
	e.log.WithFields(logrus.Fields{"prefix": key, "response": r}).Debug("Received etcd Get response")
	return r
}

// DeleteKey deletes etcd key.
func (e *EtcdClient) DeleteKey(t *testing.T, key string, opts ...clientv3.OpOption) {
	ctx, cancel := context.WithTimeout(context.Background(), EtcdRequestTimeout)
	defer cancel()

	r, err := e.Delete(ctx, key, opts...)
	assert.NoError(t, err, fmt.Sprintf("deleting etcd resource from etcd failed\n response: %+v", r))
}

// CheckKeyDoesNotExist checks that there is no value on given key.
func (e *EtcdClient) CheckKeyDoesNotExist(t *testing.T, key string) {
	gr := e.GetKey(t, key, clientv3.WithPrefix())
	assert.Equal(t, int64(0), gr.Count, fmt.Sprintf("key %v should be empty", key))
}

// WatchKey watches value changes for provided key and returns collect method that collect captured values.
func (e *EtcdClient) WatchKey(key string) (collect func() []string) {
	var result []string
	ctx, cancel := context.WithCancel(context.Background())
	wchan := e.Client.Watch(ctx, key)

	go func() {
		for val := range wchan {
			if len(val.Events) > 0 {
				result = append(result, string(val.Events[0].Kv.Value))
			}
		}
	}()

	return func() (vals []string) {
		cancel()
		return result
	}
}

// GetString gets a string value in Etcd
func (e *EtcdClient) GetString(t *testing.T, key string) (value string, revision int64) {
	err := e.Client.Sync(context.Background())
	assert.NoError(t, err)

	kvHandle := clientv3.NewKV(e.Client)
	response, err := kvHandle.Get(context.Background(), key)
	require.NoError(t, err)
	require.NotEmpty(t, response.Kvs)

	return string(response.Kvs[0].Value), response.Header.Revision
}

// ExpectValue gets key and checks if value and revision match.
func (e *EtcdClient) ExpectValue(t *testing.T, key string, value string, revision int64) {
	nextVal, nextRev := e.GetString(t, key)
	assert.Equal(t, value, nextVal)
	assert.Equal(t, revision, nextRev)
}

// Close closes connection to etcd.
func (e *EtcdClient) Close(t *testing.T) {
	err := e.Client.Close()
	assert.NoError(t, err, "closing etcd.clientv3.Client failed")
}
