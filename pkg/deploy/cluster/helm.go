package cluster

import (
	"github.com/pkg/errors"
)

type helmDeployer struct {
	deployCluster
}

func (h *helmDeployer) Deploy() error {
	//TODO(ijohnson) Support daemon manager with etcd
	return errors.New("helm is an unsupported provisioner type")
}
