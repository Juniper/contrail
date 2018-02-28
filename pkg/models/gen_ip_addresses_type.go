package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeIpAddressesType makes IpAddressesType
// nolint
func MakeIpAddressesType() *IpAddressesType {
	return &IpAddressesType{
		//TODO(nati): Apply default
		IPAddress: "",
	}
}

// MakeIpAddressesType makes IpAddressesType
// nolint
func InterfaceToIpAddressesType(i interface{}) *IpAddressesType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &IpAddressesType{
		//TODO(nati): Apply default
		IPAddress: common.InterfaceToString(m["ip_address"]),
	}
}

// MakeIpAddressesTypeSlice() makes a slice of IpAddressesType
// nolint
func MakeIpAddressesTypeSlice() []*IpAddressesType {
	return []*IpAddressesType{}
}

// InterfaceToIpAddressesTypeSlice() makes a slice of IpAddressesType
// nolint
func InterfaceToIpAddressesTypeSlice(i interface{}) []*IpAddressesType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IpAddressesType{}
	for _, item := range list {
		result = append(result, InterfaceToIpAddressesType(item))
	}
	return result
}
