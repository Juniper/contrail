package models

// AllocationPoolType

// AllocationPoolType
//proteus:generate
type AllocationPoolType struct {
	VrouterSpecificPool bool   `json:"vrouter_specific_pool"`
	Start               string `json:"start,omitempty"`
	End                 string `json:"end,omitempty"`
}

// MakeAllocationPoolType makes AllocationPoolType
func MakeAllocationPoolType() *AllocationPoolType {
	return &AllocationPoolType{
		//TODO(nati): Apply default
		VrouterSpecificPool: false,
		Start:               "",
		End:                 "",
	}
}

// MakeAllocationPoolTypeSlice() makes a slice of AllocationPoolType
func MakeAllocationPoolTypeSlice() []*AllocationPoolType {
	return []*AllocationPoolType{}
}
