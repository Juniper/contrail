package models

// RbacPermType

import "encoding/json"

type RbacPermType struct {
	RoleName string `json:"role_name"`
	RoleCrud string `json:"role_crud"`
}

func (model *RbacPermType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeRbacPermType() *RbacPermType {
	return &RbacPermType{
		//TODO(nati): Apply default
		RoleCrud: "",
		RoleName: "",
	}
}

func InterfaceToRbacPermType(iData interface{}) *RbacPermType {
	data := iData.(map[string]interface{})
	return &RbacPermType{
		RoleCrud: data["role_crud"].(string),

		//{"Title":"","Description":"String CRUD representing permissions for C=create, R=read, U=update, D=delete.","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoleCrud","GoType":"string"}
		RoleName: data["role_name"].(string),

		//{"Title":"","Description":"Name of the role","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoleName","GoType":"string"}

	}
}

func InterfaceToRbacPermTypeSlice(data interface{}) []*RbacPermType {
	list := data.([]interface{})
	result := MakeRbacPermTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRbacPermType(item))
	}
	return result
}

func MakeRbacPermTypeSlice() []*RbacPermType {
	return []*RbacPermType{}
}
