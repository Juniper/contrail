package logic

import (
	"github.com/Juniper/contrail/pkg/format"
)

func (s *Subnet) ApplyMap(m map[string]interface{}) error {
	_, ok := m[SubnetFieldIpamFQName].(string)
	if ok {
		delete(m, SubnetFieldIpamFQName)
	}
	type alias Subnet
	obj := (*alias)(s)
	return format.ApplyMap(m, &obj)
}


