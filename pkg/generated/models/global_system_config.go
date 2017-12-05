package models

// GlobalSystemConfig

import "encoding/json"

type GlobalSystemConfig struct {
	BgpaasParameters          *BGPaaServiceParametersType    `json:"bgpaas_parameters"`
	MacAgingTime              MACAgingTime                   `json:"mac_aging_time"`
	GracefulRestartParameters *GracefulRestartParametersType `json:"graceful_restart_parameters"`
	AutonomousSystem          AutonomousSystemType           `json:"autonomous_system"`
	UUID                      string                         `json:"uuid"`
	FQName                    []string                       `json:"fq_name"`
	AlarmEnable               bool                           `json:"alarm_enable"`
	IbgpAutoMesh              bool                           `json:"ibgp_auto_mesh"`
	BGPAlwaysCompareMed       bool                           `json:"bgp_always_compare_med"`
	MacLimitControl           *MACLimitControlType           `json:"mac_limit_control"`
	IDPerms                   *IdPermsType                   `json:"id_perms"`
	ConfigVersion             string                         `json:"config_version"`
	MacMoveControl            *MACMoveLimitControlType       `json:"mac_move_control"`
	UserDefinedLogStatistics  *UserDefinedLogStatList        `json:"user_defined_log_statistics"`
	DisplayName               string                         `json:"display_name"`
	Perms2                    *PermType2                     `json:"perms2"`
	PluginTuning              *PluginProperties              `json:"plugin_tuning"`
	IPFabricSubnets           *SubnetListType                `json:"ip_fabric_subnets"`
	Annotations               *KeyValuePairs                 `json:"annotations"`

	// bgp_router <utils.Reference Value>
	BGPRouterRefs []*GlobalSystemConfigBGPRouterRef

	ConfigRoots []*GlobalSystemConfigConfigRoot
}

// <utils.Reference Value>
type GlobalSystemConfigBGPRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

type GlobalSystemConfigConfigRoot struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

func (model *GlobalSystemConfig) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeGlobalSystemConfig() *GlobalSystemConfig {
	return &GlobalSystemConfig{
		//TODO(nati): Apply default
		BgpaasParameters:          MakeBGPaaServiceParametersType(),
		MacAgingTime:              MakeMACAgingTime(),
		GracefulRestartParameters: MakeGracefulRestartParametersType(),
		AutonomousSystem:          MakeAutonomousSystemType(),
		UUID:                      "",
		AlarmEnable:               false,
		IbgpAutoMesh:              false,
		BGPAlwaysCompareMed:       false,
		MacLimitControl:           MakeMACLimitControlType(),
		IDPerms:                   MakeIdPermsType(),
		FQName:                    []string{},
		ConfigVersion:             "",
		MacMoveControl:            MakeMACMoveLimitControlType(),
		UserDefinedLogStatistics:  MakeUserDefinedLogStatList(),
		DisplayName:               "",
		Perms2:                    MakePermType2(),
		PluginTuning:              MakePluginProperties(),
		IPFabricSubnets:           MakeSubnetListType(),
		Annotations:               MakeKeyValuePairs(),
	}
}

