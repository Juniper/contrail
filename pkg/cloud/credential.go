package cloud

import (
	"errors"
)

func (p *provisionData) selectCloudCredential() error {
	//TODO(madhukar) Support daemon manager with etcd
	return errors.New("selecting cloud credential is not supported")
}
