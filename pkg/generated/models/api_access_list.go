package models

// APIAccessList

import "encoding/json"

// APIAccessList
type APIAccessList struct {
	ParentUUID           string               `json:"parent_uuid"`
	FQName               []string             `json:"fq_name"`
	IDPerms              *IdPermsType         `json:"id_perms"`
	APIAccessListEntries *RbacRuleEntriesType `json:"api_access_list_entries"`
	DisplayName          string               `json:"display_name"`
	Annotations          *KeyValuePairs       `json:"annotations"`
	Perms2               *PermType2           `json:"perms2"`
	UUID                 string               `json:"uuid"`
	ParentType           string               `json:"parent_type"`
}

// String returns json representation of the object
func (model *APIAccessList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAPIAccessList makes APIAccessList
func MakeAPIAccessList() *APIAccessList {
	return &APIAccessList{
		//TODO(nati): Apply default
		APIAccessListEntries: MakeRbacRuleEntriesType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		UUID:                 "",
		ParentType:           "",
		ParentUUID:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
	}
}

// MakeAPIAccessListSlice() makes a slice of APIAccessList
func MakeAPIAccessListSlice() []*APIAccessList {
	return []*APIAccessList{}
}
