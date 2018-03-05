package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeGlobalQosConfig makes GlobalQosConfig
// nolint
func MakeGlobalQosConfig() *GlobalQosConfig {
	return &GlobalQosConfig{
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
		ControlTrafficDSCP:   MakeControlTrafficDscpType(),
	}
}

// MakeGlobalQosConfig makes GlobalQosConfig
// nolint
func InterfaceToGlobalQosConfig(i interface{}) *GlobalQosConfig {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &GlobalQosConfig{
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
		ControlTrafficDSCP:   InterfaceToControlTrafficDscpType(m["control_traffic_dscp"]),
	}
}

// MakeGlobalQosConfigSlice() makes a slice of GlobalQosConfig
// nolint
func MakeGlobalQosConfigSlice() []*GlobalQosConfig {
	return []*GlobalQosConfig{}
}

// InterfaceToGlobalQosConfigSlice() makes a slice of GlobalQosConfig
// nolint
func InterfaceToGlobalQosConfigSlice(i interface{}) []*GlobalQosConfig {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*GlobalQosConfig{}
	for _, item := range list {
		result = append(result, InterfaceToGlobalQosConfig(item))
	}
	return result
}
