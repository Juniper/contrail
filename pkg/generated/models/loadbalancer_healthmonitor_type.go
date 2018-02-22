package models


// MakeLoadbalancerHealthmonitorType makes LoadbalancerHealthmonitorType
func MakeLoadbalancerHealthmonitorType() *LoadbalancerHealthmonitorType{
    return &LoadbalancerHealthmonitorType{
    //TODO(nati): Apply default
    Delay: 0,
        ExpectedCodes: "",
        MaxRetries: 0,
        HTTPMethod: "",
        AdminState: false,
        Timeout: 0,
        URLPath: "",
        MonitorType: "",
        
    }
}

// MakeLoadbalancerHealthmonitorTypeSlice() makes a slice of LoadbalancerHealthmonitorType
func MakeLoadbalancerHealthmonitorTypeSlice() []*LoadbalancerHealthmonitorType {
    return []*LoadbalancerHealthmonitorType{}
}


