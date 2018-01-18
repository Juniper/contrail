package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services,omitempty"`
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
	FQName                     []string                       `json:"fq_name,omitempty"`
	Perms2                     *PermType2                     `json:"perms2,omitempty"`
	UUID                       string                         `json:"uuid,omitempty"`
	ParentUUID                 string                         `json:"parent_uuid,omitempty"`
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list,omitempty"`
	IDPerms                    *IdPermsType                   `json:"id_perms,omitempty"`
	Annotations                *KeyValuePairs                 `json:"annotations,omitempty"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields,omitempty"`
	FlowExportRate             int                            `json:"flow_export_rate,omitempty"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities,omitempty"`
	DisplayName                string                         `json:"display_name,omitempty"`
	ParentType                 string                         `json:"parent_type,omitempty"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode,omitempty"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`

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
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		FQName:                   []string{},
		Perms2:                   MakePermType2(),
		UUID:                     "",
		ParentUUID:               "",
		LinklocalServices:        MakeLinklocalServicesTypes(),
		IDPerms:                  MakeIdPermsType(),
		Annotations:              MakeKeyValuePairs(),
		FlowAgingTimeoutList:     MakeFlowAgingTimeoutList(),
		FlowExportRate:           0,
		EncapsulationPriorities:  MakeEncapsulationPrioritiesType(),
		DisplayName:              "",
		ParentType:               "",
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
		EnableSecurityLogging:    false,
		ForwardingMode:           MakeForwardingModeType(),
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
