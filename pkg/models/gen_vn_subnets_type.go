package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVnSubnetsType makes VnSubnetsType
// nolint
func MakeVnSubnetsType() *VnSubnetsType {
	return &VnSubnetsType{
		//TODO(nati): Apply default

		IpamSubnets: MakeIpamSubnetTypeSlice(),

		HostRoutes: MakeRouteTableType(),
	}
}

// MakeVnSubnetsType makes VnSubnetsType
// nolint
func InterfaceToVnSubnetsType(i interface{}) *VnSubnetsType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VnSubnetsType{
		//TODO(nati): Apply default

		IpamSubnets: InterfaceToIpamSubnetTypeSlice(m["ipam_subnets"]),

		HostRoutes: InterfaceToRouteTableType(m["host_routes"]),
	}
}

// MakeVnSubnetsTypeSlice() makes a slice of VnSubnetsType
// nolint
func MakeVnSubnetsTypeSlice() []*VnSubnetsType {
	return []*VnSubnetsType{}
}

// InterfaceToVnSubnetsTypeSlice() makes a slice of VnSubnetsType
// nolint
func InterfaceToVnSubnetsTypeSlice(i interface{}) []*VnSubnetsType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VnSubnetsType{}
	for _, item := range list {
		result = append(result, InterfaceToVnSubnetsType(item))
	}
	return result
}
