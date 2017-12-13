package models

// PermType2

import "encoding/json"

// PermType2
type PermType2 struct {
	OwnerAccess  AccessType   `json:"owner_access"`
	GlobalAccess AccessType   `json:"global_access"`
	Share        []*ShareType `json:"share"`
	Owner        string       `json:"owner"`
}

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

		//{"description":"Owner tenant of the object","type":"string"}
		OwnerAccess: InterfaceToAccessType(data["owner_access"]),

		//{"description":"Owner permissions of the object","type":"integer","minimum":0,"maximum":7}
		GlobalAccess: InterfaceToAccessType(data["global_access"]),

		//{"description":"Globally(others) shared object and permissions for others of the object","type":"integer","minimum":0,"maximum":7}

		Share: InterfaceToShareTypeSlice(data["share"]),

		//{"description":"Selectively shared object, List of (tenant, permissions)","type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}

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
