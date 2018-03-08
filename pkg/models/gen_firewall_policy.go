package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFirewallPolicy makes FirewallPolicy
// nolint
func MakeFirewallPolicy() *FirewallPolicy {
	return &FirewallPolicy{
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

// MakeFirewallPolicy makes FirewallPolicy
// nolint
func InterfaceToFirewallPolicy(i interface{}) *FirewallPolicy {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallPolicy{
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

		FirewallRuleRefs: InterfaceToFirewallPolicyFirewallRuleRefs(m["firewall_rule_refs"]),

		SecurityLoggingObjectRefs: InterfaceToFirewallPolicySecurityLoggingObjectRefs(m["security_logging_object_refs"]),
	}
}

func InterfaceToFirewallPolicyFirewallRuleRefs(i interface{}) []*FirewallPolicyFirewallRuleRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*FirewallPolicyFirewallRuleRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &FirewallPolicyFirewallRuleRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),

			Attr: InterfaceToFirewallSequence(m["attr"]),
		})
	}

	return result
}

func InterfaceToFirewallPolicySecurityLoggingObjectRefs(i interface{}) []*FirewallPolicySecurityLoggingObjectRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*FirewallPolicySecurityLoggingObjectRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &FirewallPolicySecurityLoggingObjectRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeFirewallPolicySlice() makes a slice of FirewallPolicy
// nolint
func MakeFirewallPolicySlice() []*FirewallPolicy {
	return []*FirewallPolicy{}
}

// InterfaceToFirewallPolicySlice() makes a slice of FirewallPolicy
// nolint
func InterfaceToFirewallPolicySlice(i interface{}) []*FirewallPolicy {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallPolicy{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallPolicy(item))
	}
	return result
}
