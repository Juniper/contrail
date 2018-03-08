package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeGlobalSystemConfig makes GlobalSystemConfig
// nolint
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
		ConfigurationVersion:      0,
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
// nolint
func InterfaceToGlobalSystemConfig(i interface{}) *GlobalSystemConfig {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &GlobalSystemConfig{
		//TODO(nati): Apply default
		UUID:                      common.InterfaceToString(m["uuid"]),
		ParentUUID:                common.InterfaceToString(m["parent_uuid"]),
		ParentType:                common.InterfaceToString(m["parent_type"]),
		FQName:                    common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               common.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:      common.InterfaceToInt64(m["configuration_version"]),
		ConfigVersion:             common.InterfaceToString(m["config_version"]),
		BgpaasParameters:          InterfaceToBGPaaServiceParametersType(m["bgpaas_parameters"]),
		AlarmEnable:               common.InterfaceToBool(m["alarm_enable"]),
		MacMoveControl:            InterfaceToMACMoveLimitControlType(m["mac_move_control"]),
		PluginTuning:              InterfaceToPluginProperties(m["plugin_tuning"]),
		IbgpAutoMesh:              common.InterfaceToBool(m["ibgp_auto_mesh"]),
		MacAgingTime:              common.InterfaceToInt64(m["mac_aging_time"]),
		BGPAlwaysCompareMed:       common.InterfaceToBool(m["bgp_always_compare_med"]),
		UserDefinedLogStatistics:  InterfaceToUserDefinedLogStatList(m["user_defined_log_statistics"]),
		GracefulRestartParameters: InterfaceToGracefulRestartParametersType(m["graceful_restart_parameters"]),
		IPFabricSubnets:           InterfaceToSubnetListType(m["ip_fabric_subnets"]),
		AutonomousSystem:          common.InterfaceToInt64(m["autonomous_system"]),
		MacLimitControl:           InterfaceToMACLimitControlType(m["mac_limit_control"]),

		BGPRouterRefs: InterfaceToGlobalSystemConfigBGPRouterRefs(m["bgp_router_refs"]),
	}
}

func InterfaceToGlobalSystemConfigBGPRouterRefs(i interface{}) []*GlobalSystemConfigBGPRouterRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*GlobalSystemConfigBGPRouterRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &GlobalSystemConfigBGPRouterRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
// nolint
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
	return []*GlobalSystemConfig{}
}

// InterfaceToGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
// nolint
func InterfaceToGlobalSystemConfigSlice(i interface{}) []*GlobalSystemConfig {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*GlobalSystemConfig{}
	for _, item := range list {
		result = append(result, InterfaceToGlobalSystemConfig(item))
	}
	return result
}
