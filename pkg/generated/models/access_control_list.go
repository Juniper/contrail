package models

// AccessControlList

import "encoding/json"

// AccessControlList
type AccessControlList struct {
	AccessControlListEntries *AclEntriesType        `json:"access_control_list_entries,omitempty"`
	UUID                     string                 `json:"uuid,omitempty"`
	ParentUUID               string                 `json:"parent_uuid,omitempty"`
	FQName                   []string               `json:"fq_name,omitempty"`
	DisplayName              string                 `json:"display_name,omitempty"`
	AccessControlListHash    map[string]interface{} `json:"access_control_list_hash,omitempty"`
	Perms2                   *PermType2             `json:"perms2,omitempty"`
	ParentType               string                 `json:"parent_type,omitempty"`
	IDPerms                  *IdPermsType           `json:"id_perms,omitempty"`
	Annotations              *KeyValuePairs         `json:"annotations,omitempty"`
}

// String returns json representation of the object
func (model *AccessControlList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAccessControlList makes AccessControlList
func MakeAccessControlList() *AccessControlList {
	return &AccessControlList{
		//TODO(nati): Apply default
		Annotations:              MakeKeyValuePairs(),
		AccessControlListHash:    map[string]interface{}{},
		Perms2:                   MakePermType2(),
		ParentType:               "",
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		AccessControlListEntries: MakeAclEntriesType(),
		UUID:       "",
		ParentUUID: "",
		FQName:     []string{},
	}
}

// MakeAccessControlListSlice() makes a slice of AccessControlList
func MakeAccessControlListSlice() []*AccessControlList {
	return []*AccessControlList{}
}
