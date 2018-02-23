package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeVnSubnetsType makes VnSubnetsType
func MakeVnSubnetsType() *VnSubnetsType {
	return &VnSubnetsType{
		//TODO(nati): Apply default

		IpamSubnets: MakeIpamSubnetTypeSlice(),

		HostRoutes: MakeRouteTableType(),
	}
}

// MakeVnSubnetsType makes VnSubnetsType
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
func MakeVnSubnetsTypeSlice() []*VnSubnetsType {
	return []*VnSubnetsType{}
}

// InterfaceToVnSubnetsTypeSlice() makes a slice of VnSubnetsType
func InterfaceToVnSubnetsTypeSlice(i interface{}) []*VnSubnetsType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VnSubnetsType{}
	for _, item := range list {
		result = append(result, InterfaceToVnSubnetsType(item))
	}
	return result
}
