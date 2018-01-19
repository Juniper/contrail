package models

// SecurityGroup

import "encoding/json"

// SecurityGroup
type SecurityGroup struct {
	UUID                      string                        `json:"uuid,omitempty"`
	FQName                    []string                      `json:"fq_name,omitempty"`
	DisplayName               string                        `json:"display_name,omitempty"`
	ConfiguredSecurityGroupID ConfiguredSecurityGroupIdType `json:"configured_security_group_id,omitempty"`
	SecurityGroupID           SecurityGroupIdType           `json:"security_group_id,omitempty"`
	Annotations               *KeyValuePairs                `json:"annotations,omitempty"`
	Perms2                    *PermType2                    `json:"perms2,omitempty"`
	SecurityGroupEntries      *PolicyEntriesType            `json:"security_group_entries,omitempty"`
	ParentUUID                string                        `json:"parent_uuid,omitempty"`
	ParentType                string                        `json:"parent_type,omitempty"`
	IDPerms                   *IdPermsType                  `json:"id_perms,omitempty"`

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
		SecurityGroupEntries:      MakePolicyEntriesType(),
		ParentUUID:                "",
		ParentType:                "",
		IDPerms:                   MakeIdPermsType(),
		UUID:                      "",
		FQName:                    []string{},
		DisplayName:               "",
		ConfiguredSecurityGroupID: MakeConfiguredSecurityGroupIdType(),
		SecurityGroupID:           MakeSecurityGroupIdType(),
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
	}
}

// MakeSecurityGroupSlice() makes a slice of SecurityGroup
func MakeSecurityGroupSlice() []*SecurityGroup {
	return []*SecurityGroup{}
}
