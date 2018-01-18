package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	Annotations                *KeyValuePairs                 `json:"annotations,omitempty"`
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list,omitempty"`
	FlowExportRate             int                            `json:"flow_export_rate,omitempty"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities,omitempty"`
	ParentUUID                 string                         `json:"parent_uuid,omitempty"`
	ParentType                 string                         `json:"parent_type,omitempty"`
	IDPerms                    *IdPermsType                   `json:"id_perms,omitempty"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields,omitempty"`
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`
	UUID                       string                         `json:"uuid,omitempty"`
	DisplayName                string                         `json:"display_name,omitempty"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode,omitempty"`
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services,omitempty"`
	FQName                     []string                       `json:"fq_name,omitempty"`
	Perms2                     *PermType2                     `json:"perms2,omitempty"`

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
		FlowExportRate:             0,
		EncapsulationPriorities:    MakeEncapsulationPrioritiesType(),
		ParentUUID:                 "",
		ParentType:                 "",
		IDPerms:                    MakeIdPermsType(),
		Annotations:                MakeKeyValuePairs(),
		EcmpHashingIncludeFields:   MakeEcmpHashingIncludeFields(),
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		EnableSecurityLogging:      false,
		UUID:              "",
		DisplayName:       "",
		ForwardingMode:    MakeForwardingModeType(),
		LinklocalServices: MakeLinklocalServicesTypes(),
		FQName:            []string{},
		Perms2:            MakePermType2(),
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
