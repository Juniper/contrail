package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancer makes Loadbalancer
// nolint
func MakeLoadbalancer() *Loadbalancer {
	return &Loadbalancer{
		//TODO(nati): Apply default
		UUID:                   "",
		ParentUUID:             "",
		ParentType:             "",
		FQName:                 []string{},
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		Annotations:            MakeKeyValuePairs(),
		Perms2:                 MakePermType2(),
		LoadbalancerProperties: MakeLoadbalancerType(),
		LoadbalancerProvider:   "",
	}
}

// MakeLoadbalancer makes Loadbalancer
// nolint
func InterfaceToLoadbalancer(i interface{}) *Loadbalancer {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Loadbalancer{
		//TODO(nati): Apply default
		UUID:                   common.InterfaceToString(m["uuid"]),
		ParentUUID:             common.InterfaceToString(m["parent_uuid"]),
		ParentType:             common.InterfaceToString(m["parent_type"]),
		FQName:                 common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:            common.InterfaceToString(m["display_name"]),
		Annotations:            InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                 InterfaceToPermType2(m["perms2"]),
		LoadbalancerProperties: InterfaceToLoadbalancerType(m["loadbalancer_properties"]),
		LoadbalancerProvider:   common.InterfaceToString(m["loadbalancer_provider"]),
	}
}

// MakeLoadbalancerSlice() makes a slice of Loadbalancer
// nolint
func MakeLoadbalancerSlice() []*Loadbalancer {
	return []*Loadbalancer{}
}

// InterfaceToLoadbalancerSlice() makes a slice of Loadbalancer
// nolint
func InterfaceToLoadbalancerSlice(i interface{}) []*Loadbalancer {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Loadbalancer{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancer(item))
	}
	return result
}
