package models

// MemberType

import "encoding/json"

// MemberType
type MemberType struct {
	Role string `json:"role"`
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

// InterfaceToMemberType makes MemberType from interface
func InterfaceToMemberType(iData interface{}) *MemberType {
	data := iData.(map[string]interface{})
	return &MemberType{
		Role: data["role"].(string),

		//{"description":"User role for the project","type":"string"}

	}
}

// InterfaceToMemberTypeSlice makes a slice of MemberType from interface
func InterfaceToMemberTypeSlice(data interface{}) []*MemberType {
	list := data.([]interface{})
	result := MakeMemberTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMemberType(item))
	}
	return result
}

// MakeMemberTypeSlice() makes a slice of MemberType
func MakeMemberTypeSlice() []*MemberType {
	return []*MemberType{}
}
