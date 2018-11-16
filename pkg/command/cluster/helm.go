package cluster

import (
	"errors"
)

type helmProvisioner struct {
	provisionCluster
}

func (h *helmProvisioner) Provision() error {
	//TODO(ijohnson) Support daemon manager with etcd
	return errors.New("helm is an unsupported provisioner type")
}
