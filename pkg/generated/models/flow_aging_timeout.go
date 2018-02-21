package models


// MakeFlowAgingTimeout makes FlowAgingTimeout
func MakeFlowAgingTimeout() *FlowAgingTimeout{
    return &FlowAgingTimeout{
    //TODO(nati): Apply default
    TimeoutInSeconds: 0,
        Protocol: "",
        Port: 0,
        
    }
}

// MakeFlowAgingTimeoutSlice() makes a slice of FlowAgingTimeout
func MakeFlowAgingTimeoutSlice() []*FlowAgingTimeout {
    return []*FlowAgingTimeout{}
}


