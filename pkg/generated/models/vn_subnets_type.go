package models
// VnSubnetsType



import "encoding/json"

// VnSubnetsType 
//proteus:generate
type VnSubnetsType struct {

    IpamSubnets []*IpamSubnetType `json:"ipam_subnets,omitempty"`
    HostRoutes *RouteTableType `json:"host_routes,omitempty"`


}



// String returns json representation of the object
func (model *VnSubnetsType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVnSubnetsType makes VnSubnetsType
func MakeVnSubnetsType() *VnSubnetsType{
    return &VnSubnetsType{
    //TODO(nati): Apply default
    
            
                IpamSubnets:  MakeIpamSubnetTypeSlice(),
            
        HostRoutes: MakeRouteTableType(),
        
    }
}



// MakeVnSubnetsTypeSlice() makes a slice of VnSubnetsType
func MakeVnSubnetsTypeSlice() []*VnSubnetsType {
    return []*VnSubnetsType{}
}
