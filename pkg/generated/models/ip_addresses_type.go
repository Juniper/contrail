package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeIpAddressesType makes IpAddressesType
func MakeIpAddressesType() *IpAddressesType {
	return &IpAddressesType{
		//TODO(nati): Apply default
		IPAddress: "",
	}
}

// MakeIpAddressesType makes IpAddressesType
func InterfaceToIpAddressesType(i interface{}) *IpAddressesType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &IpAddressesType{
		//TODO(nati): Apply default
		IPAddress: schema.InterfaceToString(m["ip_address"]),
	}
}

// MakeIpAddressesTypeSlice() makes a slice of IpAddressesType
func MakeIpAddressesTypeSlice() []*IpAddressesType {
	return []*IpAddressesType{}
}

// InterfaceToIpAddressesTypeSlice() makes a slice of IpAddressesType
func InterfaceToIpAddressesTypeSlice(i interface{}) []*IpAddressesType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IpAddressesType{}
	for _, item := range list {
		result = append(result, InterfaceToIpAddressesType(item))
	}
	return result
}
