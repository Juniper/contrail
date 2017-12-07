package models

// PermType

import "encoding/json"

// PermType
type PermType struct {
	OtherAccess AccessType `json:"other_access"`
	Group       string     `json:"group"`
	GroupAccess AccessType `json:"group_access"`
	Owner       string     `json:"owner"`
	OwnerAccess AccessType `json:"owner_access"`
}

//  parents relation object

// String returns json representation of the object
func (model *PermType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePermType makes PermType
func MakePermType() *PermType {
	return &PermType{
		//TODO(nati): Apply default
		GroupAccess: MakeAccessType(),
		Owner:       "",
		OwnerAccess: MakeAccessType(),
		OtherAccess: MakeAccessType(),
		Group:       "",
	}
}

// InterfaceToPermType makes PermType from interface
func InterfaceToPermType(iData interface{}) *PermType {
	data := iData.(map[string]interface{})
	return &PermType{
		Owner: data["owner"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true}
		OwnerAccess: InterfaceToAccessType(data["owner_access"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}
		OtherAccess: InterfaceToAccessType(data["other_access"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"OtherAccess","GoType":"AccessType","GoPremitive":false}
		Group: data["group"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Group","GoType":"string","GoPremitive":true}
		GroupAccess: InterfaceToAccessType(data["group_access"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"GroupAccess","GoType":"AccessType","GoPremitive":false}

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
