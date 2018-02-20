package models

// SequenceType

// SequenceType
//proteus:generate
type SequenceType struct {
	Major int `json:"major,omitempty"`
	Minor int `json:"minor,omitempty"`
}

// MakeSequenceType makes SequenceType
func MakeSequenceType() *SequenceType {
	return &SequenceType{
		//TODO(nati): Apply default
		Major: 0,
		Minor: 0,
	}
}

// MakeSequenceTypeSlice() makes a slice of SequenceType
func MakeSequenceTypeSlice() []*SequenceType {
	return []*SequenceType{}
}
