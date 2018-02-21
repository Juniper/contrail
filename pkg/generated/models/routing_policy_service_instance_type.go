package models


// MakeRoutingPolicyServiceInstanceType makes RoutingPolicyServiceInstanceType
func MakeRoutingPolicyServiceInstanceType() *RoutingPolicyServiceInstanceType{
    return &RoutingPolicyServiceInstanceType{
    //TODO(nati): Apply default
    RightSequence: "",
        LeftSequence: "",
        
    }
}

// MakeRoutingPolicyServiceInstanceTypeSlice() makes a slice of RoutingPolicyServiceInstanceType
func MakeRoutingPolicyServiceInstanceTypeSlice() []*RoutingPolicyServiceInstanceType {
    return []*RoutingPolicyServiceInstanceType{}
}


