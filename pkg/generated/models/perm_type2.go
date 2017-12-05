package models

// PermType2

import "encoding/json"

type PermType2 struct {
	Owner        string       `json:"owner"`
	OwnerAccess  AccessType   `json:"owner_access"`
	GlobalAccess AccessType   `json:"global_access"`
	Share        []*ShareType `json:"share"`
}

func (model *PermType2) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakePermType2() *PermType2 {
	return &PermType2{
		//TODO(nati): Apply default
		Owner:        "",
		OwnerAccess:  MakeAccessType(),
		GlobalAccess: MakeAccessType(),

		Share: MakeShareTypeSlice(),
	}
}

func InterfaceToPermType2(iData interface{}) *PermType2 {
	data := iData.(map[string]interface{})
	return &PermType2{
		Owner: data["owner"].(string),

		//{"Title":"","Description":"Owner tenant of the object","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Owner","GoType":"string"}
		OwnerAccess: InterfaceToAccessType(data["owner_access"]),

		//{"Title":"","Description":"Owner permissions of the object","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"}
		GlobalAccess: InterfaceToAccessType(data["global_access"]),

		//{"Title":"","Description":"Globally(others) shared object and permissions for others of the object","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"GlobalAccess","GoType":"AccessType"}

		Share: InterfaceToShareTypeSlice(data["share"]),

		//{"Title":"","Description":"Selectively shared object, List of (tenant, permissions)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string"},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType"},"GoName":"Share","GoType":"[]*ShareType"}

	}
}

func InterfaceToPermType2Slice(data interface{}) []*PermType2 {
	list := data.([]interface{})
	result := MakePermType2Slice()
	for _, item := range list {
		result = append(result, InterfaceToPermType2(item))
	}
	return result
}

func MakePermType2Slice() []*PermType2 {
	return []*PermType2{}
}
