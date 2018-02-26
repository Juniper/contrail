package cluster

import (
	"errors"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

// Manager inerface to manage clusters
type Manager interface {
	Manage()
}

// OneShotManager represents a one-shot cluster manager
type OneShotManager struct {
	cluster *Cluster
	log     *logrus.Entry
}

// Manage a clustes once
func (o *OneShotManager) Manage() error {
	provisioner, err := newProvisioner(o)
	if err != nil {
		return err
	}
	err := provisioner.provison()
	if err != nil {
		return err
	}
	return nil
}

func newOneShotManager(cluster *Cluster) (*OneShotManager, error) {
	return &OneShotManager{
		cluster: cluster,
		log:     pkglog.NewLogger("oneshot-manager"),
	}, nil
}

func newManager(cluster *Cluster) (Manager, error) {
	switch cluster.managerType {
	case "oneshot":
		return newOneShotManager(cluster)
	}
	//TODO(ijohnson) Support daemon manager with etcd
	return nil, errors.New("unsupported manager type")
}
