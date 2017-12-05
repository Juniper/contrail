package models

// PermType

import "encoding/json"

type PermType struct {
	Group       string     `json:"group"`
	GroupAccess AccessType `json:"group_access"`
	Owner       string     `json:"owner"`
	OwnerAccess AccessType `json:"owner_access"`
	OtherAccess AccessType `json:"other_access"`
}

func (model *PermType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakePermType() *PermType {
	return &PermType{
		//TODO(nati): Apply default
		Group:       "",
		GroupAccess: MakeAccessType(),
		Owner:       "",
		OwnerAccess: MakeAccessType(),
		OtherAccess: MakeAccessType(),
	}
}

func InterfaceToPermType(iData interface{}) *PermType {
	data := iData.(map[string]interface{})
	return &PermType{
		Owner: data["owner"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Owner","GoType":"string"}
		OwnerAccess: InterfaceToAccessType(data["owner_access"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"}
		OtherAccess: InterfaceToAccessType(data["other_access"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"OtherAccess","GoType":"AccessType"}
		Group: data["group"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Group","GoType":"string"}
		GroupAccess: InterfaceToAccessType(data["group_access"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"GroupAccess","GoType":"AccessType"}

	}
}

func InterfaceToPermTypeSlice(data interface{}) []*PermType {
	list := data.([]interface{})
	result := MakePermTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPermType(item))
	}
	return result
}

func MakePermTypeSlice() []*PermType {
	return []*PermType{}
}
