package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancerListener makes LoadbalancerListener
// nolint
func MakeLoadbalancerListener() *LoadbalancerListener {
	return &LoadbalancerListener{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		LoadbalancerListenerProperties: MakeLoadbalancerListenerType(),
	}
}

// MakeLoadbalancerListener makes LoadbalancerListener
// nolint
func InterfaceToLoadbalancerListener(i interface{}) *LoadbalancerListener {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerListener{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		LoadbalancerListenerProperties: InterfaceToLoadbalancerListenerType(m["loadbalancer_listener_properties"]),
	}
}

// MakeLoadbalancerListenerSlice() makes a slice of LoadbalancerListener
// nolint
func MakeLoadbalancerListenerSlice() []*LoadbalancerListener {
	return []*LoadbalancerListener{}
}

// InterfaceToLoadbalancerListenerSlice() makes a slice of LoadbalancerListener
// nolint
func InterfaceToLoadbalancerListenerSlice(i interface{}) []*LoadbalancerListener {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerListener{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerListener(item))
	}
	return result
}
