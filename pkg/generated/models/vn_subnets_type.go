package models

// VnSubnetsType

// VnSubnetsType
//proteus:generate
type VnSubnetsType struct {
	IpamSubnets []*IpamSubnetType `json:"ipam_subnets,omitempty"`
	HostRoutes  *RouteTableType   `json:"host_routes,omitempty"`
}

// MakeVnSubnetsType makes VnSubnetsType
func MakeVnSubnetsType() *VnSubnetsType {
	return &VnSubnetsType{
		//TODO(nati): Apply default

		IpamSubnets: MakeIpamSubnetTypeSlice(),

		HostRoutes: MakeRouteTableType(),
	}
}

// MakeVnSubnetsTypeSlice() makes a slice of VnSubnetsType
func MakeVnSubnetsTypeSlice() []*VnSubnetsType {
	return []*VnSubnetsType{}
}
