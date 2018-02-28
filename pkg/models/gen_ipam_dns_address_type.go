package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeIpamDnsAddressType makes IpamDnsAddressType
// nolint
func MakeIpamDnsAddressType() *IpamDnsAddressType {
	return &IpamDnsAddressType{
		//TODO(nati): Apply default
		TenantDNSServerAddress: MakeIpAddressesType(),
		VirtualDNSServerName:   "",
	}
}

// MakeIpamDnsAddressType makes IpamDnsAddressType
// nolint
func InterfaceToIpamDnsAddressType(i interface{}) *IpamDnsAddressType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &IpamDnsAddressType{
		//TODO(nati): Apply default
		TenantDNSServerAddress: InterfaceToIpAddressesType(m["tenant_dns_server_address"]),
		VirtualDNSServerName:   common.InterfaceToString(m["virtual_dns_server_name"]),
	}
}

// MakeIpamDnsAddressTypeSlice() makes a slice of IpamDnsAddressType
// nolint
func MakeIpamDnsAddressTypeSlice() []*IpamDnsAddressType {
	return []*IpamDnsAddressType{}
}

// InterfaceToIpamDnsAddressTypeSlice() makes a slice of IpamDnsAddressType
// nolint
func InterfaceToIpamDnsAddressTypeSlice(i interface{}) []*IpamDnsAddressType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IpamDnsAddressType{}
	for _, item := range list {
		result = append(result, InterfaceToIpamDnsAddressType(item))
	}
	return result
}
