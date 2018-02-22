package models


// MakeLoadbalancerPoolType makes LoadbalancerPoolType
func MakeLoadbalancerPoolType() *LoadbalancerPoolType{
    return &LoadbalancerPoolType{
    //TODO(nati): Apply default
    Status: "",
        Protocol: "",
        SubnetID: "",
        SessionPersistence: "",
        AdminState: false,
        PersistenceCookieName: "",
        StatusDescription: "",
        LoadbalancerMethod: "",
        
    }
}

// MakeLoadbalancerPoolTypeSlice() makes a slice of LoadbalancerPoolType
func MakeLoadbalancerPoolTypeSlice() []*LoadbalancerPoolType {
    return []*LoadbalancerPoolType{}
}


