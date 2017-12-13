package models

// PermType

import "encoding/json"

// PermType
type PermType struct {
	OwnerAccess AccessType `json:"owner_access"`
	OtherAccess AccessType `json:"other_access"`
	Group       string     `json:"group"`
	GroupAccess AccessType `json:"group_access"`
	Owner       string     `json:"owner"`
}

// String returns json representation of the object
func (model *PermType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePermType makes PermType
func MakePermType() *PermType {
	return &PermType{
		//TODO(nati): Apply default
		Owner:       "",
		OwnerAccess: MakeAccessType(),
		OtherAccess: MakeAccessType(),
		Group:       "",
		GroupAccess: MakeAccessType(),
	}
}

// InterfaceToPermType makes PermType from interface
func InterfaceToPermType(iData interface{}) *PermType {
	data := iData.(map[string]interface{})
	return &PermType{
		Owner: data["owner"].(string),

		//{"type":"string"}
		OwnerAccess: InterfaceToAccessType(data["owner_access"]),

		//{"type":"integer","minimum":0,"maximum":7}
		OtherAccess: InterfaceToAccessType(data["other_access"]),

		//{"type":"integer","minimum":0,"maximum":7}
		Group: data["group"].(string),

		//{"type":"string"}
		GroupAccess: InterfaceToAccessType(data["group_access"]),

		//{"type":"integer","minimum":0,"maximum":7}

	}
}

// InterfaceToPermTypeSlice makes a slice of PermType from interface
func InterfaceToPermTypeSlice(data interface{}) []*PermType {
	list := data.([]interface{})
	result := MakePermTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPermType(item))
	}
	return result
}

// MakePermTypeSlice() makes a slice of PermType
func MakePermTypeSlice() []*PermType {
	return []*PermType{}
}
