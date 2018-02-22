package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

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
		ConfiguredSecurityGroupID: 0,
		SecurityGroupID:           0,
	}
}

// MakeSecurityGroup makes SecurityGroup
func InterfaceToSecurityGroup(i interface{}) *SecurityGroup {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SecurityGroup{
		//TODO(nati): Apply default
		UUID:                      schema.InterfaceToString(m["uuid"]),
		ParentUUID:                schema.InterfaceToString(m["parent_uuid"]),
		ParentType:                schema.InterfaceToString(m["parent_type"]),
		FQName:                    schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               schema.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
		SecurityGroupEntries:      InterfaceToPolicyEntriesType(m["security_group_entries"]),
		ConfiguredSecurityGroupID: schema.InterfaceToInt64(m["configured_security_group_id"]),
		SecurityGroupID:           schema.InterfaceToInt64(m["security_group_id"]),
	}
}

// MakeSecurityGroupSlice() makes a slice of SecurityGroup
func MakeSecurityGroupSlice() []*SecurityGroup {
	return []*SecurityGroup{}
}

// InterfaceToSecurityGroupSlice() makes a slice of SecurityGroup
func InterfaceToSecurityGroupSlice(i interface{}) []*SecurityGroup {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SecurityGroup{}
	for _, item := range list {
		result = append(result, InterfaceToSecurityGroup(item))
	}
	return result
}
