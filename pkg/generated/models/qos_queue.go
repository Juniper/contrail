package models

// QosQueue

import "encoding/json"

// QosQueue
type QosQueue struct {
	MaxBandwidth       int            `json:"max_bandwidth,omitempty"`
	FQName             []string       `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType   `json:"id_perms,omitempty"`
	QosQueueIdentifier int            `json:"qos_queue_identifier,omitempty"`
	MinBandwidth       int            `json:"min_bandwidth,omitempty"`
	Perms2             *PermType2     `json:"perms2,omitempty"`
	UUID               string         `json:"uuid,omitempty"`
	ParentUUID         string         `json:"parent_uuid,omitempty"`
	ParentType         string         `json:"parent_type,omitempty"`
	DisplayName        string         `json:"display_name,omitempty"`
	Annotations        *KeyValuePairs `json:"annotations,omitempty"`
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
		Perms2:             MakePermType2(),
		UUID:               "",
		ParentUUID:         "",
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		QosQueueIdentifier: 0,
		MinBandwidth:       0,
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		MaxBandwidth:       0,
	}
}

// MakeQosQueueSlice() makes a slice of QosQueue
func MakeQosQueueSlice() []*QosQueue {
	return []*QosQueue{}
}
