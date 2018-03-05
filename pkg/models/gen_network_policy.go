package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeNetworkPolicy makes NetworkPolicy
// nolint
func MakeNetworkPolicy() *NetworkPolicy {
	return &NetworkPolicy{
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
		NetworkPolicyEntries: MakePolicyEntriesType(),
	}
}

// MakeNetworkPolicy makes NetworkPolicy
// nolint
func InterfaceToNetworkPolicy(i interface{}) *NetworkPolicy {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &NetworkPolicy{
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
		NetworkPolicyEntries: InterfaceToPolicyEntriesType(m["network_policy_entries"]),
	}
}

// MakeNetworkPolicySlice() makes a slice of NetworkPolicy
// nolint
func MakeNetworkPolicySlice() []*NetworkPolicy {
	return []*NetworkPolicy{}
}

// InterfaceToNetworkPolicySlice() makes a slice of NetworkPolicy
// nolint
func InterfaceToNetworkPolicySlice(i interface{}) []*NetworkPolicy {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*NetworkPolicy{}
	for _, item := range list {
		result = append(result, InterfaceToNetworkPolicy(item))
	}
	return result
}
