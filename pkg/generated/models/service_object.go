package models

// ServiceObject

import "encoding/json"

// ServiceObject
type ServiceObject struct {
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
}

// String returns json representation of the object
func (model *ServiceObject) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceObject makes ServiceObject
func MakeServiceObject() *ServiceObject {
	return &ServiceObject{
		//TODO(nati): Apply default
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
	}
}

// MakeServiceObjectSlice() makes a slice of ServiceObject
func MakeServiceObjectSlice() []*ServiceObject {
	return []*ServiceObject{}
}
