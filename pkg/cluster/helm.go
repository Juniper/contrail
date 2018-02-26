package cluster

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type helmProvisioner struct {
	provisionCommon
}

func (h *helmProvisioner) provison() error {
	//TODO(ijohnson) Support daemon manager with etcd
	return errors.New("helm is an unsupported provisioner type")
}
