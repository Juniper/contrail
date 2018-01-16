package models

// GlobalSystemConfig

import "encoding/json"

// GlobalSystemConfig
type GlobalSystemConfig struct {
	BGPAlwaysCompareMed       bool                           `json:"bgp_always_compare_med,omitempty"`
	IPFabricSubnets           *SubnetListType                `json:"ip_fabric_subnets,omitempty"`
	ParentType                string                         `json:"parent_type,omitempty"`
	DisplayName               string                         `json:"display_name,omitempty"`
	BgpaasParameters          *BGPaaServiceParametersType    `json:"bgpaas_parameters,omitempty"`
	Annotations               *KeyValuePairs                 `json:"annotations,omitempty"`
	ParentUUID                string                         `json:"parent_uuid,omitempty"`
	PluginTuning              *PluginProperties              `json:"plugin_tuning,omitempty"`
	AlarmEnable               bool                           `json:"alarm_enable,omitempty"`
	UserDefinedLogStatistics  *UserDefinedLogStatList        `json:"user_defined_log_statistics,omitempty"`
	MacLimitControl           *MACLimitControlType           `json:"mac_limit_control,omitempty"`
	FQName                    []string                       `json:"fq_name,omitempty"`
	Perms2                    *PermType2                     `json:"perms2,omitempty"`
	UUID                      string                         `json:"uuid,omitempty"`
	ConfigVersion             string                         `json:"config_version,omitempty"`
	IbgpAutoMesh              bool                           `json:"ibgp_auto_mesh,omitempty"`
	MacAgingTime              MACAgingTime                   `json:"mac_aging_time,omitempty"`
	GracefulRestartParameters *GracefulRestartParametersType `json:"graceful_restart_parameters,omitempty"`
	AutonomousSystem          AutonomousSystemType           `json:"autonomous_system,omitempty"`
	IDPerms                   *IdPermsType                   `json:"id_perms,omitempty"`
	MacMoveControl            *MACMoveLimitControlType       `json:"mac_move_control,omitempty"`

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
		IbgpAutoMesh:              false,
		MacAgingTime:              MakeMACAgingTime(),
		GracefulRestartParameters: MakeGracefulRestartParametersType(),
		AutonomousSystem:          MakeAutonomousSystemType(),
		IDPerms:                   MakeIdPermsType(),
		MacMoveControl:            MakeMACMoveLimitControlType(),
		BGPAlwaysCompareMed:       false,
		IPFabricSubnets:           MakeSubnetListType(),
		ParentType:                "",
		DisplayName:               "",
		BgpaasParameters:          MakeBGPaaServiceParametersType(),
		Annotations:               MakeKeyValuePairs(),
		ParentUUID:                "",
		PluginTuning:              MakePluginProperties(),
		AlarmEnable:               false,
		UserDefinedLogStatistics:  MakeUserDefinedLogStatList(),
		MacLimitControl:           MakeMACLimitControlType(),
		FQName:                    []string{},
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ConfigVersion:             "",
	}
}

// MakeGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
	return []*GlobalSystemConfig{}
}
