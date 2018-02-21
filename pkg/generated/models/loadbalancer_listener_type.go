package models


// MakeLoadbalancerListenerType makes LoadbalancerListenerType
func MakeLoadbalancerListenerType() *LoadbalancerListenerType{
    return &LoadbalancerListenerType{
    //TODO(nati): Apply default
    DefaultTLSContainer: "",
        Protocol: "",
        ConnectionLimit: 0,
        AdminState: false,
        SniContainers: []string{},
        ProtocolPort: 0,
        
    }
}

// MakeLoadbalancerListenerTypeSlice() makes a slice of LoadbalancerListenerType
func MakeLoadbalancerListenerTypeSlice() []*LoadbalancerListenerType {
    return []*LoadbalancerListenerType{}
}


