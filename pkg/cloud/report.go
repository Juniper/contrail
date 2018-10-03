package cloud

import (
	"errors"
)

func (p *provisionData) reportStatus() error {
	//TODO(madhukar) Support daemon manager with etcd
	return errors.New("Cannot report status at this time")
}
