package cluster

import (
	"errors"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

// manager inerface to manage clusters
type manager interface {
	manage() error
}

type oneShotManager struct {
	cluster *Cluster
	log     *logrus.Entry
}

// Manage a cluster once
func (o *oneShotManager) manage() error {
	provisioner, err := newProvisioner(o.cluster)
	if err != nil {
		return err
	}
	return provisioner.provision()
}

func newOneShotManager(cluster *Cluster) (*oneShotManager, error) {
	// create logger for oneshot manager
	logger := pkglog.NewFileLogger("oneshot-manager", cluster.config.LogFile)
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	return &oneShotManager{
		cluster: cluster,
		log:     logger,
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
