package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	Perms2                     *PermType2                     `json:"perms2,omitempty"`
	IDPerms                    *IdPermsType                   `json:"id_perms,omitempty"`
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services,omitempty"`
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
	ParentType                 string                         `json:"parent_type,omitempty"`
	FQName                     []string                       `json:"fq_name,omitempty"`
	DisplayName                string                         `json:"display_name,omitempty"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields,omitempty"`
	FlowExportRate             int                            `json:"flow_export_rate,omitempty"`
	ParentUUID                 string                         `json:"parent_uuid,omitempty"`
	Annotations                *KeyValuePairs                 `json:"annotations,omitempty"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities,omitempty"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`
	UUID                       string                         `json:"uuid,omitempty"`
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list,omitempty"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode,omitempty"`

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
		LinklocalServices:          MakeLinklocalServicesTypes(),
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		ParentType:                 "",
		FQName:                     []string{},
		DisplayName:                "",
		EcmpHashingIncludeFields:   MakeEcmpHashingIncludeFields(),
		FlowExportRate:             0,
		ParentUUID:                 "",
		Annotations:                MakeKeyValuePairs(),
		EncapsulationPriorities:    MakeEncapsulationPrioritiesType(),
		EnableSecurityLogging:      false,
		UUID:                 "",
		FlowAgingTimeoutList: MakeFlowAgingTimeoutList(),
		ForwardingMode:       MakeForwardingModeType(),
		Perms2:               MakePermType2(),
		IDPerms:              MakeIdPermsType(),
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
