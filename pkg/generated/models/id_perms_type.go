package models

// IdPermsType

import "encoding/json"

// IdPermsType
type IdPermsType struct {
	Creator      string    `json:"creator"`
	UserVisible  bool      `json:"user_visible"`
	LastModified string    `json:"last_modified"`
	Permissions  *PermType `json:"permissions"`
	Enable       bool      `json:"enable"`
	Description  string    `json:"description"`
	Created      string    `json:"created"`
}

// String returns json representation of the object
func (model *IdPermsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIdPermsType makes IdPermsType
func MakeIdPermsType() *IdPermsType {
	return &IdPermsType{
		//TODO(nati): Apply default
		Enable:       false,
		Description:  "",
		Created:      "",
		Creator:      "",
		UserVisible:  false,
		LastModified: "",
		Permissions:  MakePermType(),
	}
}

// InterfaceToIdPermsType makes IdPermsType from interface
func InterfaceToIdPermsType(iData interface{}) *IdPermsType {
	data := iData.(map[string]interface{})
	return &IdPermsType{
		LastModified: data["last_modified"].(string),

		//{"description":"Time when this object was created","type":"string"}
		Permissions: InterfaceToPermType(data["permissions"]),

		//{"description":"No longer used, will be removed","type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}}
		Enable: data["enable"].(bool),

		//{"description":"Administratively Enable/Disable this object","type":"boolean"}
		Description: data["description"].(string),

		//{"description":"User provided text","type":"string"}
		Created: data["created"].(string),

		//{"description":"Time when this object was created","type":"string"}
		Creator: data["creator"].(string),

		//{"description":"Id of tenant who created this object","type":"string"}
		UserVisible: data["user_visible"].(bool),

		//{"description":"System created internal objects will have this flag set and will not be visible","type":"boolean"}

	}
}

// InterfaceToIdPermsTypeSlice makes a slice of IdPermsType from interface
func InterfaceToIdPermsTypeSlice(data interface{}) []*IdPermsType {
	list := data.([]interface{})
	result := MakeIdPermsTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIdPermsType(item))
	}
	return result
}

// MakeIdPermsTypeSlice() makes a slice of IdPermsType
func MakeIdPermsTypeSlice() []*IdPermsType {
	return []*IdPermsType{}
}
