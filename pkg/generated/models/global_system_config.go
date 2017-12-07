package models

// GlobalSystemConfig

import "encoding/json"

// GlobalSystemConfig
type GlobalSystemConfig struct {
	ConfigVersion             string                         `json:"config_version"`
	UUID                      string                         `json:"uuid"`
	MacMoveControl            *MACMoveLimitControlType       `json:"mac_move_control"`
	IDPerms                   *IdPermsType                   `json:"id_perms"`
	AutonomousSystem          AutonomousSystemType           `json:"autonomous_system"`
	MacLimitControl           *MACLimitControlType           `json:"mac_limit_control"`
	DisplayName               string                         `json:"display_name"`
	AlarmEnable               bool                           `json:"alarm_enable"`
	MacAgingTime              MACAgingTime                   `json:"mac_aging_time"`
	UserDefinedLogStatistics  *UserDefinedLogStatList        `json:"user_defined_log_statistics"`
	IPFabricSubnets           *SubnetListType                `json:"ip_fabric_subnets"`
	GracefulRestartParameters *GracefulRestartParametersType `json:"graceful_restart_parameters"`
	Annotations               *KeyValuePairs                 `json:"annotations"`
	Perms2                    *PermType2                     `json:"perms2"`
	FQName                    []string                       `json:"fq_name"`
	BgpaasParameters          *BGPaaServiceParametersType    `json:"bgpaas_parameters"`
	PluginTuning              *PluginProperties              `json:"plugin_tuning"`
	IbgpAutoMesh              bool                           `json:"ibgp_auto_mesh"`
	BGPAlwaysCompareMed       bool                           `json:"bgp_always_compare_med"`

	// bgp_router <common.Reference Value>
	BGPRouterRefs []*GlobalSystemConfigBGPRouterRef `json:"bgp_router_refs"`

	ConfigRoots []*GlobalSystemConfigConfigRoot
}

// GlobalSystemConfigBGPRouterRef references each other
type GlobalSystemConfigBGPRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// GlobalSystemConfig parents relation object

type GlobalSystemConfigConfigRoot struct {
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
		IDPerms:                   MakeIdPermsType(),
		MacMoveControl:            MakeMACMoveLimitControlType(),
		MacAgingTime:              MakeMACAgingTime(),
		UserDefinedLogStatistics:  MakeUserDefinedLogStatList(),
		IPFabricSubnets:           MakeSubnetListType(),
		AutonomousSystem:          MakeAutonomousSystemType(),
		MacLimitControl:           MakeMACLimitControlType(),
		DisplayName:               "",
		AlarmEnable:               false,
		PluginTuning:              MakePluginProperties(),
		IbgpAutoMesh:              false,
		BGPAlwaysCompareMed:       false,
		GracefulRestartParameters: MakeGracefulRestartParametersType(),
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		FQName:                    []string{},
		BgpaasParameters:          MakeBGPaaServiceParametersType(),
		UUID:                      "",
		ConfigVersion:             "",
	}
}

