package testutil

import (
	"context"
	"testing"
	"time"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
)

// Etcd connection constants
const (
	EtcdEndpoint = "localhost:2379"

	dialTimeout = 60 * time.Second
)

// EtcdClient is etcd client extending etcd.clientv3 with test functionality and using etcd v3 API.
type EtcdClient struct {
	*clientv3.Client
	log *logrus.Entry
}

// NewEtcdClient is a constructor of etcd client.
func NewEtcdClient(t *testing.T) *EtcdClient {
	l := pkglog.NewLogger("etcd-client")
	l.WithFields(logrus.Fields{"endpoint": EtcdEndpoint, "dial-timeout": dialTimeout}).Debug("Connecting")
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{EtcdEndpoint},
		DialTimeout: dialTimeout,
	})
	if err != nil {
		t.Fatalf("Cannot connect to etcd: %s", err)
	}

	return &EtcdClient{
		Client: c,
		log:    l,
	}
}

// GetAllFromPrefix return etcd Get response with all keys starting from given prefix.
func (e *EtcdClient) GetAllFromPrefix(t *testing.T, prefix string) *clientv3.GetResponse {
	r, err := e.Get(context.Background(), prefix, clientv3.WithPrefix())
	logrus.WithFields(logrus.Fields{"prefix": prefix, "response": r}).Debug("Received etcd Get response")
	if err != nil {
		t.Errorf("Cannot get resource from etcd: %s", err)
	}
	return r
}

// CloseClient closes connection to etcd.
func (e *EtcdClient) CloseClient(t *testing.T) {
	err := e.Close()
	if err != nil {
		t.Errorf("Cannot close connection to etcd: %s", err)
	}
}
