package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeFirewallRuleMatchTagsType makes FirewallRuleMatchTagsType
func MakeFirewallRuleMatchTagsType() *FirewallRuleMatchTagsType {
	return &FirewallRuleMatchTagsType{
		//TODO(nati): Apply default
		TagList: []string{},
	}
}

// MakeFirewallRuleMatchTagsType makes FirewallRuleMatchTagsType
func InterfaceToFirewallRuleMatchTagsType(i interface{}) *FirewallRuleMatchTagsType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallRuleMatchTagsType{
		//TODO(nati): Apply default
		TagList: schema.InterfaceToStringList(m["tag_list"]),
	}
}

// MakeFirewallRuleMatchTagsTypeSlice() makes a slice of FirewallRuleMatchTagsType
func MakeFirewallRuleMatchTagsTypeSlice() []*FirewallRuleMatchTagsType {
	return []*FirewallRuleMatchTagsType{}
}

// InterfaceToFirewallRuleMatchTagsTypeSlice() makes a slice of FirewallRuleMatchTagsType
func InterfaceToFirewallRuleMatchTagsTypeSlice(i interface{}) []*FirewallRuleMatchTagsType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallRuleMatchTagsType{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleMatchTagsType(item))
	}
	return result
}
