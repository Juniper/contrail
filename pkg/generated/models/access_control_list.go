package models

// AccessControlList

import "encoding/json"

// AccessControlList
type AccessControlList struct {
	AccessControlListHash    map[string]interface{} `json:"access_control_list_hash,omitempty"`
	AccessControlListEntries *AclEntriesType        `json:"access_control_list_entries,omitempty"`
	ParentType               string                 `json:"parent_type,omitempty"`
	FQName                   []string               `json:"fq_name,omitempty"`
	ParentUUID               string                 `json:"parent_uuid,omitempty"`
	IDPerms                  *IdPermsType           `json:"id_perms,omitempty"`
	DisplayName              string                 `json:"display_name,omitempty"`
	Annotations              *KeyValuePairs         `json:"annotations,omitempty"`
	Perms2                   *PermType2             `json:"perms2,omitempty"`
	UUID                     string                 `json:"uuid,omitempty"`
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
		UUID:                     "",
		ParentUUID:               "",
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		FQName:                   []string{},
		AccessControlListHash:    map[string]interface{}{},
		AccessControlListEntries: MakeAclEntriesType(),
		ParentType:               "",
	}
}

// MakeAccessControlListSlice() makes a slice of AccessControlList
func MakeAccessControlListSlice() []*AccessControlList {
	return []*AccessControlList{}
}
