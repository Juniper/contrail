package models

// SecurityGroup

import "encoding/json"

// SecurityGroup
type SecurityGroup struct {
	SecurityGroupID           SecurityGroupIdType           `json:"security_group_id"`
	UUID                      string                        `json:"uuid"`
	ParentUUID                string                        `json:"parent_uuid"`
	ParentType                string                        `json:"parent_type"`
	IDPerms                   *IdPermsType                  `json:"id_perms"`
	DisplayName               string                        `json:"display_name"`
	Annotations               *KeyValuePairs                `json:"annotations"`
	Perms2                    *PermType2                    `json:"perms2"`
	SecurityGroupEntries      *PolicyEntriesType            `json:"security_group_entries"`
	ConfiguredSecurityGroupID ConfiguredSecurityGroupIdType `json:"configured_security_group_id"`
	FQName                    []string                      `json:"fq_name"`

	AccessControlLists []*AccessControlList `json:"access_control_lists"`
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
		ConfiguredSecurityGroupID: MakeConfiguredSecurityGroupIdType(),
		FQName:          []string{},
		SecurityGroupID: MakeSecurityGroupIdType(),
		UUID:            "",
		ParentUUID:      "",
		ParentType:      "",
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
	}
}

// MakeSecurityGroupSlice() makes a slice of SecurityGroup
func MakeSecurityGroupSlice() []*SecurityGroup {
	return []*SecurityGroup{}
}
