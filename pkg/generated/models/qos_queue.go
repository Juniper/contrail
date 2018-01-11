package models

// QosQueue

import "encoding/json"

// QosQueue
type QosQueue struct {
	FQName             []string       `json:"fq_name"`
	IDPerms            *IdPermsType   `json:"id_perms"`
	DisplayName        string         `json:"display_name"`
	UUID               string         `json:"uuid"`
	QosQueueIdentifier int            `json:"qos_queue_identifier"`
	MinBandwidth       int            `json:"min_bandwidth"`
	ParentUUID         string         `json:"parent_uuid"`
	Perms2             *PermType2     `json:"perms2"`
	MaxBandwidth       int            `json:"max_bandwidth"`
	ParentType         string         `json:"parent_type"`
	Annotations        *KeyValuePairs `json:"annotations"`
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
		ParentUUID:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		UUID:               "",
		QosQueueIdentifier: 0,
		ParentType:         "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		MaxBandwidth:       0,
	}
}

// MakeQosQueueSlice() makes a slice of QosQueue
func MakeQosQueueSlice() []*QosQueue {
	return []*QosQueue{}
}
