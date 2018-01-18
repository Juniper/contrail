package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	Perms2                     *PermType2                     `json:"perms2,omitempty"`
	UUID                       string                         `json:"uuid,omitempty"`
	ParentUUID                 string                         `json:"parent_uuid,omitempty"`
	IDPerms                    *IdPermsType                   `json:"id_perms,omitempty"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities,omitempty"`
	FlowExportRate             int                            `json:"flow_export_rate,omitempty"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`
	DisplayName                string                         `json:"display_name,omitempty"`
	Annotations                *KeyValuePairs                 `json:"annotations,omitempty"`
	FQName                     []string                       `json:"fq_name,omitempty"`
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list,omitempty"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode,omitempty"`
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services,omitempty"`
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
	ParentType                 string                         `json:"parent_type,omitempty"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields,omitempty"`

	SecurityLoggingObjects []*SecurityLoggingObject `json:"security_logging_objects,omitempty"`
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
		EncapsulationPriorities: MakeEncapsulationPrioritiesType(),
		Perms2:                  MakePermType2(),
		UUID:                    "",
		ParentUUID:              "",
		IDPerms:                 MakeIdPermsType(),
		FlowAgingTimeoutList:    MakeFlowAgingTimeoutList(),
		FlowExportRate:          0,
		EnableSecurityLogging:   false,
		DisplayName:             "",
		Annotations:             MakeKeyValuePairs(),
		FQName:                  []string{},
		EcmpHashingIncludeFields:   MakeEcmpHashingIncludeFields(),
		ForwardingMode:             MakeForwardingModeType(),
		LinklocalServices:          MakeLinklocalServicesTypes(),
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		ParentType:                 "",
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
