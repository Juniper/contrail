package models
// IpamType



import "encoding/json"

// IpamType 
//proteus:generate
type IpamType struct {

    IpamMethod IpamMethodType `json:"ipam_method,omitempty"`
    IpamDNSMethod IpamDnsMethodType `json:"ipam_dns_method,omitempty"`
    IpamDNSServer *IpamDnsAddressType `json:"ipam_dns_server,omitempty"`
    DHCPOptionList *DhcpOptionsListType `json:"dhcp_option_list,omitempty"`
    HostRoutes *RouteTableType `json:"host_routes,omitempty"`
    CidrBlock *SubnetType `json:"cidr_block,omitempty"`


}



// String returns json representation of the object
func (model *IpamType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeIpamType makes IpamType
func MakeIpamType() *IpamType{
    return &IpamType{
    //TODO(nati): Apply default
    IpamMethod: MakeIpamMethodType(),
        IpamDNSMethod: MakeIpamDnsMethodType(),
        IpamDNSServer: MakeIpamDnsAddressType(),
        DHCPOptionList: MakeDhcpOptionsListType(),
        HostRoutes: MakeRouteTableType(),
        CidrBlock: MakeSubnetType(),
        
    }
}



// MakeIpamTypeSlice() makes a slice of IpamType
func MakeIpamTypeSlice() []*IpamType {
    return []*IpamType{}
}
