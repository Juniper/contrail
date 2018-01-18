package models

// GlobalSystemConfig

import "encoding/json"

// GlobalSystemConfig
type GlobalSystemConfig struct {
	MacAgingTime              MACAgingTime                   `json:"mac_aging_time,omitempty"`
	ParentUUID                string                         `json:"parent_uuid,omitempty"`
	FQName                    []string                       `json:"fq_name,omitempty"`
	DisplayName               string                         `json:"display_name,omitempty"`
	ConfigVersion             string                         `json:"config_version,omitempty"`
	AlarmEnable               bool                           `json:"alarm_enable"`
	Perms2                    *PermType2                     `json:"perms2,omitempty"`
	ParentType                string                         `json:"parent_type,omitempty"`
	Annotations               *KeyValuePairs                 `json:"annotations,omitempty"`
	IbgpAutoMesh              bool                           `json:"ibgp_auto_mesh"`
	BGPAlwaysCompareMed       bool                           `json:"bgp_always_compare_med"`
	UserDefinedLogStatistics  *UserDefinedLogStatList        `json:"user_defined_log_statistics,omitempty"`
	GracefulRestartParameters *GracefulRestartParametersType `json:"graceful_restart_parameters,omitempty"`
	IPFabricSubnets           *SubnetListType                `json:"ip_fabric_subnets,omitempty"`
	AutonomousSystem          AutonomousSystemType           `json:"autonomous_system,omitempty"`
	MacLimitControl           *MACLimitControlType           `json:"mac_limit_control,omitempty"`
	BgpaasParameters          *BGPaaServiceParametersType    `json:"bgpaas_parameters,omitempty"`
	MacMoveControl            *MACMoveLimitControlType       `json:"mac_move_control,omitempty"`
	PluginTuning              *PluginProperties              `json:"plugin_tuning,omitempty"`
	UUID                      string                         `json:"uuid,omitempty"`
	IDPerms                   *IdPermsType                   `json:"id_perms,omitempty"`

	BGPRouterRefs []*GlobalSystemConfigBGPRouterRef `json:"bgp_router_refs,omitempty"`

	Alarms               []*Alarm               `json:"alarms,omitempty"`
	AnalyticsNodes       []*AnalyticsNode       `json:"analytics_nodes,omitempty"`
	APIAccessLists       []*APIAccessList       `json:"api_access_lists,omitempty"`
	ConfigNodes          []*ConfigNode          `json:"config_nodes,omitempty"`
	DatabaseNodes        []*DatabaseNode        `json:"database_nodes,omitempty"`
	GlobalQosConfigs     []*GlobalQosConfig     `json:"global_qos_configs,omitempty"`
	GlobalVrouterConfigs []*GlobalVrouterConfig `json:"global_vrouter_configs,omitempty"`
	PhysicalRouters      []*PhysicalRouter      `json:"physical_routers,omitempty"`
	ServiceApplianceSets []*ServiceApplianceSet `json:"service_appliance_sets,omitempty"`
	VirtualRouters       []*VirtualRouter       `json:"virtual_routers,omitempty"`
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
		AlarmEnable:               false,
		Perms2:                    MakePermType2(),
		ParentType:                "",
		Annotations:               MakeKeyValuePairs(),
		ConfigVersion:             "",
		BGPAlwaysCompareMed:       false,
		UserDefinedLogStatistics:  MakeUserDefinedLogStatList(),
		GracefulRestartParameters: MakeGracefulRestartParametersType(),
		IPFabricSubnets:           MakeSubnetListType(),
		AutonomousSystem:          MakeAutonomousSystemType(),
		MacLimitControl:           MakeMACLimitControlType(),
		IbgpAutoMesh:              false,
		MacMoveControl:            MakeMACMoveLimitControlType(),
		PluginTuning:              MakePluginProperties(),
		UUID:                      "",
		IDPerms:                   MakeIdPermsType(),
		BgpaasParameters:          MakeBGPaaServiceParametersType(),
		ParentUUID:                "",
		FQName:                    []string{},
		DisplayName:               "",
		MacAgingTime:              MakeMACAgingTime(),
	}
}

// MakeGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
	return []*GlobalSystemConfig{}
}
