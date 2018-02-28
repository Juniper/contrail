package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeIpamSubnets makes IpamSubnets
// nolint
func MakeIpamSubnets() *IpamSubnets {
	return &IpamSubnets{
		//TODO(nati): Apply default

		Subnets: MakeIpamSubnetTypeSlice(),
	}
}

// MakeIpamSubnets makes IpamSubnets
// nolint
func InterfaceToIpamSubnets(i interface{}) *IpamSubnets {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &IpamSubnets{
		//TODO(nati): Apply default

		Subnets: InterfaceToIpamSubnetTypeSlice(m["subnets"]),
	}
}

// MakeIpamSubnetsSlice() makes a slice of IpamSubnets
// nolint
func MakeIpamSubnetsSlice() []*IpamSubnets {
	return []*IpamSubnets{}
}

// InterfaceToIpamSubnetsSlice() makes a slice of IpamSubnets
// nolint
func InterfaceToIpamSubnetsSlice(i interface{}) []*IpamSubnets {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IpamSubnets{}
	for _, item := range list {
		result = append(result, InterfaceToIpamSubnets(item))
	}
	return result
}
