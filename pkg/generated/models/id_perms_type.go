package models

// IdPermsType

import "encoding/json"

// IdPermsType
type IdPermsType struct {
	Description  string    `json:"description"`
	Created      string    `json:"created"`
	Creator      string    `json:"creator"`
	UserVisible  bool      `json:"user_visible"`
	LastModified string    `json:"last_modified"`
	Permissions  *PermType `json:"permissions"`
	Enable       bool      `json:"enable"`
}

//  parents relation object

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
		Permissions: InterfaceToPermType(data["permissions"]),

		//{"Title":"","Description":"No longer used, will be removed","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Group","GoType":"string","GoPremitive":true},"group_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"GroupAccess","GoType":"AccessType","GoPremitive":false},"other_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"OtherAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType","GoPremitive":false}
		Enable: data["enable"].(bool),

		//{"Title":"","Description":"Administratively Enable/Disable this object","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true}
		Description: data["description"].(string),

		//{"Title":"","Description":"User provided text","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Description","GoType":"string","GoPremitive":true}
		Created: data["created"].(string),

		//{"Title":"","Description":"Time when this object was created","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Created","GoType":"string","GoPremitive":true}
		Creator: data["creator"].(string),

		//{"Title":"","Description":"Id of tenant who created this object","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Creator","GoType":"string","GoPremitive":true}
		UserVisible: data["user_visible"].(bool),

		//{"Title":"","Description":"System created internal objects will have this flag set and will not be visible","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UserVisible","GoType":"bool","GoPremitive":true}
		LastModified: data["last_modified"].(string),

		//{"Title":"","Description":"Time when this object was created","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LastModified","GoType":"string","GoPremitive":true}

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
