package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
	DisplayName                string                         `json:"display_name,omitempty"`
	Annotations                *KeyValuePairs                 `json:"annotations,omitempty"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields,omitempty"`
	ParentType                 string                         `json:"parent_type,omitempty"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`
	ParentUUID                 string                         `json:"parent_uuid,omitempty"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities,omitempty"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode,omitempty"`
	FlowExportRate             int                            `json:"flow_export_rate,omitempty"`
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services,omitempty"`
	Perms2                     *PermType2                     `json:"perms2,omitempty"`
	UUID                       string                         `json:"uuid,omitempty"`
	FQName                     []string                       `json:"fq_name,omitempty"`
	IDPerms                    *IdPermsType                   `json:"id_perms,omitempty"`
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list,omitempty"`

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
		Perms2:                     MakePermType2(),
		UUID:                       "",
		FQName:                     []string{},
		IDPerms:                    MakeIdPermsType(),
		FlowAgingTimeoutList:       MakeFlowAgingTimeoutList(),
		ForwardingMode:             MakeForwardingModeType(),
		FlowExportRate:             0,
		LinklocalServices:          MakeLinklocalServicesTypes(),
		EcmpHashingIncludeFields:   MakeEcmpHashingIncludeFields(),
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		DisplayName:                "",
		Annotations:                MakeKeyValuePairs(),
		ParentType:                 "",
		EncapsulationPriorities:    MakeEncapsulationPrioritiesType(),
		EnableSecurityLogging:      false,
		ParentUUID:                 "",
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
