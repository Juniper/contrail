package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeLoadbalancerHealthmonitor makes LoadbalancerHealthmonitor
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
func InterfaceToLoadbalancerHealthmonitor(i interface{}) *LoadbalancerHealthmonitor {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerHealthmonitor{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		LoadbalancerHealthmonitorProperties: InterfaceToLoadbalancerHealthmonitorType(m["loadbalancer_healthmonitor_properties"]),
	}
}

// MakeLoadbalancerHealthmonitorSlice() makes a slice of LoadbalancerHealthmonitor
func MakeLoadbalancerHealthmonitorSlice() []*LoadbalancerHealthmonitor {
	return []*LoadbalancerHealthmonitor{}
}

// InterfaceToLoadbalancerHealthmonitorSlice() makes a slice of LoadbalancerHealthmonitor
func InterfaceToLoadbalancerHealthmonitorSlice(i interface{}) []*LoadbalancerHealthmonitor {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerHealthmonitor{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerHealthmonitor(item))
	}
	return result
}
