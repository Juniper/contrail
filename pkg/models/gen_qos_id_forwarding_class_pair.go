package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeQosIdForwardingClassPair makes QosIdForwardingClassPair
// nolint
func MakeQosIdForwardingClassPair() *QosIdForwardingClassPair {
	return &QosIdForwardingClassPair{
		//TODO(nati): Apply default
		Key:               0,
		ForwardingClassID: 0,
	}
}

// MakeQosIdForwardingClassPair makes QosIdForwardingClassPair
// nolint
func InterfaceToQosIdForwardingClassPair(i interface{}) *QosIdForwardingClassPair {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &QosIdForwardingClassPair{
		//TODO(nati): Apply default
		Key:               common.InterfaceToInt64(m["key"]),
		ForwardingClassID: common.InterfaceToInt64(m["forwarding_class_id"]),
	}
}

// MakeQosIdForwardingClassPairSlice() makes a slice of QosIdForwardingClassPair
// nolint
func MakeQosIdForwardingClassPairSlice() []*QosIdForwardingClassPair {
	return []*QosIdForwardingClassPair{}
}

// InterfaceToQosIdForwardingClassPairSlice() makes a slice of QosIdForwardingClassPair
// nolint
func InterfaceToQosIdForwardingClassPairSlice(i interface{}) []*QosIdForwardingClassPair {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*QosIdForwardingClassPair{}
	for _, item := range list {
		result = append(result, InterfaceToQosIdForwardingClassPair(item))
	}
	return result
}
