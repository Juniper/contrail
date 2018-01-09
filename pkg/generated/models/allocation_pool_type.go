package models

// AllocationPoolType

import "encoding/json"

// AllocationPoolType
type AllocationPoolType struct {
	Start               string `json:"start"`
	End                 string `json:"end"`
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
		VrouterSpecificPool: false,
		Start:               "",
		End:                 "",
	}
}

// InterfaceToAllocationPoolType makes AllocationPoolType from interface
func InterfaceToAllocationPoolType(iData interface{}) *AllocationPoolType {
	data := iData.(map[string]interface{})
	return &AllocationPoolType{
		VrouterSpecificPool: data["vrouter_specific_pool"].(bool),

		//{"type":"boolean"}
		Start: data["start"].(string),

		//{"type":"string"}
		End: data["end"].(string),

		//{"type":"string"}

	}
}

// InterfaceToAllocationPoolTypeSlice makes a slice of AllocationPoolType from interface
func InterfaceToAllocationPoolTypeSlice(data interface{}) []*AllocationPoolType {
	list := data.([]interface{})
	result := MakeAllocationPoolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAllocationPoolType(item))
	}
	return result
}

// MakeAllocationPoolTypeSlice() makes a slice of AllocationPoolType
func MakeAllocationPoolTypeSlice() []*AllocationPoolType {
	return []*AllocationPoolType{}
}
