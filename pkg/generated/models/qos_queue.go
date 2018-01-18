package models

// QosQueue

import "encoding/json"

// QosQueue
type QosQueue struct {
	QosQueueIdentifier int            `json:"qos_queue_identifier,omitempty"`
	MinBandwidth       int            `json:"min_bandwidth,omitempty"`
	Perms2             *PermType2     `json:"perms2,omitempty"`
	ParentUUID         string         `json:"parent_uuid,omitempty"`
	Annotations        *KeyValuePairs `json:"annotations,omitempty"`
	MaxBandwidth       int            `json:"max_bandwidth,omitempty"`
	UUID               string         `json:"uuid,omitempty"`
	ParentType         string         `json:"parent_type,omitempty"`
	FQName             []string       `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName        string         `json:"display_name,omitempty"`
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
		DisplayName:        "",
		MaxBandwidth:       0,
		UUID:               "",
		ParentType:         "",
		FQName:             []string{},
		Annotations:        MakeKeyValuePairs(),
		QosQueueIdentifier: 0,
		MinBandwidth:       0,
		Perms2:             MakePermType2(),
		ParentUUID:         "",
	}
}

// MakeQosQueueSlice() makes a slice of QosQueue
func MakeQosQueueSlice() []*QosQueue {
	return []*QosQueue{}
}
