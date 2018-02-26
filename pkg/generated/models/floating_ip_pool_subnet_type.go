package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeFloatingIpPoolSubnetType makes FloatingIpPoolSubnetType
func MakeFloatingIpPoolSubnetType() *FloatingIpPoolSubnetType {
	return &FloatingIpPoolSubnetType{
		//TODO(nati): Apply default
		SubnetUUID: []string{},
	}
}

// MakeFloatingIpPoolSubnetType makes FloatingIpPoolSubnetType
func InterfaceToFloatingIpPoolSubnetType(i interface{}) *FloatingIpPoolSubnetType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FloatingIpPoolSubnetType{
		//TODO(nati): Apply default
		SubnetUUID: schema.InterfaceToStringList(m["subnet_uuid"]),
	}
}

// MakeFloatingIpPoolSubnetTypeSlice() makes a slice of FloatingIpPoolSubnetType
func MakeFloatingIpPoolSubnetTypeSlice() []*FloatingIpPoolSubnetType {
	return []*FloatingIpPoolSubnetType{}
}

// InterfaceToFloatingIpPoolSubnetTypeSlice() makes a slice of FloatingIpPoolSubnetType
func InterfaceToFloatingIpPoolSubnetTypeSlice(i interface{}) []*FloatingIpPoolSubnetType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FloatingIpPoolSubnetType{}
	for _, item := range list {
		result = append(result, InterfaceToFloatingIpPoolSubnetType(item))
	}
	return result
}
