package models

// AccessControlList

import "encoding/json"

// AccessControlList
type AccessControlList struct {
	IDPerms                  *IdPermsType           `json:"id_perms"`
	DisplayName              string                 `json:"display_name"`
	AccessControlListHash    map[string]interface{} `json:"access_control_list_hash"`
	AccessControlListEntries *AclEntriesType        `json:"access_control_list_entries"`
	Annotations              *KeyValuePairs         `json:"annotations"`
	Perms2                   *PermType2             `json:"perms2"`
	ParentUUID               string                 `json:"parent_uuid"`
	UUID                     string                 `json:"uuid"`
	ParentType               string                 `json:"parent_type"`
	FQName                   []string               `json:"fq_name"`
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
		Perms2:                   MakePermType2(),
		ParentUUID:               "",
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		AccessControlListHash:    map[string]interface{}{},
		AccessControlListEntries: MakeAclEntriesType(),
		Annotations:              MakeKeyValuePairs(),
		UUID:                     "",
		ParentType:               "",
		FQName:                   []string{},
	}
}

// MakeAccessControlListSlice() makes a slice of AccessControlList
func MakeAccessControlListSlice() []*AccessControlList {
	return []*AccessControlList{}
}
