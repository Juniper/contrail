package models

// IpamDnsAddressType

// IpamDnsAddressType
//proteus:generate
type IpamDnsAddressType struct {
	TenantDNSServerAddress *IpAddressesType `json:"tenant_dns_server_address,omitempty"`
	VirtualDNSServerName   string           `json:"virtual_dns_server_name,omitempty"`
}

// MakeIpamDnsAddressType makes IpamDnsAddressType
func MakeIpamDnsAddressType() *IpamDnsAddressType {
	return &IpamDnsAddressType{
		//TODO(nati): Apply default
		TenantDNSServerAddress: MakeIpAddressesType(),
		VirtualDNSServerName:   "",
	}
}

// MakeIpamDnsAddressTypeSlice() makes a slice of IpamDnsAddressType
func MakeIpamDnsAddressTypeSlice() []*IpamDnsAddressType {
	return []*IpamDnsAddressType{}
}
