package cluster

import (
	"errors"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

type manager interface {
	manage()
}

type oneShotManager struct {
	cluster *Cluster
	log     *logrus.Entry
}

func (m *oneShotManager) manage() error {
	provisioner := newProvisioner(m)
	err := provisioner.provison()
	if err != nil {
		return err
	}
	return nil
}

func newOneShotManager(cluster *Cluster) (*oneShotManager, error) {
	return &oneShotManager{
		cluster: cluster,
		log:     pkglog.NewLogger("oneshot-manager"),
	}, nil
}

func newManager(cluster *Cluster) (manager, error) {
	switch cluster.managerType {
	case "oneshot":
		return newOneShotManager(cluster)
	}
	//TODO(ijohnson) Support daemon manager with etcd
	return nil, errors.New("unsupported manager type")
}
