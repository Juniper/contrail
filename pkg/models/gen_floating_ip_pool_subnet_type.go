package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFloatingIpPoolSubnetType makes FloatingIpPoolSubnetType
// nolint
func MakeFloatingIpPoolSubnetType() *FloatingIpPoolSubnetType {
	return &FloatingIpPoolSubnetType{
		//TODO(nati): Apply default
		SubnetUUID: []string{},
	}
}

// MakeFloatingIpPoolSubnetType makes FloatingIpPoolSubnetType
// nolint
func InterfaceToFloatingIpPoolSubnetType(i interface{}) *FloatingIpPoolSubnetType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FloatingIpPoolSubnetType{
		//TODO(nati): Apply default
		SubnetUUID: common.InterfaceToStringList(m["subnet_uuid"]),
	}
}

// MakeFloatingIpPoolSubnetTypeSlice() makes a slice of FloatingIpPoolSubnetType
// nolint
func MakeFloatingIpPoolSubnetTypeSlice() []*FloatingIpPoolSubnetType {
	return []*FloatingIpPoolSubnetType{}
}

// InterfaceToFloatingIpPoolSubnetTypeSlice() makes a slice of FloatingIpPoolSubnetType
// nolint
func InterfaceToFloatingIpPoolSubnetTypeSlice(i interface{}) []*FloatingIpPoolSubnetType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FloatingIpPoolSubnetType{}
	for _, item := range list {
		result = append(result, InterfaceToFloatingIpPoolSubnetType(item))
	}
	return result
}
