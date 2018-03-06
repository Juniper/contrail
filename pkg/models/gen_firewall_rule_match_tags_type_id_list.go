package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFirewallRuleMatchTagsTypeIdList makes FirewallRuleMatchTagsTypeIdList
// nolint
func MakeFirewallRuleMatchTagsTypeIdList() *FirewallRuleMatchTagsTypeIdList {
	return &FirewallRuleMatchTagsTypeIdList{
		//TODO(nati): Apply default

		TagType: []int64{},
	}
}

// MakeFirewallRuleMatchTagsTypeIdList makes FirewallRuleMatchTagsTypeIdList
// nolint
func InterfaceToFirewallRuleMatchTagsTypeIdList(i interface{}) *FirewallRuleMatchTagsTypeIdList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallRuleMatchTagsTypeIdList{
		//TODO(nati): Apply default

		TagType: common.InterfaceToInt64List(m["tag_type"]),
	}
}

// MakeFirewallRuleMatchTagsTypeIdListSlice() makes a slice of FirewallRuleMatchTagsTypeIdList
// nolint
func MakeFirewallRuleMatchTagsTypeIdListSlice() []*FirewallRuleMatchTagsTypeIdList {
	return []*FirewallRuleMatchTagsTypeIdList{}
}

// InterfaceToFirewallRuleMatchTagsTypeIdListSlice() makes a slice of FirewallRuleMatchTagsTypeIdList
// nolint
func InterfaceToFirewallRuleMatchTagsTypeIdListSlice(i interface{}) []*FirewallRuleMatchTagsTypeIdList {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallRuleMatchTagsTypeIdList{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleMatchTagsTypeIdList(item))
	}
	return result
}
