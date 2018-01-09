package models

// SequenceType

import "encoding/json"

// SequenceType
type SequenceType struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
}

// String returns json representation of the object
func (model *SequenceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSequenceType makes SequenceType
func MakeSequenceType() *SequenceType {
	return &SequenceType{
		//TODO(nati): Apply default
		Minor: 0,
		Major: 0,
	}
}

// InterfaceToSequenceType makes SequenceType from interface
func InterfaceToSequenceType(iData interface{}) *SequenceType {
	data := iData.(map[string]interface{})
	return &SequenceType{
		Major: data["major"].(int),

		//{"type":"integer"}
		Minor: data["minor"].(int),

		//{"type":"integer"}

	}
}

// InterfaceToSequenceTypeSlice makes a slice of SequenceType from interface
func InterfaceToSequenceTypeSlice(data interface{}) []*SequenceType {
	list := data.([]interface{})
	result := MakeSequenceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSequenceType(item))
	}
	return result
}

// MakeSequenceTypeSlice() makes a slice of SequenceType
func MakeSequenceTypeSlice() []*SequenceType {
	return []*SequenceType{}
}
