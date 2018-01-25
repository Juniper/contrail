package models
// FloatingIpPoolSubnetType



import "encoding/json"

// FloatingIpPoolSubnetType 
//proteus:generate
type FloatingIpPoolSubnetType struct {

    SubnetUUID []string `json:"subnet_uuid,omitempty"`


}



// String returns json representation of the object
func (model *FloatingIpPoolSubnetType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFloatingIpPoolSubnetType makes FloatingIpPoolSubnetType
func MakeFloatingIpPoolSubnetType() *FloatingIpPoolSubnetType{
    return &FloatingIpPoolSubnetType{
    //TODO(nati): Apply default
    SubnetUUID: []string{},
        
    }
}



// MakeFloatingIpPoolSubnetTypeSlice() makes a slice of FloatingIpPoolSubnetType
func MakeFloatingIpPoolSubnetTypeSlice() []*FloatingIpPoolSubnetType {
    return []*FloatingIpPoolSubnetType{}
}
