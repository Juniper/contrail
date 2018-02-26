package cluster

import (
	"errors"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

type helmProvisioner struct {
	manager     *Manager
	clusterID   string
	action      string
	clusterInfo map[interface{}]interface{}
	log         *logrus.Entry
	reporter    *Reporter
}

func (h *helmProvisioner) provison() error {
	//TODO(ijohnson) Support daemon manager with etcd
	return nil, errors.New("helm is an unsupported provisioner type")
}
