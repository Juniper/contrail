package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeQosIdForwardingClassPairs makes QosIdForwardingClassPairs
// nolint
func MakeQosIdForwardingClassPairs() *QosIdForwardingClassPairs {
	return &QosIdForwardingClassPairs{
		//TODO(nati): Apply default

		QosIDForwardingClassPair: MakeQosIdForwardingClassPairSlice(),
	}
}

// MakeQosIdForwardingClassPairs makes QosIdForwardingClassPairs
// nolint
func InterfaceToQosIdForwardingClassPairs(i interface{}) *QosIdForwardingClassPairs {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &QosIdForwardingClassPairs{
		//TODO(nati): Apply default

		QosIDForwardingClassPair: InterfaceToQosIdForwardingClassPairSlice(m["qos_id_forwarding_class_pair"]),
	}
}

// MakeQosIdForwardingClassPairsSlice() makes a slice of QosIdForwardingClassPairs
// nolint
func MakeQosIdForwardingClassPairsSlice() []*QosIdForwardingClassPairs {
	return []*QosIdForwardingClassPairs{}
}

// InterfaceToQosIdForwardingClassPairsSlice() makes a slice of QosIdForwardingClassPairs
// nolint
func InterfaceToQosIdForwardingClassPairsSlice(i interface{}) []*QosIdForwardingClassPairs {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*QosIdForwardingClassPairs{}
	for _, item := range list {
		result = append(result, InterfaceToQosIdForwardingClassPairs(item))
	}
	return result
}
