package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeQosQueue makes QosQueue
func MakeQosQueue() *QosQueue {
	return &QosQueue{
		//TODO(nati): Apply default
		UUID:               "",
		ParentUUID:         "",
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		QosQueueIdentifier: 0,
		MaxBandwidth:       0,
		MinBandwidth:       0,
	}
}

// MakeQosQueue makes QosQueue
func InterfaceToQosQueue(i interface{}) *QosQueue {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &QosQueue{
		//TODO(nati): Apply default
		UUID:               schema.InterfaceToString(m["uuid"]),
		ParentUUID:         schema.InterfaceToString(m["parent_uuid"]),
		ParentType:         schema.InterfaceToString(m["parent_type"]),
		FQName:             schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:            InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:        schema.InterfaceToString(m["display_name"]),
		Annotations:        InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:             InterfaceToPermType2(m["perms2"]),
		QosQueueIdentifier: schema.InterfaceToInt64(m["qos_queue_identifier"]),
		MaxBandwidth:       schema.InterfaceToInt64(m["max_bandwidth"]),
		MinBandwidth:       schema.InterfaceToInt64(m["min_bandwidth"]),
	}
}

// MakeQosQueueSlice() makes a slice of QosQueue
func MakeQosQueueSlice() []*QosQueue {
	return []*QosQueue{}
}

// InterfaceToQosQueueSlice() makes a slice of QosQueue
func InterfaceToQosQueueSlice(i interface{}) []*QosQueue {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*QosQueue{}
	for _, item := range list {
		result = append(result, InterfaceToQosQueue(item))
	}
	return result
}
