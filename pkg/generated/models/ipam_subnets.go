package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeIpamSubnets makes IpamSubnets
func MakeIpamSubnets() *IpamSubnets {
	return &IpamSubnets{
		//TODO(nati): Apply default

		Subnets: MakeIpamSubnetTypeSlice(),
	}
}

// MakeIpamSubnets makes IpamSubnets
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
func MakeIpamSubnetsSlice() []*IpamSubnets {
	return []*IpamSubnets{}
}

// InterfaceToIpamSubnetsSlice() makes a slice of IpamSubnets
func InterfaceToIpamSubnetsSlice(i interface{}) []*IpamSubnets {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IpamSubnets{}
	for _, item := range list {
		result = append(result, InterfaceToIpamSubnets(item))
	}
	return result
}
