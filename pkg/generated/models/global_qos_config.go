package models

// GlobalQosConfig

import "encoding/json"

// GlobalQosConfig
type GlobalQosConfig struct {
	DisplayName        string                  `json:"display_name,omitempty"`
	Annotations        *KeyValuePairs          `json:"annotations,omitempty"`
	Perms2             *PermType2              `json:"perms2,omitempty"`
	ParentType         string                  `json:"parent_type,omitempty"`
	ControlTrafficDSCP *ControlTrafficDscpType `json:"control_traffic_dscp,omitempty"`
	IDPerms            *IdPermsType            `json:"id_perms,omitempty"`
	FQName             []string                `json:"fq_name,omitempty"`
	UUID               string                  `json:"uuid,omitempty"`
	ParentUUID         string                  `json:"parent_uuid,omitempty"`

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
		UUID:               "",
		ParentUUID:         "",
		FQName:             []string{},
		ControlTrafficDSCP: MakeControlTrafficDscpType(),
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		ParentType:         "",
	}
}

// MakeGlobalQosConfigSlice() makes a slice of GlobalQosConfig
func MakeGlobalQosConfigSlice() []*GlobalQosConfig {
	return []*GlobalQosConfig{}
}
