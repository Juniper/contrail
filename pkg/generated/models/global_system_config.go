package models

// GlobalSystemConfig

import "encoding/json"

// GlobalSystemConfig
type GlobalSystemConfig struct {
	MacLimitControl           *MACLimitControlType           `json:"mac_limit_control,omitempty"`
	Perms2                    *PermType2                     `json:"perms2,omitempty"`
	UUID                      string                         `json:"uuid,omitempty"`
	ConfigVersion             string                         `json:"config_version,omitempty"`
	BgpaasParameters          *BGPaaServiceParametersType    `json:"bgpaas_parameters,omitempty"`
	PluginTuning              *PluginProperties              `json:"plugin_tuning,omitempty"`
	GracefulRestartParameters *GracefulRestartParametersType `json:"graceful_restart_parameters,omitempty"`
	IPFabricSubnets           *SubnetListType                `json:"ip_fabric_subnets,omitempty"`
	UserDefinedLogStatistics  *UserDefinedLogStatList        `json:"user_defined_log_statistics,omitempty"`
	DisplayName               string                         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs                 `json:"annotations,omitempty"`
	IDPerms                   *IdPermsType                   `json:"id_perms,omitempty"`
	MacMoveControl            *MACMoveLimitControlType       `json:"mac_move_control,omitempty"`
	IbgpAutoMesh              bool                           `json:"ibgp_auto_mesh"`
	MacAgingTime              MACAgingTime                   `json:"mac_aging_time,omitempty"`
	AutonomousSystem          AutonomousSystemType           `json:"autonomous_system,omitempty"`
	ParentUUID                string                         `json:"parent_uuid,omitempty"`
	AlarmEnable               bool                           `json:"alarm_enable"`
	BGPAlwaysCompareMed       bool                           `json:"bgp_always_compare_med"`
	ParentType                string                         `json:"parent_type,omitempty"`
	FQName                    []string                       `json:"fq_name,omitempty"`

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
		UserDefinedLogStatistics:  MakeUserDefinedLogStatList(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		MacMoveControl:            MakeMACMoveLimitControlType(),
		IbgpAutoMesh:              false,
		MacAgingTime:              MakeMACAgingTime(),
		AutonomousSystem:          MakeAutonomousSystemType(),
		ParentUUID:                "",
		IDPerms:                   MakeIdPermsType(),
		AlarmEnable:               false,
		BGPAlwaysCompareMed:       false,
		ParentType:                "",
		FQName:                    []string{},
		UUID:                      "",
		ConfigVersion:             "",
		BgpaasParameters:          MakeBGPaaServiceParametersType(),
		PluginTuning:              MakePluginProperties(),
		GracefulRestartParameters: MakeGracefulRestartParametersType(),
		IPFabricSubnets:           MakeSubnetListType(),
		MacLimitControl:           MakeMACLimitControlType(),
		Perms2:                    MakePermType2(),
	}
}

// MakeGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
	return []*GlobalSystemConfig{}
}
