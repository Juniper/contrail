package models

// AllocationPoolType

import "encoding/json"

// AllocationPoolType
type AllocationPoolType struct {
	Start               string `json:"start,omitempty"`
	End                 string `json:"end,omitempty"`
	VrouterSpecificPool bool   `json:"vrouter_specific_pool"`
}

// String returns json representation of the object
func (model *AllocationPoolType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAllocationPoolType makes AllocationPoolType
func MakeAllocationPoolType() *AllocationPoolType {
	return &AllocationPoolType{
		//TODO(nati): Apply default
		Start:               "",
		End:                 "",
		VrouterSpecificPool: false,
	}
}

// MakeAllocationPoolTypeSlice() makes a slice of AllocationPoolType
func MakeAllocationPoolTypeSlice() []*AllocationPoolType {
	return []*AllocationPoolType{}
}
