package models

// SecurityGroup

import "encoding/json"

// SecurityGroup
type SecurityGroup struct {
	IDPerms                   *IdPermsType                  `json:"id_perms,omitempty"`
	DisplayName               string                        `json:"display_name,omitempty"`
	ConfiguredSecurityGroupID ConfiguredSecurityGroupIdType `json:"configured_security_group_id,omitempty"`
	Perms2                    *PermType2                    `json:"perms2,omitempty"`
	FQName                    []string                      `json:"fq_name,omitempty"`
	UUID                      string                        `json:"uuid,omitempty"`
	ParentUUID                string                        `json:"parent_uuid,omitempty"`
	ParentType                string                        `json:"parent_type,omitempty"`
	SecurityGroupEntries      *PolicyEntriesType            `json:"security_group_entries,omitempty"`
	SecurityGroupID           SecurityGroupIdType           `json:"security_group_id,omitempty"`
	Annotations               *KeyValuePairs                `json:"annotations,omitempty"`

	AccessControlLists []*AccessControlList `json:"access_control_lists,omitempty"`
}

// String returns json representation of the object
func (model *SecurityGroup) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSecurityGroup makes SecurityGroup
func MakeSecurityGroup() *SecurityGroup {
	return &SecurityGroup{
		//TODO(nati): Apply default
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		ConfiguredSecurityGroupID: MakeConfiguredSecurityGroupIdType(),
		Perms2:               MakePermType2(),
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		SecurityGroupEntries: MakePolicyEntriesType(),
		SecurityGroupID:      MakeSecurityGroupIdType(),
	}
}

// MakeSecurityGroupSlice() makes a slice of SecurityGroup
func MakeSecurityGroupSlice() []*SecurityGroup {
	return []*SecurityGroup{}
}
