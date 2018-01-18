package models

// SecurityGroup

import "encoding/json"

// SecurityGroup
type SecurityGroup struct {
	Annotations               *KeyValuePairs                `json:"annotations,omitempty"`
	ParentType                string                        `json:"parent_type,omitempty"`
	FQName                    []string                      `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType                  `json:"id_perms,omitempty"`
	DisplayName               string                        `json:"display_name,omitempty"`
	SecurityGroupEntries      *PolicyEntriesType            `json:"security_group_entries,omitempty"`
	ConfiguredSecurityGroupID ConfiguredSecurityGroupIdType `json:"configured_security_group_id,omitempty"`
	SecurityGroupID           SecurityGroupIdType           `json:"security_group_id,omitempty"`
	ParentUUID                string                        `json:"parent_uuid,omitempty"`
	Perms2                    *PermType2                    `json:"perms2,omitempty"`
	UUID                      string                        `json:"uuid,omitempty"`

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
		Annotations:               MakeKeyValuePairs(),
		SecurityGroupID:           MakeSecurityGroupIdType(),
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		SecurityGroupEntries:      MakePolicyEntriesType(),
		ConfiguredSecurityGroupID: MakeConfiguredSecurityGroupIdType(),
		Perms2: MakePermType2(),
		UUID:   "",
	}
}

// MakeSecurityGroupSlice() makes a slice of SecurityGroup
func MakeSecurityGroupSlice() []*SecurityGroup {
	return []*SecurityGroup{}
}
