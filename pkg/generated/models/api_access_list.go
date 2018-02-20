package models

// APIAccessList

// APIAccessList
//proteus:generate
type APIAccessList struct {
	UUID                 string               `json:"uuid,omitempty"`
	ParentUUID           string               `json:"parent_uuid,omitempty"`
	ParentType           string               `json:"parent_type,omitempty"`
	FQName               []string             `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType         `json:"id_perms,omitempty"`
	DisplayName          string               `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2               *PermType2           `json:"perms2,omitempty"`
	APIAccessListEntries *RbacRuleEntriesType `json:"api_access_list_entries,omitempty"`
}

// MakeAPIAccessList makes APIAccessList
func MakeAPIAccessList() *APIAccessList {
	return &APIAccessList{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		APIAccessListEntries: MakeRbacRuleEntriesType(),
	}
}

// MakeAPIAccessListSlice() makes a slice of APIAccessList
func MakeAPIAccessListSlice() []*APIAccessList {
	return []*APIAccessList{}
}
