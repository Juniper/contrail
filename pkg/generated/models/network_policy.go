package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeNetworkPolicy makes NetworkPolicy
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
		NetworkPolicyEntries: MakePolicyEntriesType(),
	}
}

// MakeNetworkPolicy makes NetworkPolicy
func InterfaceToNetworkPolicy(i interface{}) *NetworkPolicy {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &NetworkPolicy{
		//TODO(nati): Apply default
		UUID:                 schema.InterfaceToString(m["uuid"]),
		ParentUUID:           schema.InterfaceToString(m["parent_uuid"]),
		ParentType:           schema.InterfaceToString(m["parent_type"]),
		FQName:               schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          schema.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		NetworkPolicyEntries: InterfaceToPolicyEntriesType(m["network_policy_entries"]),
	}
}

// MakeNetworkPolicySlice() makes a slice of NetworkPolicy
func MakeNetworkPolicySlice() []*NetworkPolicy {
	return []*NetworkPolicy{}
}

// InterfaceToNetworkPolicySlice() makes a slice of NetworkPolicy
func InterfaceToNetworkPolicySlice(i interface{}) []*NetworkPolicy {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*NetworkPolicy{}
	for _, item := range list {
		result = append(result, InterfaceToNetworkPolicy(item))
	}
	return result
}
