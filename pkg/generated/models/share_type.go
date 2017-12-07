package models

// ShareType

import "encoding/json"

// ShareType
type ShareType struct {
	TenantAccess AccessType `json:"tenant_access"`
	Tenant       string     `json:"tenant"`
}

//  parents relation object

// String returns json representation of the object
func (model *ShareType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeShareType makes ShareType
func MakeShareType() *ShareType {
	return &ShareType{
		//TODO(nati): Apply default
		TenantAccess: MakeAccessType(),
		Tenant:       "",
	}
}

// InterfaceToShareType makes ShareType from interface
func InterfaceToShareType(iData interface{}) *ShareType {
	data := iData.(map[string]interface{})
	return &ShareType{
		TenantAccess: InterfaceToAccessType(data["tenant_access"]),

		//{"Title":"","Description":"Allowed permissions in sharing","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType","GoPremitive":false}
		Tenant: data["tenant"].(string),

		//{"Title":"","Description":"Name of tenant with whom the object is shared","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToShareTypeSlice makes a slice of ShareType from interface
func InterfaceToShareTypeSlice(data interface{}) []*ShareType {
	list := data.([]interface{})
	result := MakeShareTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToShareType(item))
	}
	return result
}

// MakeShareTypeSlice() makes a slice of ShareType
func MakeShareTypeSlice() []*ShareType {
	return []*ShareType{}
}
