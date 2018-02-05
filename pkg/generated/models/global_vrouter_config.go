package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services,omitempty"`
	ParentType                 string                         `json:"parent_type,omitempty"`
	DisplayName                string                         `json:"display_name,omitempty"`
	Annotations                *KeyValuePairs                 `json:"annotations,omitempty"`
	Perms2                     *PermType2                     `json:"perms2,omitempty"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode,omitempty"`
	FlowExportRate             int                            `json:"flow_export_rate,omitempty"`
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list,omitempty"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`
	FQName                     []string                       `json:"fq_name,omitempty"`
	IDPerms                    *IdPermsType                   `json:"id_perms,omitempty"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities,omitempty"`
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
	ParentUUID                 string                         `json:"parent_uuid,omitempty"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields,omitempty"`
	UUID                       string                         `json:"uuid,omitempty"`

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
		FlowAgingTimeoutList:       MakeFlowAgingTimeoutList(),
		EnableSecurityLogging:      false,
		FQName:                     []string{},
		IDPerms:                    MakeIdPermsType(),
		EncapsulationPriorities:    MakeEncapsulationPrioritiesType(),
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		ParentUUID:                 "",
		EcmpHashingIncludeFields:   MakeEcmpHashingIncludeFields(),
		UUID:              "",
		LinklocalServices: MakeLinklocalServicesTypes(),
		ParentType:        "",
		DisplayName:       "",
		Annotations:       MakeKeyValuePairs(),
		Perms2:            MakePermType2(),
		ForwardingMode:    MakeForwardingModeType(),
		FlowExportRate:    0,
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
