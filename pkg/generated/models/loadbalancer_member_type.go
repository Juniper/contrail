package models


// MakeLoadbalancerMemberType makes LoadbalancerMemberType
func MakeLoadbalancerMemberType() *LoadbalancerMemberType{
    return &LoadbalancerMemberType{
    //TODO(nati): Apply default
    Status: "",
        StatusDescription: "",
        Weight: 0,
        AdminState: false,
        Address: "",
        ProtocolPort: 0,
        
    }
}

// MakeLoadbalancerMemberTypeSlice() makes a slice of LoadbalancerMemberType
func MakeLoadbalancerMemberTypeSlice() []*LoadbalancerMemberType {
    return []*LoadbalancerMemberType{}
}


