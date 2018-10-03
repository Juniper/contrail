package cloud

import (
	"errors"
)

func (p *provisionData) modifyTopology() error {
	//TODO(madhukar) Support daemon manager with etcd
	return errors.New("modify topology yml file for multi-cloud is not supported")
}
