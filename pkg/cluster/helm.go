package cluster

import (
	"errors"
)

type helmProvisioner struct {
	provisionCommon
}

func (h *helmProvisioner) provision() error {
	//TODO(ijohnson) Support daemon manager with etcd
	return errors.New("helm is an unsupported provisioner type")
}
