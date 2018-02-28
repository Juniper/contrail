package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceGroup makes ServiceGroup
// nolint
func MakeServiceGroup() *ServiceGroup {
	return &ServiceGroup{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ServiceGroupFirewallServiceList: MakeFirewallServiceGroupType(),
	}
}

// MakeServiceGroup makes ServiceGroup
// nolint
func InterfaceToServiceGroup(i interface{}) *ServiceGroup {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceGroup{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		ServiceGroupFirewallServiceList: InterfaceToFirewallServiceGroupType(m["service_group_firewall_service_list"]),
	}
}

// MakeServiceGroupSlice() makes a slice of ServiceGroup
// nolint
func MakeServiceGroupSlice() []*ServiceGroup {
	return []*ServiceGroup{}
}

// InterfaceToServiceGroupSlice() makes a slice of ServiceGroup
// nolint
func InterfaceToServiceGroupSlice(i interface{}) []*ServiceGroup {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceGroup{}
	for _, item := range list {
		result = append(result, InterfaceToServiceGroup(item))
	}
	return result
}
