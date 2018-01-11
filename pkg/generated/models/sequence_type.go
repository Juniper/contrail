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

// MakeSequenceTypeSlice() makes a slice of SequenceType
func MakeSequenceTypeSlice() []*SequenceType {
	return []*SequenceType{}
}
