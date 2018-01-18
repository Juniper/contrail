package models

// APIAccessList

import "encoding/json"

// APIAccessList
type APIAccessList struct {
	FQName               []string             `json:"fq_name,omitempty"`
	Annotations          *KeyValuePairs       `json:"annotations,omitempty"`
	UUID                 string               `json:"uuid,omitempty"`
	ParentUUID           string               `json:"parent_uuid,omitempty"`
	ParentType           string               `json:"parent_type,omitempty"`
	APIAccessListEntries *RbacRuleEntriesType `json:"api_access_list_entries,omitempty"`
	DisplayName          string               `json:"display_name,omitempty"`
	Perms2               *PermType2           `json:"perms2,omitempty"`
	IDPerms              *IdPermsType         `json:"id_perms,omitempty"`
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
		DisplayName:          "",
		Perms2:               MakePermType2(),
		IDPerms:              MakeIdPermsType(),
		FQName:               []string{},
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		APIAccessListEntries: MakeRbacRuleEntriesType(),
	}
}

// MakeAPIAccessListSlice() makes a slice of APIAccessList
func MakeAPIAccessListSlice() []*APIAccessList {
	return []*APIAccessList{}
}
