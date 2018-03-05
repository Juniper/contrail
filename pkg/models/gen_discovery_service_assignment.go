package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeDiscoveryServiceAssignment makes DiscoveryServiceAssignment
// nolint
func MakeDiscoveryServiceAssignment() *DiscoveryServiceAssignment {
	return &DiscoveryServiceAssignment{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		ConfigurationVersion: 0,
	}
}

// MakeDiscoveryServiceAssignment makes DiscoveryServiceAssignment
// nolint
func InterfaceToDiscoveryServiceAssignment(i interface{}) *DiscoveryServiceAssignment {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DiscoveryServiceAssignment{
		//TODO(nati): Apply default
		UUID:                 common.InterfaceToString(m["uuid"]),
		ParentUUID:           common.InterfaceToString(m["parent_uuid"]),
		ParentType:           common.InterfaceToString(m["parent_type"]),
		FQName:               common.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          common.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion: common.InterfaceToInt64(m["configuration_version"]),
	}
}

// MakeDiscoveryServiceAssignmentSlice() makes a slice of DiscoveryServiceAssignment
// nolint
func MakeDiscoveryServiceAssignmentSlice() []*DiscoveryServiceAssignment {
	return []*DiscoveryServiceAssignment{}
}

// InterfaceToDiscoveryServiceAssignmentSlice() makes a slice of DiscoveryServiceAssignment
// nolint
func InterfaceToDiscoveryServiceAssignmentSlice(i interface{}) []*DiscoveryServiceAssignment {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DiscoveryServiceAssignment{}
	for _, item := range list {
		result = append(result, InterfaceToDiscoveryServiceAssignment(item))
	}
	return result
}
