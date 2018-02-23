package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeFirewallPolicy makes FirewallPolicy
func MakeFirewallPolicy() *FirewallPolicy {
	return &FirewallPolicy{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeFirewallPolicy makes FirewallPolicy
func InterfaceToFirewallPolicy(i interface{}) *FirewallPolicy {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallPolicy{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
	}
}

// MakeFirewallPolicySlice() makes a slice of FirewallPolicy
func MakeFirewallPolicySlice() []*FirewallPolicy {
	return []*FirewallPolicy{}
}

// InterfaceToFirewallPolicySlice() makes a slice of FirewallPolicy
func InterfaceToFirewallPolicySlice(i interface{}) []*FirewallPolicy {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallPolicy{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallPolicy(item))
	}
	return result
}
