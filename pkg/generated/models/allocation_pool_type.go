package models

// AllocationPoolType

import "encoding/json"

// AllocationPoolType
type AllocationPoolType struct {
	VrouterSpecificPool bool   `json:"vrouter_specific_pool,omitempty"`
	Start               string `json:"start,omitempty"`
	End                 string `json:"end,omitempty"`
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
		VrouterSpecificPool: false,
		Start:               "",
		End:                 "",
	}
}

// MakeAllocationPoolTypeSlice() makes a slice of AllocationPoolType
func MakeAllocationPoolTypeSlice() []*AllocationPoolType {
	return []*AllocationPoolType{}
}
