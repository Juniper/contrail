package models

// SequenceType

import "encoding/json"

type SequenceType struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
}

func (model *SequenceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeSequenceType() *SequenceType {
	return &SequenceType{
		//TODO(nati): Apply default
		Major: 0,
		Minor: 0,
	}
}

func InterfaceToSequenceType(iData interface{}) *SequenceType {
	data := iData.(map[string]interface{})
	return &SequenceType{
		Major: data["major"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Major","GoType":"int"}
		Minor: data["minor"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Minor","GoType":"int"}

	}
}

func InterfaceToSequenceTypeSlice(data interface{}) []*SequenceType {
	list := data.([]interface{})
	result := MakeSequenceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSequenceType(item))
	}
	return result
}

func MakeSequenceTypeSlice() []*SequenceType {
	return []*SequenceType{}
}
