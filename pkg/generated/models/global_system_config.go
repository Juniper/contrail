package models

// GlobalSystemConfig

import "encoding/json"

// GlobalSystemConfig
type GlobalSystemConfig struct {
	ConfigVersion             string                         `json:"config_version"`
	BgpaasParameters          *BGPaaServiceParametersType    `json:"bgpaas_parameters"`
	AlarmEnable               bool                           `json:"alarm_enable"`
	IbgpAutoMesh              bool                           `json:"ibgp_auto_mesh"`
	GracefulRestartParameters *GracefulRestartParametersType `json:"graceful_restart_parameters"`
	AutonomousSystem          AutonomousSystemType           `json:"autonomous_system"`
	DisplayName               string                         `json:"display_name"`
	Annotations               *KeyValuePairs                 `json:"annotations"`
	UUID                      string                         `json:"uuid"`
	ParentType                string                         `json:"parent_type"`
	PluginTuning              *PluginProperties              `json:"plugin_tuning"`
	UserDefinedLogStatistics  *UserDefinedLogStatList        `json:"user_defined_log_statistics"`
	MacMoveControl            *MACMoveLimitControlType       `json:"mac_move_control"`
	MacLimitControl           *MACLimitControlType           `json:"mac_limit_control"`
	MacAgingTime              MACAgingTime                   `json:"mac_aging_time"`
	BGPAlwaysCompareMed       bool                           `json:"bgp_always_compare_med"`
	IPFabricSubnets           *SubnetListType                `json:"ip_fabric_subnets"`
	Perms2                    *PermType2                     `json:"perms2"`
	ParentUUID                string                         `json:"parent_uuid"`
	FQName                    []string                       `json:"fq_name"`
	IDPerms                   *IdPermsType                   `json:"id_perms"`

	BGPRouterRefs []*GlobalSystemConfigBGPRouterRef `json:"bgp_router_refs"`

	Alarms               []*Alarm               `json:"alarms"`
	AnalyticsNodes       []*AnalyticsNode       `json:"analytics_nodes"`
	APIAccessLists       []*APIAccessList       `json:"api_access_lists"`
	ConfigNodes          []*ConfigNode          `json:"config_nodes"`
	DatabaseNodes        []*DatabaseNode        `json:"database_nodes"`
	GlobalQosConfigs     []*GlobalQosConfig     `json:"global_qos_configs"`
	GlobalVrouterConfigs []*GlobalVrouterConfig `json:"global_vrouter_configs"`
	PhysicalRouters      []*PhysicalRouter      `json:"physical_routers"`
	ServiceApplianceSets []*ServiceApplianceSet `json:"service_appliance_sets"`
	VirtualRouters       []*VirtualRouter       `json:"virtual_routers"`
}

// GlobalSystemConfigBGPRouterRef references each other
type GlobalSystemConfigBGPRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *GlobalSystemConfig) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeGlobalSystemConfig makes GlobalSystemConfig
func MakeGlobalSystemConfig() *GlobalSystemConfig {
	return &GlobalSystemConfig{
		//TODO(nati): Apply default
		AutonomousSystem:          MakeAutonomousSystemType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		ConfigVersion:             "",
		BgpaasParameters:          MakeBGPaaServiceParametersType(),
		AlarmEnable:               false,
		IbgpAutoMesh:              false,
		GracefulRestartParameters: MakeGracefulRestartParametersType(),
		UUID:                     "",
		ParentType:               "",
		PluginTuning:             MakePluginProperties(),
		UserDefinedLogStatistics: MakeUserDefinedLogStatList(),
		MacMoveControl:           MakeMACMoveLimitControlType(),
		MacLimitControl:          MakeMACLimitControlType(),
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		MacAgingTime:             MakeMACAgingTime(),
		BGPAlwaysCompareMed:      false,
		IPFabricSubnets:          MakeSubnetListType(),
		Perms2:                   MakePermType2(),
		ParentUUID:               "",
	}
}

// MakeGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
	return []*GlobalSystemConfig{}
}
