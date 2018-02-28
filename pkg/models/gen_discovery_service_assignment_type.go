package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeDiscoveryServiceAssignmentType makes DiscoveryServiceAssignmentType
// nolint
func MakeDiscoveryServiceAssignmentType() *DiscoveryServiceAssignmentType {
	return &DiscoveryServiceAssignmentType{
		//TODO(nati): Apply default

		Subscriber: MakeDiscoveryPubSubEndPointTypeSlice(),

		Publisher: MakeDiscoveryPubSubEndPointType(),
	}
}

// MakeDiscoveryServiceAssignmentType makes DiscoveryServiceAssignmentType
// nolint
func InterfaceToDiscoveryServiceAssignmentType(i interface{}) *DiscoveryServiceAssignmentType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DiscoveryServiceAssignmentType{
		//TODO(nati): Apply default

		Subscriber: InterfaceToDiscoveryPubSubEndPointTypeSlice(m["subscriber"]),

		Publisher: InterfaceToDiscoveryPubSubEndPointType(m["publisher"]),
	}
}

// MakeDiscoveryServiceAssignmentTypeSlice() makes a slice of DiscoveryServiceAssignmentType
// nolint
func MakeDiscoveryServiceAssignmentTypeSlice() []*DiscoveryServiceAssignmentType {
	return []*DiscoveryServiceAssignmentType{}
}

// InterfaceToDiscoveryServiceAssignmentTypeSlice() makes a slice of DiscoveryServiceAssignmentType
// nolint
func InterfaceToDiscoveryServiceAssignmentTypeSlice(i interface{}) []*DiscoveryServiceAssignmentType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DiscoveryServiceAssignmentType{}
	for _, item := range list {
		result = append(result, InterfaceToDiscoveryServiceAssignmentType(item))
	}
	return result
}
