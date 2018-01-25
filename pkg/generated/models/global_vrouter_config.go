package models

// GlobalVrouterConfig

// GlobalVrouterConfig
//proteus:generate
type GlobalVrouterConfig struct {
	UUID                       string                         `json:"uuid,omitempty"`
	ParentUUID                 string                         `json:"parent_uuid,omitempty"`
	ParentType                 string                         `json:"parent_type,omitempty"`
	FQName                     []string                       `json:"fq_name,omitempty"`
	IDPerms                    *IdPermsType                   `json:"id_perms,omitempty"`
	DisplayName                string                         `json:"display_name,omitempty"`
	Annotations                *KeyValuePairs                 `json:"annotations,omitempty"`
	Perms2                     *PermType2                     `json:"perms2,omitempty"`
	EcmpHashingIncludeFields   *EcmpHashingIncludeFields      `json:"ecmp_hashing_include_fields,omitempty"`
	FlowAgingTimeoutList       *FlowAgingTimeoutList          `json:"flow_aging_timeout_list,omitempty"`
	ForwardingMode             ForwardingModeType             `json:"forwarding_mode,omitempty"`
	FlowExportRate             int                            `json:"flow_export_rate,omitempty"`
	LinklocalServices          *LinklocalServicesTypes        `json:"linklocal_services,omitempty"`
	EncapsulationPriorities    *EncapsulationPrioritiesType   `json:"encapsulation_priorities,omitempty"`
	VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
	EnableSecurityLogging      bool                           `json:"enable_security_logging"`

	SecurityLoggingObjects []*SecurityLoggingObject `json:"security_logging_objects,omitempty"`
}

// MakeGlobalVrouterConfig makes GlobalVrouterConfig
func MakeGlobalVrouterConfig() *GlobalVrouterConfig {
	return &GlobalVrouterConfig{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		EcmpHashingIncludeFields:   MakeEcmpHashingIncludeFields(),
		FlowAgingTimeoutList:       MakeFlowAgingTimeoutList(),
		ForwardingMode:             MakeForwardingModeType(),
		FlowExportRate:             0,
		LinklocalServices:          MakeLinklocalServicesTypes(),
		EncapsulationPriorities:    MakeEncapsulationPrioritiesType(),
		VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
		EnableSecurityLogging:      false,
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}
