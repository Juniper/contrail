package models

// GlobalQosConfig

import "encoding/json"

// GlobalQosConfig
type GlobalQosConfig struct {
	Annotations        *KeyValuePairs          `json:"annotations"`
	Perms2             *PermType2              `json:"perms2"`
	ControlTrafficDSCP *ControlTrafficDscpType `json:"control_traffic_dscp"`
	FQName             []string                `json:"fq_name"`
	IDPerms            *IdPermsType            `json:"id_perms"`
	DisplayName        string                  `json:"display_name"`
	UUID               string                  `json:"uuid"`
	ParentUUID         string                  `json:"parent_uuid"`
	ParentType         string                  `json:"parent_type"`

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
		IDPerms:            MakeIdPermsType(),
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		ControlTrafficDSCP: MakeControlTrafficDscpType(),
		FQName:             []string{},
		ParentType:         "",
		DisplayName:        "",
		UUID:               "",
		ParentUUID:         "",
	}
}

// InterfaceToGlobalQosConfig makes GlobalQosConfig from interface
func InterfaceToGlobalQosConfig(iData interface{}) *GlobalQosConfig {
	data := iData.(map[string]interface{})
	return &GlobalQosConfig{
		ControlTrafficDSCP: InterfaceToControlTrafficDscpType(data["control_traffic_dscp"]),

		//{"description":"DSCP value of IP header for control traffic","type":"object","properties":{"analytics":{"type":"integer","minimum":0,"maximum":63},"control":{"type":"integer","minimum":0,"maximum":63},"dns":{"type":"integer","minimum":0,"maximum":63}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

	}
}

// InterfaceToGlobalQosConfigSlice makes a slice of GlobalQosConfig from interface
func InterfaceToGlobalQosConfigSlice(data interface{}) []*GlobalQosConfig {
	list := data.([]interface{})
	result := MakeGlobalQosConfigSlice()
	for _, item := range list {
		result = append(result, InterfaceToGlobalQosConfig(item))
	}
	return result
}

// MakeGlobalQosConfigSlice() makes a slice of GlobalQosConfig
func MakeGlobalQosConfigSlice() []*GlobalQosConfig {
	return []*GlobalQosConfig{}
}
