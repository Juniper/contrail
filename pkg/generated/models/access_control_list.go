package models

// AccessControlList

// AccessControlList
//proteus:generate
type AccessControlList struct {
	UUID                     string          `json:"uuid,omitempty"`
	ParentUUID               string          `json:"parent_uuid,omitempty"`
	ParentType               string          `json:"parent_type,omitempty"`
	FQName                   []string        `json:"fq_name,omitempty"`
	IDPerms                  *IdPermsType    `json:"id_perms,omitempty"`
	DisplayName              string          `json:"display_name,omitempty"`
	Annotations              *KeyValuePairs  `json:"annotations,omitempty"`
	Perms2                   *PermType2      `json:"perms2,omitempty"`
	AccessControlListHash    int             `json:"access_control_list_hash,omitempty"`
	AccessControlListEntries *AclEntriesType `json:"access_control_list_entries,omitempty"`
}

// MakeAccessControlList makes AccessControlList
func MakeAccessControlList() *AccessControlList {
	return &AccessControlList{
		//TODO(nati): Apply default
		UUID:                     "",
		ParentUUID:               "",
		ParentType:               "",
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		AccessControlListHash:    0,
		AccessControlListEntries: MakeAclEntriesType(),
	}
}

// MakeAccessControlListSlice() makes a slice of AccessControlList
func MakeAccessControlListSlice() []*AccessControlList {
	return []*AccessControlList{}
}
