package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeIpamDnsAddressType makes IpamDnsAddressType
func MakeIpamDnsAddressType() *IpamDnsAddressType {
	return &IpamDnsAddressType{
		//TODO(nati): Apply default
		TenantDNSServerAddress: MakeIpAddressesType(),
		VirtualDNSServerName:   "",
	}
}

// MakeIpamDnsAddressType makes IpamDnsAddressType
func InterfaceToIpamDnsAddressType(i interface{}) *IpamDnsAddressType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &IpamDnsAddressType{
		//TODO(nati): Apply default
		TenantDNSServerAddress: InterfaceToIpAddressesType(m["tenant_dns_server_address"]),
		VirtualDNSServerName:   schema.InterfaceToString(m["virtual_dns_server_name"]),
	}
}

// MakeIpamDnsAddressTypeSlice() makes a slice of IpamDnsAddressType
func MakeIpamDnsAddressTypeSlice() []*IpamDnsAddressType {
	return []*IpamDnsAddressType{}
}

// InterfaceToIpamDnsAddressTypeSlice() makes a slice of IpamDnsAddressType
func InterfaceToIpamDnsAddressTypeSlice(i interface{}) []*IpamDnsAddressType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IpamDnsAddressType{}
	for _, item := range list {
		result = append(result, InterfaceToIpamDnsAddressType(item))
	}
	return result
}
