package models


// MakeAllocationPoolType makes AllocationPoolType
func MakeAllocationPoolType() *AllocationPoolType{
    return &AllocationPoolType{
    //TODO(nati): Apply default
    VrouterSpecificPool: false,
        Start: "",
        End: "",
        
    }
}

// MakeAllocationPoolTypeSlice() makes a slice of AllocationPoolType
func MakeAllocationPoolTypeSlice() []*AllocationPoolType {
    return []*AllocationPoolType{}
}


