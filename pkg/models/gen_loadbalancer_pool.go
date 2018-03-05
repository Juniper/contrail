package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancerPool makes LoadbalancerPool
// nolint
func MakeLoadbalancerPool() *LoadbalancerPool {
	return &LoadbalancerPool{
		//TODO(nati): Apply default
		UUID:                             "",
		ParentUUID:                       "",
		ParentType:                       "",
		FQName:                           []string{},
		IDPerms:                          MakeIdPermsType(),
		DisplayName:                      "",
		Annotations:                      MakeKeyValuePairs(),
		Perms2:                           MakePermType2(),
		ConfigurationVersion:             0,
		LoadbalancerPoolProperties:       MakeLoadbalancerPoolType(),
		LoadbalancerPoolCustomAttributes: MakeKeyValuePairs(),
		LoadbalancerPoolProvider:         "",
	}
}

// MakeLoadbalancerPool makes LoadbalancerPool
// nolint
func InterfaceToLoadbalancerPool(i interface{}) *LoadbalancerPool {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerPool{
		//TODO(nati): Apply default
		UUID:                             common.InterfaceToString(m["uuid"]),
		ParentUUID:                       common.InterfaceToString(m["parent_uuid"]),
		ParentType:                       common.InterfaceToString(m["parent_type"]),
		FQName:                           common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                          InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                      common.InterfaceToString(m["display_name"]),
		Annotations:                      InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                           InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:             common.InterfaceToInt64(m["configuration_version"]),
		LoadbalancerPoolProperties:       InterfaceToLoadbalancerPoolType(m["loadbalancer_pool_properties"]),
		LoadbalancerPoolCustomAttributes: InterfaceToKeyValuePairs(m["loadbalancer_pool_custom_attributes"]),
		LoadbalancerPoolProvider:         common.InterfaceToString(m["loadbalancer_pool_provider"]),
	}
}

// MakeLoadbalancerPoolSlice() makes a slice of LoadbalancerPool
// nolint
func MakeLoadbalancerPoolSlice() []*LoadbalancerPool {
	return []*LoadbalancerPool{}
}

// InterfaceToLoadbalancerPoolSlice() makes a slice of LoadbalancerPool
// nolint
func InterfaceToLoadbalancerPoolSlice(i interface{}) []*LoadbalancerPool {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerPool{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerPool(item))
	}
	return result
}
