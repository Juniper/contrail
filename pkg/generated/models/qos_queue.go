package models

// QosQueue

import "encoding/json"

// QosQueue
type QosQueue struct {
	QosQueueIdentifier int            `json:"qos_queue_identifier,omitempty"`
	MaxBandwidth       int            `json:"max_bandwidth,omitempty"`
	Annotations        *KeyValuePairs `json:"annotations,omitempty"`
	FQName             []string       `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName        string         `json:"display_name,omitempty"`
	MinBandwidth       int            `json:"min_bandwidth,omitempty"`
	Perms2             *PermType2     `json:"perms2,omitempty"`
	UUID               string         `json:"uuid,omitempty"`
	ParentUUID         string         `json:"parent_uuid,omitempty"`
	ParentType         string         `json:"parent_type,omitempty"`
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
		ParentType:         "",
		DisplayName:        "",
		MinBandwidth:       0,
		Perms2:             MakePermType2(),
		UUID:               "",
		ParentUUID:         "",
		IDPerms:            MakeIdPermsType(),
		QosQueueIdentifier: 0,
		MaxBandwidth:       0,
		Annotations:        MakeKeyValuePairs(),
		FQName:             []string{},
	}
}

// MakeQosQueueSlice() makes a slice of QosQueue
func MakeQosQueueSlice() []*QosQueue {
	return []*QosQueue{}
}
