package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeSecurityGroup makes SecurityGroup
// nolint
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
		ConfigurationVersion:      0,
		SecurityGroupEntries:      MakePolicyEntriesType(),
		ConfiguredSecurityGroupID: 0,
		SecurityGroupID:           0,
	}
}

// MakeSecurityGroup makes SecurityGroup
// nolint
func InterfaceToSecurityGroup(i interface{}) *SecurityGroup {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SecurityGroup{
		//TODO(nati): Apply default
		UUID:                      common.InterfaceToString(m["uuid"]),
		ParentUUID:                common.InterfaceToString(m["parent_uuid"]),
		ParentType:                common.InterfaceToString(m["parent_type"]),
		FQName:                    common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               common.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:      common.InterfaceToInt64(m["configuration_version"]),
		SecurityGroupEntries:      InterfaceToPolicyEntriesType(m["security_group_entries"]),
		ConfiguredSecurityGroupID: common.InterfaceToInt64(m["configured_security_group_id"]),
		SecurityGroupID:           common.InterfaceToInt64(m["security_group_id"]),
	}
}

// MakeSecurityGroupSlice() makes a slice of SecurityGroup
// nolint
func MakeSecurityGroupSlice() []*SecurityGroup {
	return []*SecurityGroup{}
}

// InterfaceToSecurityGroupSlice() makes a slice of SecurityGroup
// nolint
func InterfaceToSecurityGroupSlice(i interface{}) []*SecurityGroup {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SecurityGroup{}
	for _, item := range list {
		result = append(result, InterfaceToSecurityGroup(item))
	}
	return result
}
