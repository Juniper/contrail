package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeNetworkDeviceConfig makes NetworkDeviceConfig
// nolint
func MakeNetworkDeviceConfig() *NetworkDeviceConfig {
	return &NetworkDeviceConfig{
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

// MakeNetworkDeviceConfig makes NetworkDeviceConfig
// nolint
func InterfaceToNetworkDeviceConfig(i interface{}) *NetworkDeviceConfig {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &NetworkDeviceConfig{
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

		PhysicalRouterRefs: InterfaceToNetworkDeviceConfigPhysicalRouterRefs(m["physical_router_refs"]),
	}
}

func InterfaceToNetworkDeviceConfigPhysicalRouterRefs(i interface{}) []*NetworkDeviceConfigPhysicalRouterRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*NetworkDeviceConfigPhysicalRouterRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &NetworkDeviceConfigPhysicalRouterRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeNetworkDeviceConfigSlice() makes a slice of NetworkDeviceConfig
// nolint
func MakeNetworkDeviceConfigSlice() []*NetworkDeviceConfig {
	return []*NetworkDeviceConfig{}
}

// InterfaceToNetworkDeviceConfigSlice() makes a slice of NetworkDeviceConfig
// nolint
func InterfaceToNetworkDeviceConfigSlice(i interface{}) []*NetworkDeviceConfig {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*NetworkDeviceConfig{}
	for _, item := range list {
		result = append(result, InterfaceToNetworkDeviceConfig(item))
	}
	return result
}
