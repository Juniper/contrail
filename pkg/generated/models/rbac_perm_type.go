package models

// RbacPermType

import "encoding/json"

// RbacPermType
type RbacPermType struct {
	RoleCrud string `json:"role_crud"`
	RoleName string `json:"role_name"`
}

// String returns json representation of the object
func (model *RbacPermType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRbacPermType makes RbacPermType
func MakeRbacPermType() *RbacPermType {
	return &RbacPermType{
		//TODO(nati): Apply default
		RoleCrud: "",
		RoleName: "",
	}
}

// MakeRbacPermTypeSlice() makes a slice of RbacPermType
func MakeRbacPermTypeSlice() []*RbacPermType {
	return []*RbacPermType{}
}
