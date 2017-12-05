package models

// ShareType

import "encoding/json"

type ShareType struct {
	TenantAccess AccessType `json:"tenant_access"`
	Tenant       string     `json:"tenant"`
}

func (model *ShareType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeShareType() *ShareType {
	return &ShareType{
		//TODO(nati): Apply default
		TenantAccess: MakeAccessType(),
		Tenant:       "",
	}
}

func InterfaceToShareType(iData interface{}) *ShareType {
	data := iData.(map[string]interface{})
	return &ShareType{
		TenantAccess: InterfaceToAccessType(data["tenant_access"]),

		//{"Title":"","Description":"Allowed permissions in sharing","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType"}
		Tenant: data["tenant"].(string),

		//{"Title":"","Description":"Name of tenant with whom the object is shared","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string"}

	}
}

func InterfaceToShareTypeSlice(data interface{}) []*ShareType {
	list := data.([]interface{})
	result := MakeShareTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToShareType(item))
	}
	return result
}

func MakeShareTypeSlice() []*ShareType {
	return []*ShareType{}
}
