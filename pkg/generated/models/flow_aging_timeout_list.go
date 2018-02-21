package models


// MakeFlowAgingTimeoutList makes FlowAgingTimeoutList
func MakeFlowAgingTimeoutList() *FlowAgingTimeoutList{
    return &FlowAgingTimeoutList{
    //TODO(nati): Apply default
    
            
                FlowAgingTimeout:  MakeFlowAgingTimeoutSlice(),
            
        
    }
}

// MakeFlowAgingTimeoutListSlice() makes a slice of FlowAgingTimeoutList
func MakeFlowAgingTimeoutListSlice() []*FlowAgingTimeoutList {
    return []*FlowAgingTimeoutList{}
}


