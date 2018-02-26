package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeGlobalSystemConfig makes GlobalSystemConfig
func MakeGlobalSystemConfig() *GlobalSystemConfig {
	return &GlobalSystemConfig{
		//TODO(nati): Apply default
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		ConfigVersion:             "",
		BgpaasParameters:          MakeBGPaaServiceParametersType(),
		AlarmEnable:               false,
		MacMoveControl:            MakeMACMoveLimitControlType(),
		PluginTuning:              MakePluginProperties(),
		IbgpAutoMesh:              false,
		MacAgingTime:              0,
		BGPAlwaysCompareMed:       false,
		UserDefinedLogStatistics:  MakeUserDefinedLogStatList(),
		GracefulRestartParameters: MakeGracefulRestartParametersType(),
		IPFabricSubnets:           MakeSubnetListType(),
		AutonomousSystem:          0,
		MacLimitControl:           MakeMACLimitControlType(),
	}
}

// MakeGlobalSystemConfig makes GlobalSystemConfig
func InterfaceToGlobalSystemConfig(i interface{}) *GlobalSystemConfig {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &GlobalSystemConfig{
		//TODO(nati): Apply default
		UUID:                      schema.InterfaceToString(m["uuid"]),
		ParentUUID:                schema.InterfaceToString(m["parent_uuid"]),
		ParentType:                schema.InterfaceToString(m["parent_type"]),
		FQName:                    schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               schema.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
		ConfigVersion:             schema.InterfaceToString(m["config_version"]),
		BgpaasParameters:          InterfaceToBGPaaServiceParametersType(m["bgpaas_parameters"]),
		AlarmEnable:               schema.InterfaceToBool(m["alarm_enable"]),
		MacMoveControl:            InterfaceToMACMoveLimitControlType(m["mac_move_control"]),
		PluginTuning:              InterfaceToPluginProperties(m["plugin_tuning"]),
		IbgpAutoMesh:              schema.InterfaceToBool(m["ibgp_auto_mesh"]),
		MacAgingTime:              schema.InterfaceToInt64(m["mac_aging_time"]),
		BGPAlwaysCompareMed:       schema.InterfaceToBool(m["bgp_always_compare_med"]),
		UserDefinedLogStatistics:  InterfaceToUserDefinedLogStatList(m["user_defined_log_statistics"]),
		GracefulRestartParameters: InterfaceToGracefulRestartParametersType(m["graceful_restart_parameters"]),
		IPFabricSubnets:           InterfaceToSubnetListType(m["ip_fabric_subnets"]),
		AutonomousSystem:          schema.InterfaceToInt64(m["autonomous_system"]),
		MacLimitControl:           InterfaceToMACLimitControlType(m["mac_limit_control"]),
	}
}

// MakeGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
	return []*GlobalSystemConfig{}
}

// InterfaceToGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
func InterfaceToGlobalSystemConfigSlice(i interface{}) []*GlobalSystemConfig {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*GlobalSystemConfig{}
	for _, item := range list {
		result = append(result, InterfaceToGlobalSystemConfig(item))
	}
	return result
}
