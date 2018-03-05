package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeBGPRouter makes BGPRouter
// nolint
func MakeBGPRouter() *BGPRouter {
	return &BGPRouter{
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
	}
}

// MakeBGPRouter makes BGPRouter
// nolint
func InterfaceToBGPRouter(i interface{}) *BGPRouter {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BGPRouter{
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
	}
}

// MakeBGPRouterSlice() makes a slice of BGPRouter
// nolint
func MakeBGPRouterSlice() []*BGPRouter {
	return []*BGPRouter{}
}

// InterfaceToBGPRouterSlice() makes a slice of BGPRouter
// nolint
func InterfaceToBGPRouterSlice(i interface{}) []*BGPRouter {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BGPRouter{}
	for _, item := range list {
		result = append(result, InterfaceToBGPRouter(item))
	}
	return result
}
