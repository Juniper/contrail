package models

// IpamDnsAddressType

import "encoding/json"

// IpamDnsAddressType
type IpamDnsAddressType struct {
	TenantDNSServerAddress *IpAddressesType `json:"tenant_dns_server_address,omitempty"`
	VirtualDNSServerName   string           `json:"virtual_dns_server_name,omitempty"`
}

// String returns json representation of the object
func (model *IpamDnsAddressType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIpamDnsAddressType makes IpamDnsAddressType
func MakeIpamDnsAddressType() *IpamDnsAddressType {
	return &IpamDnsAddressType{
		//TODO(nati): Apply default
		VirtualDNSServerName:   "",
		TenantDNSServerAddress: MakeIpAddressesType(),
	}
}

// MakeIpamDnsAddressTypeSlice() makes a slice of IpamDnsAddressType
func MakeIpamDnsAddressTypeSlice() []*IpamDnsAddressType {
	return []*IpamDnsAddressType{}
}
