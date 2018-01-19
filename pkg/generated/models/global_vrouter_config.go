package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list,omitempty"`
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
	Perms2                     *PermType2                     `json:"perms2,omitempty"`
	IDPerms                    *IdPermsType                   `json:"id_perms,omitempty"`
	DisplayName                string                         `json:"display_name,omitempty"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields,omitempty"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode,omitempty"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities,omitempty"`
	FQName                     []string                       `json:"fq_name,omitempty"`
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services,omitempty"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`
	ParentUUID                 string                         `json:"parent_uuid,omitempty"`
	ParentType                 string                         `json:"parent_type,omitempty"`
	FlowExportRate             int                            `json:"flow_export_rate,omitempty"`
	Annotations                *KeyValuePairs                 `json:"annotations,omitempty"`
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
		FlowExportRate:             0,
		Annotations:                MakeKeyValuePairs(),
		UUID:                       "",
		FlowAgingTimeoutList:       MakeFlowAgingTimeoutList(),
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		Perms2: MakePermType2(),
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
		ForwardingMode:           MakeForwardingModeType(),
		EncapsulationPriorities:  MakeEncapsulationPrioritiesType(),
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		LinklocalServices:        MakeLinklocalServicesTypes(),
		EnableSecurityLogging:    false,
		ParentUUID:               "",
		ParentType:               "",
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
