package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancerHealthmonitor makes LoadbalancerHealthmonitor
// nolint
func MakeLoadbalancerHealthmonitor() *LoadbalancerHealthmonitor {
	return &LoadbalancerHealthmonitor{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		LoadbalancerHealthmonitorProperties: MakeLoadbalancerHealthmonitorType(),
	}
}

// MakeLoadbalancerHealthmonitor makes LoadbalancerHealthmonitor
// nolint
func InterfaceToLoadbalancerHealthmonitor(i interface{}) *LoadbalancerHealthmonitor {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerHealthmonitor{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		LoadbalancerHealthmonitorProperties: InterfaceToLoadbalancerHealthmonitorType(m["loadbalancer_healthmonitor_properties"]),
	}
}

// MakeLoadbalancerHealthmonitorSlice() makes a slice of LoadbalancerHealthmonitor
// nolint
func MakeLoadbalancerHealthmonitorSlice() []*LoadbalancerHealthmonitor {
	return []*LoadbalancerHealthmonitor{}
}

// InterfaceToLoadbalancerHealthmonitorSlice() makes a slice of LoadbalancerHealthmonitor
// nolint
func InterfaceToLoadbalancerHealthmonitorSlice(i interface{}) []*LoadbalancerHealthmonitor {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerHealthmonitor{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerHealthmonitor(item))
	}
	return result
}
