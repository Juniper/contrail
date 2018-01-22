package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list,omitempty"`
	ParentUUID                 string                         `json:"parent_uuid,omitempty"`
	FlowExportRate             int                            `json:"flow_export_rate,omitempty"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities,omitempty"`
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
	Annotations                *KeyValuePairs                 `json:"annotations,omitempty"`
	Perms2                     *PermType2                     `json:"perms2,omitempty"`
	DisplayName                string                         `json:"display_name,omitempty"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields,omitempty"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode,omitempty"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`
	UUID                       string                         `json:"uuid,omitempty"`
	FQName                     []string                       `json:"fq_name,omitempty"`
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services,omitempty"`
	ParentType                 string                         `json:"parent_type,omitempty"`
	IDPerms                    *IdPermsType                   `json:"id_perms,omitempty"`

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
		DisplayName:                "",
		FlowExportRate:             0,
		EncapsulationPriorities:    MakeEncapsulationPrioritiesType(),
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		Annotations:                MakeKeyValuePairs(),
		Perms2:                     MakePermType2(),
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
		ForwardingMode:           MakeForwardingModeType(),
		EnableSecurityLogging:    false,
		UUID:                 "",
		FQName:               []string{},
		LinklocalServices:    MakeLinklocalServicesTypes(),
		ParentType:           "",
		IDPerms:              MakeIdPermsType(),
		FlowAgingTimeoutList: MakeFlowAgingTimeoutList(),
		ParentUUID:           "",
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
