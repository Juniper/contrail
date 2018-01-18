package models

// MemberType

import "encoding/json"

// MemberType
type MemberType struct {
	Role string `json:"role,omitempty"`
}

// String returns json representation of the object
func (model *MemberType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeMemberType makes MemberType
func MakeMemberType() *MemberType {
	return &MemberType{
		//TODO(nati): Apply default
		Role: "",
	}
}

// MakeMemberTypeSlice() makes a slice of MemberType
func MakeMemberTypeSlice() []*MemberType {
	return []*MemberType{}
}
