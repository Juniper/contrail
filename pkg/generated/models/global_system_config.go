package models

// GlobalSystemConfig

import "encoding/json"

// GlobalSystemConfig
type GlobalSystemConfig struct {
	MacAgingTime              MACAgingTime                   `json:"mac_aging_time,omitempty"`
	UserDefinedLogStatistics  *UserDefinedLogStatList        `json:"user_defined_log_statistics,omitempty"`
	AutonomousSystem          AutonomousSystemType           `json:"autonomous_system,omitempty"`
	ParentType                string                         `json:"parent_type,omitempty"`
	FQName                    []string                       `json:"fq_name,omitempty"`
	MacMoveControl            *MACMoveLimitControlType       `json:"mac_move_control,omitempty"`
	IbgpAutoMesh              bool                           `json:"ibgp_auto_mesh"`
	IDPerms                   *IdPermsType                   `json:"id_perms,omitempty"`
	DisplayName               string                         `json:"display_name,omitempty"`
	AlarmEnable               bool                           `json:"alarm_enable"`
	GracefulRestartParameters *GracefulRestartParametersType `json:"graceful_restart_parameters,omitempty"`
	BGPAlwaysCompareMed       bool                           `json:"bgp_always_compare_med"`
	IPFabricSubnets           *SubnetListType                `json:"ip_fabric_subnets,omitempty"`
	MacLimitControl           *MACLimitControlType           `json:"mac_limit_control,omitempty"`
	ParentUUID                string                         `json:"parent_uuid,omitempty"`
	BgpaasParameters          *BGPaaServiceParametersType    `json:"bgpaas_parameters,omitempty"`
	PluginTuning              *PluginProperties              `json:"plugin_tuning,omitempty"`
	UUID                      string                         `json:"uuid,omitempty"`
	Annotations               *KeyValuePairs                 `json:"annotations,omitempty"`
	ConfigVersion             string                         `json:"config_version,omitempty"`
	Perms2                    *PermType2                     `json:"perms2,omitempty"`

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
		GracefulRestartParameters: MakeGracefulRestartParametersType(),
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		BgpaasParameters:          MakeBGPaaServiceParametersType(),
		PluginTuning:              MakePluginProperties(),
		BGPAlwaysCompareMed:       false,
		IPFabricSubnets:           MakeSubnetListType(),
		MacLimitControl:           MakeMACLimitControlType(),
		ParentUUID:                "",
		ConfigVersion:             "",
		Perms2:                    MakePermType2(),
		UUID:                      "",
		Annotations:               MakeKeyValuePairs(),
		MacMoveControl:            MakeMACMoveLimitControlType(),
		IbgpAutoMesh:              false,
		MacAgingTime:              MakeMACAgingTime(),
		UserDefinedLogStatistics:  MakeUserDefinedLogStatList(),
		AutonomousSystem:          MakeAutonomousSystemType(),
		ParentType:                "",
		FQName:                    []string{},
	}
}

// MakeGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
	return []*GlobalSystemConfig{}
}
