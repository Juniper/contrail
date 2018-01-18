package models

// GlobalSystemConfig

import "encoding/json"

// GlobalSystemConfig
type GlobalSystemConfig struct {
	UserDefinedLogStatistics  *UserDefinedLogStatList        `json:"user_defined_log_statistics,omitempty"`
	MacLimitControl           *MACLimitControlType           `json:"mac_limit_control,omitempty"`
	Annotations               *KeyValuePairs                 `json:"annotations,omitempty"`
	ParentType                string                         `json:"parent_type,omitempty"`
	FQName                    []string                       `json:"fq_name,omitempty"`
	AlarmEnable               bool                           `json:"alarm_enable"`
	IbgpAutoMesh              bool                           `json:"ibgp_auto_mesh"`
	IDPerms                   *IdPermsType                   `json:"id_perms,omitempty"`
	UUID                      string                         `json:"uuid,omitempty"`
	ConfigVersion             string                         `json:"config_version,omitempty"`
	PluginTuning              *PluginProperties              `json:"plugin_tuning,omitempty"`
	MacAgingTime              MACAgingTime                   `json:"mac_aging_time,omitempty"`
	AutonomousSystem          AutonomousSystemType           `json:"autonomous_system,omitempty"`
	MacMoveControl            *MACMoveLimitControlType       `json:"mac_move_control,omitempty"`
	BGPAlwaysCompareMed       bool                           `json:"bgp_always_compare_med"`
	GracefulRestartParameters *GracefulRestartParametersType `json:"graceful_restart_parameters,omitempty"`
	IPFabricSubnets           *SubnetListType                `json:"ip_fabric_subnets,omitempty"`
	DisplayName               string                         `json:"display_name,omitempty"`
	Perms2                    *PermType2                     `json:"perms2,omitempty"`
	ParentUUID                string                         `json:"parent_uuid,omitempty"`
	BgpaasParameters          *BGPaaServiceParametersType    `json:"bgpaas_parameters,omitempty"`

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
		BgpaasParameters:          MakeBGPaaServiceParametersType(),
		BGPAlwaysCompareMed:       false,
		GracefulRestartParameters: MakeGracefulRestartParametersType(),
		IPFabricSubnets:           MakeSubnetListType(),
		DisplayName:               "",
		Perms2:                    MakePermType2(),
		ParentUUID:                "",
		AlarmEnable:               false,
		UserDefinedLogStatistics:  MakeUserDefinedLogStatList(),
		MacLimitControl:           MakeMACLimitControlType(),
		Annotations:               MakeKeyValuePairs(),
		ParentType:                "",
		FQName:                    []string{},
		ConfigVersion:             "",
		IbgpAutoMesh:              false,
		IDPerms:                   MakeIdPermsType(),
		UUID:                      "",
		MacMoveControl:            MakeMACMoveLimitControlType(),
		PluginTuning:              MakePluginProperties(),
		MacAgingTime:              MakeMACAgingTime(),
		AutonomousSystem:          MakeAutonomousSystemType(),
	}
}

// MakeGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
	return []*GlobalSystemConfig{}
}
