package models

// IpamDnsAddressType

import "encoding/json"

// IpamDnsAddressType
type IpamDnsAddressType struct {
	TenantDNSServerAddress *IpAddressesType `json:"tenant_dns_server_address"`
	VirtualDNSServerName   string           `json:"virtual_dns_server_name"`
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
		TenantDNSServerAddress: MakeIpAddressesType(),
		VirtualDNSServerName:   "",
	}
}

// InterfaceToIpamDnsAddressType makes IpamDnsAddressType from interface
func InterfaceToIpamDnsAddressType(iData interface{}) *IpamDnsAddressType {
	data := iData.(map[string]interface{})
	return &IpamDnsAddressType{
		TenantDNSServerAddress: InterfaceToIpAddressesType(data["tenant_dns_server_address"]),

		//{"description":"In case of tenant DNS server method, Ip address of DNS server. This will be given in DHCP","type":"object","properties":{"ip_address":{"type":"string"}}}
		VirtualDNSServerName: data["virtual_dns_server_name"].(string),

		//{"description":"In case of virtual DNS server, name of virtual DNS server","type":"string"}

	}
}

// InterfaceToIpamDnsAddressTypeSlice makes a slice of IpamDnsAddressType from interface
func InterfaceToIpamDnsAddressTypeSlice(data interface{}) []*IpamDnsAddressType {
	list := data.([]interface{})
	result := MakeIpamDnsAddressTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamDnsAddressType(item))
	}
	return result
}

// MakeIpamDnsAddressTypeSlice() makes a slice of IpamDnsAddressType
func MakeIpamDnsAddressTypeSlice() []*IpamDnsAddressType {
	return []*IpamDnsAddressType{}
}
