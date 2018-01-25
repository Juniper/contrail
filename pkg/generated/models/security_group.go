package models

// SecurityGroup

// SecurityGroup
//proteus:generate
type SecurityGroup struct {
	UUID                      string                        `json:"uuid,omitempty"`
	ParentUUID                string                        `json:"parent_uuid,omitempty"`
	ParentType                string                        `json:"parent_type,omitempty"`
	FQName                    []string                      `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType                  `json:"id_perms,omitempty"`
	DisplayName               string                        `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs                `json:"annotations,omitempty"`
	Perms2                    *PermType2                    `json:"perms2,omitempty"`
	SecurityGroupEntries      *PolicyEntriesType            `json:"security_group_entries,omitempty"`
	ConfiguredSecurityGroupID ConfiguredSecurityGroupIdType `json:"configured_security_group_id,omitempty"`
	SecurityGroupID           SecurityGroupIdType           `json:"security_group_id,omitempty"`

	AccessControlLists []*AccessControlList `json:"access_control_lists,omitempty"`
}

// MakeSecurityGroup makes SecurityGroup
func MakeSecurityGroup() *SecurityGroup {
	return &SecurityGroup{
		//TODO(nati): Apply default
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		SecurityGroupEntries:      MakePolicyEntriesType(),
		ConfiguredSecurityGroupID: MakeConfiguredSecurityGroupIdType(),
		SecurityGroupID:           MakeSecurityGroupIdType(),
	}
}

// MakeSecurityGroupSlice() makes a slice of SecurityGroup
func MakeSecurityGroupSlice() []*SecurityGroup {
	return []*SecurityGroup{}
}
