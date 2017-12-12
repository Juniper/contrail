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

// InterfaceToRbacPermType makes RbacPermType from interface
func InterfaceToRbacPermType(iData interface{}) *RbacPermType {
	data := iData.(map[string]interface{})
	return &RbacPermType{
		RoleCrud: data["role_crud"].(string),

		//{"description":"String CRUD representing permissions for C=create, R=read, U=update, D=delete.","type":"string"}
		RoleName: data["role_name"].(string),

		//{"description":"Name of the role","type":"string"}

	}
}

// InterfaceToRbacPermTypeSlice makes a slice of RbacPermType from interface
func InterfaceToRbacPermTypeSlice(data interface{}) []*RbacPermType {
	list := data.([]interface{})
	result := MakeRbacPermTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRbacPermType(item))
	}
	return result
}

// MakeRbacPermTypeSlice() makes a slice of RbacPermType
func MakeRbacPermTypeSlice() []*RbacPermType {
	return []*RbacPermType{}
}
