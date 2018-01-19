package models

// GlobalQosConfig

import "encoding/json"

// GlobalQosConfig
type GlobalQosConfig struct {
	ParentType         string                  `json:"parent_type,omitempty"`
	IDPerms            *IdPermsType            `json:"id_perms,omitempty"`
	DisplayName        string                  `json:"display_name,omitempty"`
	UUID               string                  `json:"uuid,omitempty"`
	ControlTrafficDSCP *ControlTrafficDscpType `json:"control_traffic_dscp,omitempty"`
	ParentUUID         string                  `json:"parent_uuid,omitempty"`
	Perms2             *PermType2              `json:"perms2,omitempty"`
	FQName             []string                `json:"fq_name,omitempty"`
	Annotations        *KeyValuePairs          `json:"annotations,omitempty"`

	ForwardingClasss []*ForwardingClass `json:"forwarding_classs,omitempty"`
	QosConfigs       []*QosConfig       `json:"qos_configs,omitempty"`
	QosQueues        []*QosQueue        `json:"qos_queues,omitempty"`
}

// String returns json representation of the object
func (model *GlobalQosConfig) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeGlobalQosConfig makes GlobalQosConfig
func MakeGlobalQosConfig() *GlobalQosConfig {
	return &GlobalQosConfig{
		//TODO(nati): Apply default
		FQName:             []string{},
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		UUID:               "",
		ControlTrafficDSCP: MakeControlTrafficDscpType(),
		ParentUUID:         "",
		ParentType:         "",
	}
}

// MakeGlobalQosConfigSlice() makes a slice of GlobalQosConfig
func MakeGlobalQosConfigSlice() []*GlobalQosConfig {
	return []*GlobalQosConfig{}
}
