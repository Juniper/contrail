package models

// QosQueue

import "encoding/json"

// QosQueue
type QosQueue struct {
	FQName             []string       `json:"fq_name,omitempty"`
	DisplayName        string         `json:"display_name,omitempty"`
	Perms2             *PermType2     `json:"perms2,omitempty"`
	UUID               string         `json:"uuid,omitempty"`
	MinBandwidth       int            `json:"min_bandwidth,omitempty"`
	MaxBandwidth       int            `json:"max_bandwidth,omitempty"`
	ParentType         string         `json:"parent_type,omitempty"`
	IDPerms            *IdPermsType   `json:"id_perms,omitempty"`
	Annotations        *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID         string         `json:"parent_uuid,omitempty"`
	QosQueueIdentifier int            `json:"qos_queue_identifier,omitempty"`
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
		MinBandwidth:       0,
		FQName:             []string{},
		DisplayName:        "",
		Perms2:             MakePermType2(),
		UUID:               "",
		ParentUUID:         "",
		QosQueueIdentifier: 0,
		MaxBandwidth:       0,
		ParentType:         "",
		IDPerms:            MakeIdPermsType(),
		Annotations:        MakeKeyValuePairs(),
	}
}

// MakeQosQueueSlice() makes a slice of QosQueue
func MakeQosQueueSlice() []*QosQueue {
	return []*QosQueue{}
}