func InterfaceToGlobalSystemConfig(iData interface{}) *GlobalSystemConfig {
	data := iData.(map[string]interface{})
	return &GlobalSystemConfig{
		AlarmEnable: data["alarm_enable"].(bool),

		//{"Title":"","Description":"Flag to enable/disable alarms configured under global-system-config. True, if not set.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"alarm_enable","Item":null,"GoName":"AlarmEnable","GoType":"bool"}
		IbgpAutoMesh: data["ibgp_auto_mesh"].(bool),

		//{"Title":"","Description":"When true, system will automatically create BGP peering mesh with all control-nodes that have same BGP AS number as global AS number.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ibgp_auto_mesh","Item":null,"GoName":"IbgpAutoMesh","GoType":"bool"}
		BGPAlwaysCompareMed: data["bgp_always_compare_med"].(bool),

		//{"Title":"","Description":"Always compare MED even if paths are received from different ASes.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"bgp_always_compare_med","Item":null,"GoName":"BGPAlwaysCompareMed","GoType":"bool"}
		MacLimitControl: InterfaceToMACLimitControlType(data["mac_limit_control"]),

		//{"Title":"","Description":"MAC limit control on the network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"mac_limit":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mac_limit","Item":null,"GoName":"MacLimit","GoType":"int"},"mac_limit_action":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["log","alarm","shutdown","drop"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitExceedActionType","CollectionType":"","Column":"mac_limit_action","Item":null,"GoName":"MacLimitAction","GoType":"MACLimitExceedActionType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitControlType","CollectionType":"","Column":"","Item":null,"GoName":"MacLimitControl","GoType":"MACLimitControlType"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string"},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string"},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string"},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool"},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string"},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string"},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType"},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType"},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType"}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string"},"GoName":"FQName","GoType":"[]string"}
		ConfigVersion: data["config_version"].(string),

		//{"Title":"","Description":"Version of OpenContrail software that generated this config.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"config_version","Item":null,"GoName":"ConfigVersion","GoType":"string"}
		MacMoveControl: InterfaceToMACMoveLimitControlType(data["mac_move_control"]),

		//{"Title":"","Description":"MAC move control on the network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"mac_move_limit":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mac_move_limit","Item":null,"GoName":"MacMoveLimit","GoType":"int"},"mac_move_limit_action":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["log","alarm","shutdown","drop"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitExceedActionType","CollectionType":"","Column":"mac_move_limit_action","Item":null,"GoName":"MacMoveLimitAction","GoType":"MACLimitExceedActionType"},"mac_move_time_window":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":60,"Ref":"types.json#/definitions/MACMoveTimeWindow","CollectionType":"","Column":"mac_move_time_window","Item":null,"GoName":"MacMoveTimeWindow","GoType":"MACMoveTimeWindow"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACMoveLimitControlType","CollectionType":"","Column":"","Item":null,"GoName":"MacMoveControl","GoType":"MACMoveLimitControlType"}
		UserDefinedLogStatistics: InterfaceToUserDefinedLogStatList(data["user_defined_log_statistics"]),

		//{"Title":"","Description":"stats name and patterns","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"statlist":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Name","GoType":"string"},"pattern":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Pattern","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/UserDefinedLogStat","CollectionType":"","Column":"","Item":null,"GoName":"Statlist","GoType":"UserDefinedLogStat"},"GoName":"Statlist","GoType":"[]*UserDefinedLogStat"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/UserDefinedLogStatList","CollectionType":"map","Column":"user_defined_log_statistics","Item":null,"GoName":"UserDefinedLogStatistics","GoType":"UserDefinedLogStatList"}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"perms2_owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"perms2_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string"},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType"},"GoName":"Share","GoType":"[]*ShareType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2"}
		PluginTuning: InterfaceToPluginProperties(data["plugin_tuning"]),

		//{"Title":"","Description":"Various Orchestration system plugin(interface) parameters, like Openstack Neutron plugin.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"plugin_property":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"plugin_property","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"property":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Property","GoType":"string"},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PluginProperty","CollectionType":"","Column":"","Item":null,"GoName":"PluginProperty","GoType":"PluginProperty"},"GoName":"PluginProperty","GoType":"[]*PluginProperty"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PluginProperties","CollectionType":"","Column":"","Item":null,"GoName":"PluginTuning","GoType":"PluginProperties"}
		IPFabricSubnets: InterfaceToSubnetListType(data["ip_fabric_subnets"]),

		//{"Title":"","Description":"List of all subnets in which vrouter ip address exist. Used by Device manager to configure dynamic GRE tunnels on the SDN gateway.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"subnet":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"subnet","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"GoName":"Subnet","GoType":"[]*SubnetType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetListType","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricSubnets","GoType":"SubnetListType"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string"},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair"},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs"}
		BgpaasParameters: InterfaceToBGPaaServiceParametersType(data["bgpaas_parameters"]),

		//{"Title":"","Description":"BGP As A Service Parameters configuration","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"port_end":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"port_end","Item":null,"GoName":"PortEnd","GoType":"L4PortType"},"port_start":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"port_start","Item":null,"GoName":"PortStart","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/BGPaaServiceParametersType","CollectionType":"","Column":"","Item":null,"GoName":"BgpaasParameters","GoType":"BGPaaServiceParametersType"}
		MacAgingTime: InterfaceToMACAgingTime(data["mac_aging_time"]),

		//{"Title":"","Description":"MAC aging time on the network","SQL":"int","Default":"300","Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":86400,"Ref":"types.json#/definitions/MACAgingTime","CollectionType":"","Column":"mac_aging_time","Item":null,"GoName":"MacAgingTime","GoType":"MACAgingTime"}
		GracefulRestartParameters: InterfaceToGracefulRestartParametersType(data["graceful_restart_parameters"]),

		//{"Title":"","Description":"Graceful Restart parameters","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"bgp_helper_enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"bgp_helper_enable","Item":null,"GoName":"BGPHelperEnable","GoType":"bool"},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"graceful_restart_parameters_enable","Item":null,"GoName":"Enable","GoType":"bool"},"end_of_rib_timeout":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":4095,"Ref":"types.json#/definitions/EndOfRibTimeType","CollectionType":"","Column":"end_of_rib_timeout","Item":null,"GoName":"EndOfRibTimeout","GoType":"EndOfRibTimeType"},"long_lived_restart_time":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":16777215,"Ref":"types.json#/definitions/LongLivedGracefulRestartTimeType","CollectionType":"","Column":"long_lived_restart_time","Item":null,"GoName":"LongLivedRestartTime","GoType":"LongLivedGracefulRestartTimeType"},"restart_time":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":4095,"Ref":"types.json#/definitions/GracefulRestartTimeType","CollectionType":"","Column":"restart_time","Item":null,"GoName":"RestartTime","GoType":"GracefulRestartTimeType"},"xmpp_helper_enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"xmpp_helper_enable","Item":null,"GoName":"XMPPHelperEnable","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/GracefulRestartParametersType","CollectionType":"","Column":"","Item":null,"GoName":"GracefulRestartParameters","GoType":"GracefulRestartParametersType"}
		AutonomousSystem: InterfaceToAutonomousSystemType(data["autonomous_system"]),

		//{"Title":"","Description":"16 bit BGP Autonomous System number for the cluster.","SQL":"int","Default":null,"Operation":"","Presence":"required","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":65534,"Ref":"types.json#/definitions/AutonomousSystemType","CollectionType":"","Column":"autonomous_system","Item":null,"GoName":"AutonomousSystem","GoType":"AutonomousSystemType"}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string"}

	}
}

func InterfaceToGlobalSystemConfigSlice(data interface{}) []*GlobalSystemConfig {
	list := data.([]interface{})
	result := MakeGlobalSystemConfigSlice()
	for _, item := range list {
		result = append(result, InterfaceToGlobalSystemConfig(item))
	}
	return result
}

func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
	return []*GlobalSystemConfig{}
}
