package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFirewallRuleMatchTagsType makes FirewallRuleMatchTagsType
// nolint
func MakeFirewallRuleMatchTagsType() *FirewallRuleMatchTagsType {
	return &FirewallRuleMatchTagsType{
		//TODO(nati): Apply default
		TagList: []string{},
	}
}

// MakeFirewallRuleMatchTagsType makes FirewallRuleMatchTagsType
// nolint
func InterfaceToFirewallRuleMatchTagsType(i interface{}) *FirewallRuleMatchTagsType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallRuleMatchTagsType{
		//TODO(nati): Apply default
		TagList: common.InterfaceToStringList(m["tag_list"]),
	}
}

// MakeFirewallRuleMatchTagsTypeSlice() makes a slice of FirewallRuleMatchTagsType
// nolint
func MakeFirewallRuleMatchTagsTypeSlice() []*FirewallRuleMatchTagsType {
	return []*FirewallRuleMatchTagsType{}
}

// InterfaceToFirewallRuleMatchTagsTypeSlice() makes a slice of FirewallRuleMatchTagsType
// nolint
func InterfaceToFirewallRuleMatchTagsTypeSlice(i interface{}) []*FirewallRuleMatchTagsType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallRuleMatchTagsType{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleMatchTagsType(item))
	}
	return result
}
