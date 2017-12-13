package models

// ShareType

import "encoding/json"

// ShareType
type ShareType struct {
	TenantAccess AccessType `json:"tenant_access"`
	Tenant       string     `json:"tenant"`
}

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

		//{"description":"Allowed permissions in sharing","type":"integer","minimum":0,"maximum":7}
		Tenant: data["tenant"].(string),

		//{"description":"Name of tenant with whom the object is shared","type":"string"}

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
