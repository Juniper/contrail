package models

// APIAccessList

import "encoding/json"

// APIAccessList
type APIAccessList struct {
	APIAccessListEntries *RbacRuleEntriesType `json:"api_access_list_entries,omitempty"`
	IDPerms              *IdPermsType         `json:"id_perms,omitempty"`
	Perms2               *PermType2           `json:"perms2,omitempty"`
	DisplayName          string               `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs       `json:"annotations,omitempty"`
	UUID                 string               `json:"uuid,omitempty"`
	ParentUUID           string               `json:"parent_uuid,omitempty"`
	ParentType           string               `json:"parent_type,omitempty"`
	FQName               []string             `json:"fq_name,omitempty"`
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
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
		APIAccessListEntries: MakeRbacRuleEntriesType(),
		IDPerms:              MakeIdPermsType(),
		Perms2:               MakePermType2(),
	}
}

// MakeAPIAccessListSlice() makes a slice of APIAccessList
func MakeAPIAccessListSlice() []*APIAccessList {
	return []*APIAccessList{}
}
