package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFirewallRule makes FirewallRule
// nolint
func MakeFirewallRule() *FirewallRule {
	return &FirewallRule{
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
		Endpoint1:            MakeFirewallRuleEndpointType(),
		Endpoint2:            MakeFirewallRuleEndpointType(),
		ActionList:           MakeActionListType(),
		Service:              MakeFirewallServiceType(),
		Direction:            "",
		MatchTagTypes:        MakeFirewallRuleMatchTagsTypeIdList(),
		MatchTags:            MakeFirewallRuleMatchTagsType(),
	}
}

// MakeFirewallRule makes FirewallRule
// nolint
func InterfaceToFirewallRule(i interface{}) *FirewallRule {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallRule{
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
		Endpoint1:            InterfaceToFirewallRuleEndpointType(m["endpoint_1"]),
		Endpoint2:            InterfaceToFirewallRuleEndpointType(m["endpoint_2"]),
		ActionList:           InterfaceToActionListType(m["action_list"]),
		Service:              InterfaceToFirewallServiceType(m["service"]),
		Direction:            common.InterfaceToString(m["direction"]),
		MatchTagTypes:        InterfaceToFirewallRuleMatchTagsTypeIdList(m["match_tag_types"]),
		MatchTags:            InterfaceToFirewallRuleMatchTagsType(m["match_tags"]),
	}
}

// MakeFirewallRuleSlice() makes a slice of FirewallRule
// nolint
func MakeFirewallRuleSlice() []*FirewallRule {
	return []*FirewallRule{}
}

// InterfaceToFirewallRuleSlice() makes a slice of FirewallRule
// nolint
func InterfaceToFirewallRuleSlice(i interface{}) []*FirewallRule {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallRule{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallRule(item))
	}
	return result
}
