package models

// PermType2

import "encoding/json"

// PermType2
type PermType2 struct {
	GlobalAccess AccessType   `json:"global_access"`
	Share        []*ShareType `json:"share"`
	Owner        string       `json:"owner"`
	OwnerAccess  AccessType   `json:"owner_access"`
}

//  parents relation object

// String returns json representation of the object
func (model *PermType2) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePermType2 makes PermType2
func MakePermType2() *PermType2 {
	return &PermType2{
		//TODO(nati): Apply default
		Owner:        "",
		OwnerAccess:  MakeAccessType(),
		GlobalAccess: MakeAccessType(),

		Share: MakeShareTypeSlice(),
	}
}

// InterfaceToPermType2 makes PermType2 from interface
func InterfaceToPermType2(iData interface{}) *PermType2 {
	data := iData.(map[string]interface{})
	return &PermType2{
		Owner: data["owner"].(string),

		//{"Title":"","Description":"Owner tenant of the object","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true}
		OwnerAccess: InterfaceToAccessType(data["owner_access"]),

		//{"Title":"","Description":"Owner permissions of the object","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}
		GlobalAccess: InterfaceToAccessType(data["global_access"]),

		//{"Title":"","Description":"Globally(others) shared object and permissions for others of the object","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"GlobalAccess","GoType":"AccessType","GoPremitive":false}

		Share: InterfaceToShareTypeSlice(data["share"]),

		//{"Title":"","Description":"Selectively shared object, List of (tenant, permissions)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string","GoPremitive":true},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType","GoPremitive":false},"GoName":"Share","GoType":"[]*ShareType","GoPremitive":true}

	}
}

// InterfaceToPermType2Slice makes a slice of PermType2 from interface
func InterfaceToPermType2Slice(data interface{}) []*PermType2 {
	list := data.([]interface{})
	result := MakePermType2Slice()
	for _, item := range list {
		result = append(result, InterfaceToPermType2(item))
	}
	return result
}

// MakePermType2Slice() makes a slice of PermType2
func MakePermType2Slice() []*PermType2 {
	return []*PermType2{}
}
