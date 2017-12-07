package models

// AllocationPoolType

import "encoding/json"

// AllocationPoolType
type AllocationPoolType struct {
	VrouterSpecificPool bool   `json:"vrouter_specific_pool"`
	Start               string `json:"start"`
	End                 string `json:"end"`
}

//  parents relation object

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

// InterfaceToAllocationPoolType makes AllocationPoolType from interface
func InterfaceToAllocationPoolType(iData interface{}) *AllocationPoolType {
	data := iData.(map[string]interface{})
	return &AllocationPoolType{
		VrouterSpecificPool: data["vrouter_specific_pool"].(bool),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VrouterSpecificPool","GoType":"bool","GoPremitive":true}
		Start: data["start"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Start","GoType":"string","GoPremitive":true}
		End: data["end"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"End","GoType":"string","GoPremitive":true}

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
