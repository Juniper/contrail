package models
// IpamSubnets



import "encoding/json"

// IpamSubnets 
//proteus:generate
type IpamSubnets struct {

    Subnets []*IpamSubnetType `json:"subnets,omitempty"`


}



// String returns json representation of the object
func (model *IpamSubnets) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeIpamSubnets makes IpamSubnets
func MakeIpamSubnets() *IpamSubnets{
    return &IpamSubnets{
    //TODO(nati): Apply default
    
            
                Subnets:  MakeIpamSubnetTypeSlice(),
            
        
    }
}



// MakeIpamSubnetsSlice() makes a slice of IpamSubnets
func MakeIpamSubnetsSlice() []*IpamSubnets {
    return []*IpamSubnets{}
}
