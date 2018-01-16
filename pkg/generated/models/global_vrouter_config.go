package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	IDPerms                    *IdPermsType                   `json:"id_perms,omitempty"`
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list,omitempty"`
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services,omitempty"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities,omitempty"`
	ParentType                 string                         `json:"parent_type,omitempty"`
	FQName                     []string                       `json:"fq_name,omitempty"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging,omitempty"`
	ParentUUID                 string                         `json:"parent_uuid,omitempty"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode,omitempty"`
	FlowExportRate             int                            `json:"flow_export_rate,omitempty"`
	Annotations                *KeyValuePairs                 `json:"annotations,omitempty"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields,omitempty"`
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
	Perms2                     *PermType2                     `json:"perms2,omitempty"`
	UUID                       string                         `json:"uuid,omitempty"`
	DisplayName                string                         `json:"display_name,omitempty"`

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
		EnableSecurityLogging:      false,
		ParentUUID:                 "",
		ForwardingMode:             MakeForwardingModeType(),
		FlowExportRate:             0,
		Annotations:                MakeKeyValuePairs(),
		DisplayName:                "",
		EcmpHashingIncludeFields:   MakeEcmpHashingIncludeFields(),
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		Perms2:                  MakePermType2(),
		UUID:                    "",
		FQName:                  []string{},
		IDPerms:                 MakeIdPermsType(),
		FlowAgingTimeoutList:    MakeFlowAgingTimeoutList(),
		LinklocalServices:       MakeLinklocalServicesTypes(),
		EncapsulationPriorities: MakeEncapsulationPrioritiesType(),
		ParentType:              "",
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
