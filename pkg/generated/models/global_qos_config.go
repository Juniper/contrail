package models

// GlobalQosConfig

import "encoding/json"

// GlobalQosConfig
type GlobalQosConfig struct {
	IDPerms            *IdPermsType            `json:"id_perms"`
	DisplayName        string                  `json:"display_name"`
	Annotations        *KeyValuePairs          `json:"annotations"`
	Perms2             *PermType2              `json:"perms2"`
	ParentType         string                  `json:"parent_type"`
	FQName             []string                `json:"fq_name"`
	ControlTrafficDSCP *ControlTrafficDscpType `json:"control_traffic_dscp"`
	UUID               string                  `json:"uuid"`
	ParentUUID         string                  `json:"parent_uuid"`

	ForwardingClasss []*ForwardingClass `json:"forwarding_classs"`
	QosConfigs       []*QosConfig       `json:"qos_configs"`
	QosQueues        []*QosQueue        `json:"qos_queues"`
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
		ControlTrafficDSCP: MakeControlTrafficDscpType(),
		UUID:               "",
		ParentUUID:         "",
		Perms2:             MakePermType2(),
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
	}
}

// MakeGlobalQosConfigSlice() makes a slice of GlobalQosConfig
func MakeGlobalQosConfigSlice() []*GlobalQosConfig {
	return []*GlobalQosConfig{}
}
