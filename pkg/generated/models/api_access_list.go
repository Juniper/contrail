package models

// APIAccessList

import "encoding/json"

// APIAccessList
type APIAccessList struct {
	ParentUUID           string               `json:"parent_uuid,omitempty"`
	DisplayName          string               `json:"display_name,omitempty"`
	IDPerms              *IdPermsType         `json:"id_perms,omitempty"`
	APIAccessListEntries *RbacRuleEntriesType `json:"api_access_list_entries,omitempty"`
	Annotations          *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2               *PermType2           `json:"perms2,omitempty"`
	UUID                 string               `json:"uuid,omitempty"`
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
		DisplayName:          "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		APIAccessListEntries: MakeRbacRuleEntriesType(),
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		UUID:                 "",
	}
}

// MakeAPIAccessListSlice() makes a slice of APIAccessList
func MakeAPIAccessListSlice() []*APIAccessList {
	return []*APIAccessList{}
}
