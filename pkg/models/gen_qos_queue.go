package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeQosQueue makes QosQueue
// nolint
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
// nolint
func InterfaceToQosQueue(i interface{}) *QosQueue {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &QosQueue{
		//TODO(nati): Apply default
		UUID:               common.InterfaceToString(m["uuid"]),
		ParentUUID:         common.InterfaceToString(m["parent_uuid"]),
		ParentType:         common.InterfaceToString(m["parent_type"]),
		FQName:             common.InterfaceToStringList(m["fq_name"]),
		IDPerms:            InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:        common.InterfaceToString(m["display_name"]),
		Annotations:        InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:             InterfaceToPermType2(m["perms2"]),
		QosQueueIdentifier: common.InterfaceToInt64(m["qos_queue_identifier"]),
		MaxBandwidth:       common.InterfaceToInt64(m["max_bandwidth"]),
		MinBandwidth:       common.InterfaceToInt64(m["min_bandwidth"]),
	}
}

// MakeQosQueueSlice() makes a slice of QosQueue
// nolint
func MakeQosQueueSlice() []*QosQueue {
	return []*QosQueue{}
}

// InterfaceToQosQueueSlice() makes a slice of QosQueue
// nolint
func InterfaceToQosQueueSlice(i interface{}) []*QosQueue {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*QosQueue{}
	for _, item := range list {
		result = append(result, InterfaceToQosQueue(item))
	}
	return result
}
