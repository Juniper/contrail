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

// Etcd is etcd client extending etcd.clientv3 with test functionality and using etcd v3 API.
type Etcd struct {
	*clientv3.Client
	log *logrus.Entry
}

// NewEtcd is a constructor of Etcd client.
func NewEtcd(t *testing.T) *Etcd {
	l := pkglog.NewLogger("etcd-client")
	l.WithFields(logrus.Fields{"endpoint": EtcdEndpoint, "dial-timeout": dialTimeout}).Debug("Connecting")
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{EtcdEndpoint},
		DialTimeout: dialTimeout,
	})
	if err != nil {
		t.Fatalf("Cannot connect to Etcd: %s", err)
	}

	return &Etcd{
		Client: c,
		log:    l,
	}
}

// GetAllFromPrefix return etcd Get response with all keys starting from given prefix.
func (e *Etcd) GetAllFromPrefix(t *testing.T, prefix string) *clientv3.GetResponse {
	r, err := e.Get(context.Background(), prefix, clientv3.WithPrefix())
	logrus.WithFields(logrus.Fields{"prefix": prefix, "response": r}).Debug("Received etcd Get response")
	if err != nil {
		t.Fatalf("Cannot get resource from etcd: %s", err)
	}
	if r.More {
		t.Fatal("Received incomplete etcd response, which currently is not handled")
	}
	return r
}

// CloseClient closes connection to Etcd.
func (e *Etcd) CloseClient(t *testing.T) {
	err := e.Close()
	if err != nil {
		t.Errorf("Cannot close connection to Etcd: %s", err)
	}
}
