package models

// QosQueue

import "encoding/json"

// QosQueue
type QosQueue struct {
	FQName             []string       `json:"fq_name,omitempty"`
	QosQueueIdentifier int            `json:"qos_queue_identifier,omitempty"`
	IDPerms            *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName        string         `json:"display_name,omitempty"`
	Annotations        *KeyValuePairs `json:"annotations,omitempty"`
	Perms2             *PermType2     `json:"perms2,omitempty"`
	ParentUUID         string         `json:"parent_uuid,omitempty"`
	ParentType         string         `json:"parent_type,omitempty"`
	MaxBandwidth       int            `json:"max_bandwidth,omitempty"`
	MinBandwidth       int            `json:"min_bandwidth,omitempty"`
	UUID               string         `json:"uuid,omitempty"`
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
		ParentUUID:         "",
		ParentType:         "",
		FQName:             []string{},
		QosQueueIdentifier: 0,
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		MaxBandwidth:       0,
		MinBandwidth:       0,
		UUID:               "",
	}
}

// MakeQosQueueSlice() makes a slice of QosQueue
func MakeQosQueueSlice() []*QosQueue {
	return []*QosQueue{}
}
