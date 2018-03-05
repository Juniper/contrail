package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRoutingPolicy makes RoutingPolicy
// nolint
func MakeRoutingPolicy() *RoutingPolicy {
	return &RoutingPolicy{
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

// MakeRoutingPolicy makes RoutingPolicy
// nolint
func InterfaceToRoutingPolicy(i interface{}) *RoutingPolicy {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RoutingPolicy{
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

// MakeRoutingPolicySlice() makes a slice of RoutingPolicy
// nolint
func MakeRoutingPolicySlice() []*RoutingPolicy {
	return []*RoutingPolicy{}
}

// InterfaceToRoutingPolicySlice() makes a slice of RoutingPolicy
// nolint
func InterfaceToRoutingPolicySlice(i interface{}) []*RoutingPolicy {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RoutingPolicy{}
	for _, item := range list {
		result = append(result, InterfaceToRoutingPolicy(item))
	}
	return result
}
