package models

// ServiceObject

import "encoding/json"

// ServiceObject
type ServiceObject struct {
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
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
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
	}
}

// MakeServiceObjectSlice() makes a slice of ServiceObject
func MakeServiceObjectSlice() []*ServiceObject {
	return []*ServiceObject{}
}
