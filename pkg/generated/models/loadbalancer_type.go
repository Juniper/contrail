package models


// MakeLoadbalancerType makes LoadbalancerType
func MakeLoadbalancerType() *LoadbalancerType{
    return &LoadbalancerType{
    //TODO(nati): Apply default
    Status: "",
        ProvisioningStatus: "",
        AdminState: false,
        VipAddress: "",
        VipSubnetID: "",
        OperatingStatus: "",
        
    }
}

// MakeLoadbalancerTypeSlice() makes a slice of LoadbalancerType
func MakeLoadbalancerTypeSlice() []*LoadbalancerType {
    return []*LoadbalancerType{}
}


