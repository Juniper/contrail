package models

// IdPermsType

import "encoding/json"

type IdPermsType struct {
	Created      string    `json:"created"`
	Creator      string    `json:"creator"`
	UserVisible  bool      `json:"user_visible"`
	LastModified string    `json:"last_modified"`
	Permissions  *PermType `json:"permissions"`
	Enable       bool      `json:"enable"`
	Description  string    `json:"description"`
}

func (model *IdPermsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

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

func InterfaceToIdPermsType(iData interface{}) *IdPermsType {
	data := iData.(map[string]interface{})
	return &IdPermsType{
		Creator: data["creator"].(string),

		//{"Title":"","Description":"Id of tenant who created this object","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Creator","GoType":"string"}
		UserVisible: data["user_visible"].(bool),

		//{"Title":"","Description":"System created internal objects will have this flag set and will not be visible","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UserVisible","GoType":"bool"}
		LastModified: data["last_modified"].(string),

		//{"Title":"","Description":"Time when this object was created","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LastModified","GoType":"string"}
		Permissions: InterfaceToPermType(data["permissions"]),

		//{"Title":"","Description":"No longer used, will be removed","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Group","GoType":"string"},"group_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"GroupAccess","GoType":"AccessType"},"other_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"OtherAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType"}
		Enable: data["enable"].(bool),

		//{"Title":"","Description":"Administratively Enable/Disable this object","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Enable","GoType":"bool"}
		Description: data["description"].(string),

		//{"Title":"","Description":"User provided text","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Description","GoType":"string"}
		Created: data["created"].(string),

		//{"Title":"","Description":"Time when this object was created","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Created","GoType":"string"}

	}
}

func InterfaceToIdPermsTypeSlice(data interface{}) []*IdPermsType {
	list := data.([]interface{})
	result := MakeIdPermsTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIdPermsType(item))
	}
	return result
}

func MakeIdPermsTypeSlice() []*IdPermsType {
	return []*IdPermsType{}
}
