package models

// IpamType

import "encoding/json"

// IpamType
type IpamType struct {
	IpamDNSMethod  IpamDnsMethodType    `json:"ipam_dns_method,omitempty"`
	IpamDNSServer  *IpamDnsAddressType  `json:"ipam_dns_server,omitempty"`
	DHCPOptionList *DhcpOptionsListType `json:"dhcp_option_list,omitempty"`
	HostRoutes     *RouteTableType      `json:"host_routes,omitempty"`
	CidrBlock      *SubnetType          `json:"cidr_block,omitempty"`
	IpamMethod     IpamMethodType       `json:"ipam_method,omitempty"`
}

// String returns json representation of the object
func (model *IpamType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIpamType makes IpamType
func MakeIpamType() *IpamType {
	return &IpamType{
		//TODO(nati): Apply default
		IpamDNSMethod:  MakeIpamDnsMethodType(),
		IpamDNSServer:  MakeIpamDnsAddressType(),
		DHCPOptionList: MakeDhcpOptionsListType(),
		HostRoutes:     MakeRouteTableType(),
		CidrBlock:      MakeSubnetType(),
		IpamMethod:     MakeIpamMethodType(),
	}
}

// MakeIpamTypeSlice() makes a slice of IpamType
func MakeIpamTypeSlice() []*IpamType {
	return []*IpamType{}
}
