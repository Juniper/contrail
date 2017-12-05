package models

// AllocationPoolType

import "encoding/json"

type AllocationPoolType struct {
	VrouterSpecificPool bool   `json:"vrouter_specific_pool"`
	Start               string `json:"start"`
	End                 string `json:"end"`
}

func (model *AllocationPoolType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeAllocationPoolType() *AllocationPoolType {
	return &AllocationPoolType{
		//TODO(nati): Apply default
		VrouterSpecificPool: false,
		Start:               "",
		End:                 "",
	}
}

func InterfaceToAllocationPoolType(iData interface{}) *AllocationPoolType {
	data := iData.(map[string]interface{})
	return &AllocationPoolType{
		VrouterSpecificPool: data["vrouter_specific_pool"].(bool),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VrouterSpecificPool","GoType":"bool"}
		Start: data["start"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Start","GoType":"string"}
		End: data["end"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"End","GoType":"string"}

	}
}

func InterfaceToAllocationPoolTypeSlice(data interface{}) []*AllocationPoolType {
	list := data.([]interface{})
	result := MakeAllocationPoolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAllocationPoolType(item))
	}
	return result
}

func MakeAllocationPoolTypeSlice() []*AllocationPoolType {
	return []*AllocationPoolType{}
}
