package models

// QosQueue

import "encoding/json"

// QosQueue
type QosQueue struct {
	Perms2             *PermType2     `json:"perms2"`
	IDPerms            *IdPermsType   `json:"id_perms"`
	Annotations        *KeyValuePairs `json:"annotations"`
	MinBandwidth       int            `json:"min_bandwidth"`
	DisplayName        string         `json:"display_name"`
	UUID               string         `json:"uuid"`
	ParentUUID         string         `json:"parent_uuid"`
	ParentType         string         `json:"parent_type"`
	FQName             []string       `json:"fq_name"`
	QosQueueIdentifier int            `json:"qos_queue_identifier"`
	MaxBandwidth       int            `json:"max_bandwidth"`
}

// String returns json representation of the object
func (model *QosQueue) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeQosQueue makes QosQueue
func MakeQosQueue() *QosQueue {
	return &QosQueue{
		//TODO(nati): Apply default
		IDPerms:            MakeIdPermsType(),
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		UUID:               "",
		ParentUUID:         "",
		ParentType:         "",
		FQName:             []string{},
		QosQueueIdentifier: 0,
		MaxBandwidth:       0,
		MinBandwidth:       0,
		DisplayName:        "",
	}
}

// InterfaceToQosQueue makes QosQueue from interface
func InterfaceToQosQueue(iData interface{}) *QosQueue {
	data := iData.(map[string]interface{})
	return &QosQueue{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		QosQueueIdentifier: data["qos_queue_identifier"].(int),

		//{"description":"Unique id for this queue.","type":"integer"}
		MaxBandwidth: data["max_bandwidth"].(int),

		//{"description":"Maximum bandwidth for this queue.","type":"integer"}
		MinBandwidth: data["min_bandwidth"].(int),

		//{"description":"Minimum bandwidth for this queue.","type":"integer"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}

	}
}

// InterfaceToQosQueueSlice makes a slice of QosQueue from interface
func InterfaceToQosQueueSlice(data interface{}) []*QosQueue {
	list := data.([]interface{})
	result := MakeQosQueueSlice()
	for _, item := range list {
		result = append(result, InterfaceToQosQueue(item))
	}
	return result
}

// MakeQosQueueSlice() makes a slice of QosQueue
func MakeQosQueueSlice() []*QosQueue {
	return []*QosQueue{}
}
