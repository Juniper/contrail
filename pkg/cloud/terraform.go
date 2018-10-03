package cloud

import (
	"errors"
)

func (p *provisionData) invokeTerraform() error {
	//TODO(madhukar) Support daemon manager with etcd
	return errors.New("invoking terraform is not supported yet")
}