// InterfaceToGlobalSystemConfig makes GlobalSystemConfig from interface
func InterfaceToGlobalSystemConfig(iData interface{}) *GlobalSystemConfig {
	data := iData.(map[string]interface{})
	return &GlobalSystemConfig{
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair","GoPremitive":false},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs","GoPremitive":false}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"perms2_owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"perms2_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string","GoPremitive":true},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType","GoPremitive":false},"GoName":"Share","GoType":"[]*ShareType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2","GoPremitive":false}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string","GoPremitive":true},"GoName":"FQName","GoType":"[]string","GoPremitive":true}
		BgpaasParameters: InterfaceToBGPaaServiceParametersType(data["bgpaas_parameters"]),

		//{"Title":"","Description":"BGP As A Service Parameters configuration","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"port_end":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"port_end","Item":null,"GoName":"PortEnd","GoType":"L4PortType","GoPremitive":false},"port_start":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"port_start","Item":null,"GoName":"PortStart","GoType":"L4PortType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/BGPaaServiceParametersType","CollectionType":"","Column":"","Item":null,"GoName":"BgpaasParameters","GoType":"BGPaaServiceParametersType","GoPremitive":false}
		PluginTuning: InterfaceToPluginProperties(data["plugin_tuning"]),

		//{"Title":"","Description":"Various Orchestration system plugin(interface) parameters, like Openstack Neutron plugin.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"plugin_property":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"plugin_property","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"property":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Property","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PluginProperty","CollectionType":"","Column":"","Item":null,"GoName":"PluginProperty","GoType":"PluginProperty","GoPremitive":false},"GoName":"PluginProperty","GoType":"[]*PluginProperty","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PluginProperties","CollectionType":"","Column":"","Item":null,"GoName":"PluginTuning","GoType":"PluginProperties","GoPremitive":false}
		IbgpAutoMesh: data["ibgp_auto_mesh"].(bool),

		//{"Title":"","Description":"When true, system will automatically create BGP peering mesh with all control-nodes that have same BGP AS number as global AS number.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ibgp_auto_mesh","Item":null,"GoName":"IbgpAutoMesh","GoType":"bool","GoPremitive":true}
		BGPAlwaysCompareMed: data["bgp_always_compare_med"].(bool),

		//{"Title":"","Description":"Always compare MED even if paths are received from different ASes.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"bgp_always_compare_med","Item":null,"GoName":"BGPAlwaysCompareMed","GoType":"bool","GoPremitive":true}
		GracefulRestartParameters: InterfaceToGracefulRestartParametersType(data["graceful_restart_parameters"]),

		//{"Title":"","Description":"Graceful Restart parameters","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"bgp_helper_enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"bgp_helper_enable","Item":null,"GoName":"BGPHelperEnable","GoType":"bool","GoPremitive":true},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"graceful_restart_parameters_enable","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true},"end_of_rib_timeout":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":4095,"Ref":"types.json#/definitions/EndOfRibTimeType","CollectionType":"","Column":"end_of_rib_timeout","Item":null,"GoName":"EndOfRibTimeout","GoType":"EndOfRibTimeType","GoPremitive":false},"long_lived_restart_time":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":16777215,"Ref":"types.json#/definitions/LongLivedGracefulRestartTimeType","CollectionType":"","Column":"long_lived_restart_time","Item":null,"GoName":"LongLivedRestartTime","GoType":"LongLivedGracefulRestartTimeType","GoPremitive":false},"restart_time":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":4095,"Ref":"types.json#/definitions/GracefulRestartTimeType","CollectionType":"","Column":"restart_time","Item":null,"GoName":"RestartTime","GoType":"GracefulRestartTimeType","GoPremitive":false},"xmpp_helper_enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"xmpp_helper_enable","Item":null,"GoName":"XMPPHelperEnable","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/GracefulRestartParametersType","CollectionType":"","Column":"","Item":null,"GoName":"GracefulRestartParameters","GoType":"GracefulRestartParametersType","GoPremitive":false}
		ConfigVersion: data["config_version"].(string),

		//{"Title":"","Description":"Version of OpenContrail software that generated this config.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"config_version","Item":null,"GoName":"ConfigVersion","GoType":"string","GoPremitive":true}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string","GoPremitive":true}
		MacMoveControl: InterfaceToMACMoveLimitControlType(data["mac_move_control"]),

		//{"Title":"","Description":"MAC move control on the network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"mac_move_limit":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mac_move_limit","Item":null,"GoName":"MacMoveLimit","GoType":"int","GoPremitive":true},"mac_move_limit_action":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["log","alarm","shutdown","drop"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitExceedActionType","CollectionType":"","Column":"mac_move_limit_action","Item":null,"GoName":"MacMoveLimitAction","GoType":"MACLimitExceedActionType","GoPremitive":false},"mac_move_time_window":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":60,"Ref":"types.json#/definitions/MACMoveTimeWindow","CollectionType":"","Column":"mac_move_time_window","Item":null,"GoName":"MacMoveTimeWindow","GoType":"MACMoveTimeWindow","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACMoveLimitControlType","CollectionType":"","Column":"","Item":null,"GoName":"MacMoveControl","GoType":"MACMoveLimitControlType","GoPremitive":false}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string","GoPremitive":true},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string","GoPremitive":true},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string","GoPremitive":true},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string","GoPremitive":true},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string","GoPremitive":true},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType","GoPremitive":false},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType","GoPremitive":false},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType","GoPremitive":false}
		MacLimitControl: InterfaceToMACLimitControlType(data["mac_limit_control"]),

		//{"Title":"","Description":"MAC limit control on the network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"mac_limit":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mac_limit","Item":null,"GoName":"MacLimit","GoType":"int","GoPremitive":true},"mac_limit_action":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["log","alarm","shutdown","drop"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitExceedActionType","CollectionType":"","Column":"mac_limit_action","Item":null,"GoName":"MacLimitAction","GoType":"MACLimitExceedActionType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitControlType","CollectionType":"","Column":"","Item":null,"GoName":"MacLimitControl","GoType":"MACLimitControlType","GoPremitive":false}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string","GoPremitive":true}
		AlarmEnable: data["alarm_enable"].(bool),

		//{"Title":"","Description":"Flag to enable/disable alarms configured under global-system-config. True, if not set.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"alarm_enable","Item":null,"GoName":"AlarmEnable","GoType":"bool","GoPremitive":true}
		MacAgingTime: InterfaceToMACAgingTime(data["mac_aging_time"]),

		//{"Title":"","Description":"MAC aging time on the network","SQL":"int","Default":"300","Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":86400,"Ref":"types.json#/definitions/MACAgingTime","CollectionType":"","Column":"mac_aging_time","Item":null,"GoName":"MacAgingTime","GoType":"MACAgingTime","GoPremitive":false}
		UserDefinedLogStatistics: InterfaceToUserDefinedLogStatList(data["user_defined_log_statistics"]),

		//{"Title":"","Description":"stats name and patterns","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"statlist":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Name","GoType":"string","GoPremitive":true},"pattern":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Pattern","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/UserDefinedLogStat","CollectionType":"","Column":"","Item":null,"GoName":"Statlist","GoType":"UserDefinedLogStat","GoPremitive":false},"GoName":"Statlist","GoType":"[]*UserDefinedLogStat","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/UserDefinedLogStatList","CollectionType":"map","Column":"user_defined_log_statistics","Item":null,"GoName":"UserDefinedLogStatistics","GoType":"UserDefinedLogStatList","GoPremitive":false}
		IPFabricSubnets: InterfaceToSubnetListType(data["ip_fabric_subnets"]),

		//{"Title":"","Description":"List of all subnets in which vrouter ip address exist. Used by Device manager to configure dynamic GRE tunnels on the SDN gateway.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"subnet":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"subnet","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType","GoPremitive":false},"GoName":"Subnet","GoType":"[]*SubnetType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetListType","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricSubnets","GoType":"SubnetListType","GoPremitive":false}
		AutonomousSystem: InterfaceToAutonomousSystemType(data["autonomous_system"]),

		//{"Title":"","Description":"16 bit BGP Autonomous System number for the cluster.","SQL":"int","Default":null,"Operation":"","Presence":"required","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":65534,"Ref":"types.json#/definitions/AutonomousSystemType","CollectionType":"","Column":"autonomous_system","Item":null,"GoName":"AutonomousSystem","GoType":"AutonomousSystemType","GoPremitive":false}

	}
}

// InterfaceToGlobalSystemConfigSlice makes a slice of GlobalSystemConfig from interface
func InterfaceToGlobalSystemConfigSlice(data interface{}) []*GlobalSystemConfig {
	list := data.([]interface{})
	result := MakeGlobalSystemConfigSlice()
	for _, item := range list {
		result = append(result, InterfaceToGlobalSystemConfig(item))
	}
	return result
}

// MakeGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
	return []*GlobalSystemConfig{}
}
