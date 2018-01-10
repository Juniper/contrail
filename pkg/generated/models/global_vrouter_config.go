package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode"`
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list"`
	IDPerms                    *IdPermsType                   `json:"id_perms"`
	DisplayName                string                         `json:"display_name"`
	Annotations                *KeyValuePairs                 `json:"annotations"`
	ParentUUID                 string                         `json:"parent_uuid"`
	ParentType                 string                         `json:"parent_type"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode"`
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services"`
	FQName                     []string                       `json:"fq_name"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`
	Perms2                     *PermType2                     `json:"perms2"`
	UUID                       string                         `json:"uuid"`
	FlowExportRate             int                            `json:"flow_export_rate"`

	SecurityLoggingObjects []*SecurityLoggingObject `json:"security_logging_objects"`
}

// String returns json representation of the object
func (model *GlobalVrouterConfig) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeGlobalVrouterConfig makes GlobalVrouterConfig
func MakeGlobalVrouterConfig() *GlobalVrouterConfig {
	return &GlobalVrouterConfig{
		//TODO(nati): Apply default
		UUID:                       "",
		FlowExportRate:             0,
		EnableSecurityLogging:      false,
		Perms2:                     MakePermType2(),
		FlowAgingTimeoutList:       MakeFlowAgingTimeoutList(),
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		Annotations:                MakeKeyValuePairs(),
		ParentUUID:                 "",
		ParentType:                 "",
		EncapsulationPriorities:    MakeEncapsulationPrioritiesType(),
		IDPerms:                    MakeIdPermsType(),
		DisplayName:                "",
		FQName:                     []string{},
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
		ForwardingMode:           MakeForwardingModeType(),
		LinklocalServices:        MakeLinklocalServicesTypes(),
	}
}

// InterfaceToGlobalVrouterConfig makes GlobalVrouterConfig from interface
func InterfaceToGlobalVrouterConfig(iData interface{}) *GlobalVrouterConfig {
	data := iData.(map[string]interface{})
	return &GlobalVrouterConfig{
		FlowAgingTimeoutList: InterfaceToFlowAgingTimeoutList(data["flow_aging_timeout_list"]),

		//{"description":"Flow aging timeout per application (protocol, port) list.","type":"object","properties":{"flow_aging_timeout":{"type":"array","item":{"type":"object","properties":{"port":{"type":"integer"},"protocol":{"type":"string"},"timeout_in_seconds":{"type":"integer"}}}}}}
		VxlanNetworkIdentifierMode: InterfaceToVxlanNetworkIdentifierModeType(data["vxlan_network_identifier_mode"]),

		//{"type":"string","enum":["configured","automatic"]}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		EncapsulationPriorities: InterfaceToEncapsulationPrioritiesType(data["encapsulation_priorities"]),

		//{"description":"Ordered list of encapsulations that vrouter will use in priority order.","type":"object","properties":{"encapsulation":{"type":"array"}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		LinklocalServices: InterfaceToLinklocalServicesTypes(data["linklocal_services"]),

		//{"description":"Global services provided on link local subnet to the virtual machines.","type":"object","properties":{"linklocal_service_entry":{"type":"array","item":{"type":"object","properties":{"ip_fabric_DNS_service_name":{"type":"string"},"ip_fabric_service_ip":{"type":"array","item":{"type":"string"}},"ip_fabric_service_port":{"type":"integer"},"linklocal_service_ip":{"type":"string"},"linklocal_service_name":{"type":"string"},"linklocal_service_port":{"type":"integer"}}}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		EcmpHashingIncludeFields: InterfaceToEcmpHashingIncludeFields(data["ecmp_hashing_include_fields"]),

		//{"description":"ECMP hashing config at global level.","type":"object","properties":{"destination_ip":{"type":"boolean"},"destination_port":{"type":"boolean"},"hashing_configured":{"type":"boolean"},"ip_protocol":{"type":"boolean"},"source_ip":{"type":"boolean"},"source_port":{"type":"boolean"}}}
		ForwardingMode: InterfaceToForwardingModeType(data["forwarding_mode"]),

		//{"description":"Packet forwarding mode for this system L2-only, L3-only OR L2-L3. L2-L3 is default.","type":"string","enum":["l2_l3","l2","l3"]}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		FlowExportRate: data["flow_export_rate"].(int),

		//{"description":"Flow export rate is global config, rate at which each vrouter will sample and export flow records to analytics","type":"integer"}
		EnableSecurityLogging: data["enable_security_logging"].(bool),

		//{"description":"Enable or disable security-logging in the system","default":true,"type":"boolean"}

	}
}

// InterfaceToGlobalVrouterConfigSlice makes a slice of GlobalVrouterConfig from interface
func InterfaceToGlobalVrouterConfigSlice(data interface{}) []*GlobalVrouterConfig {
	list := data.([]interface{})
	result := MakeGlobalVrouterConfigSlice()
	for _, item := range list {
		result = append(result, InterfaceToGlobalVrouterConfig(item))
	}
	return result
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
