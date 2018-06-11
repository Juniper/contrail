package integration

import (
	"testing"

	etcdlock "github.com/DavidCai1993/etcd-lock"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

// EtcdLocker is etcd locker extending etcdlock.Locker with test functionality and using etcd gRPC protocol.
type EtcdLocker struct {
	*etcdlock.Locker
	log *logrus.Entry
}

// NewEtcdLocker is a constructor of etcd client.
func NewEtcdLocker(t *testing.T) *EtcdLocker {
	l := pkglog.NewLogger("etcd-locker")
	l.WithFields(logrus.Fields{"endpoint": etcdEndpoint}).Debug("Connecting")

	locker, err := etcdlock.NewLocker(etcdlock.LockerOptions{
		Address:     etcdEndpoint,
		DialOptions: []grpc.DialOption{grpc.WithInsecure()},
	})
	require.NoError(t, err, "connecting etcdlock failed")

	return &EtcdLocker{
		Locker: locker,
		log:    l,
	}
}
