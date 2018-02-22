package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeGlobalQosConfig makes GlobalQosConfig
func MakeGlobalQosConfig() *GlobalQosConfig {
	return &GlobalQosConfig{
		//TODO(nati): Apply default
		UUID:               "",
		ParentUUID:         "",
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		ControlTrafficDSCP: MakeControlTrafficDscpType(),
	}
}

// MakeGlobalQosConfig makes GlobalQosConfig
func InterfaceToGlobalQosConfig(i interface{}) *GlobalQosConfig {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &GlobalQosConfig{
		//TODO(nati): Apply default
		UUID:               schema.InterfaceToString(m["uuid"]),
		ParentUUID:         schema.InterfaceToString(m["parent_uuid"]),
		ParentType:         schema.InterfaceToString(m["parent_type"]),
		FQName:             schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:            InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:        schema.InterfaceToString(m["display_name"]),
		Annotations:        InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:             InterfaceToPermType2(m["perms2"]),
		ControlTrafficDSCP: InterfaceToControlTrafficDscpType(m["control_traffic_dscp"]),
	}
}

// MakeGlobalQosConfigSlice() makes a slice of GlobalQosConfig
func MakeGlobalQosConfigSlice() []*GlobalQosConfig {
	return []*GlobalQosConfig{}
}

// InterfaceToGlobalQosConfigSlice() makes a slice of GlobalQosConfig
func InterfaceToGlobalQosConfigSlice(i interface{}) []*GlobalQosConfig {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*GlobalQosConfig{}
	for _, item := range list {
		result = append(result, InterfaceToGlobalQosConfig(item))
	}
	return result
}
