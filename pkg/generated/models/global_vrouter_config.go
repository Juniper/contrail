package models

// GlobalVrouterConfig

import "encoding/json"

// GlobalVrouterConfig
type GlobalVrouterConfig struct {
	IDPerms                    *IdPermsType                   `json:"id_perms"`
	FlowExportRate             int                            `json:"flow_export_rate"`
	Annotations                *KeyValuePairs                 `json:"annotations"`
	ParentType                 string                         `json:"parent_type"`
	Perms2                     *PermType2                     `json:"perms2"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode"`
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode"`
	DisplayName                string                         `json:"display_name"`
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list"`
	ParentUUID                 string                         `json:"parent_uuid"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`
	UUID                       string                         `json:"uuid"`
	FQName                     []string                       `json:"fq_name"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields"`
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities"`

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
		Annotations:                MakeKeyValuePairs(),
		ParentType:                 "",
		IDPerms:                    MakeIdPermsType(),
		FlowExportRate:             0,
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		DisplayName:                "",
		Perms2:                     MakePermType2(),
		ForwardingMode:             MakeForwardingModeType(),
		ParentUUID:                 "",
		FlowAgingTimeoutList:       MakeFlowAgingTimeoutList(),
		LinklocalServices:          MakeLinklocalServicesTypes(),
		EncapsulationPriorities:    MakeEncapsulationPrioritiesType(),
		EnableSecurityLogging:      false,
		UUID:   "",
		FQName: []string{},
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
