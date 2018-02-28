package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeDiscoveryPubSubEndPointType makes DiscoveryPubSubEndPointType
// nolint
func MakeDiscoveryPubSubEndPointType() *DiscoveryPubSubEndPointType {
	return &DiscoveryPubSubEndPointType{
		//TODO(nati): Apply default
		EpVersion: "",
		EpID:      "",
		EpType:    "",
		EpPrefix:  MakeSubnetType(),
	}
}

// MakeDiscoveryPubSubEndPointType makes DiscoveryPubSubEndPointType
// nolint
func InterfaceToDiscoveryPubSubEndPointType(i interface{}) *DiscoveryPubSubEndPointType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DiscoveryPubSubEndPointType{
		//TODO(nati): Apply default
		EpVersion: common.InterfaceToString(m["ep_version"]),
		EpID:      common.InterfaceToString(m["ep_id"]),
		EpType:    common.InterfaceToString(m["ep_type"]),
		EpPrefix:  InterfaceToSubnetType(m["ep_prefix"]),
	}
}

// MakeDiscoveryPubSubEndPointTypeSlice() makes a slice of DiscoveryPubSubEndPointType
// nolint
func MakeDiscoveryPubSubEndPointTypeSlice() []*DiscoveryPubSubEndPointType {
	return []*DiscoveryPubSubEndPointType{}
}

// InterfaceToDiscoveryPubSubEndPointTypeSlice() makes a slice of DiscoveryPubSubEndPointType
// nolint
func InterfaceToDiscoveryPubSubEndPointTypeSlice(i interface{}) []*DiscoveryPubSubEndPointType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DiscoveryPubSubEndPointType{}
	for _, item := range list {
		result = append(result, InterfaceToDiscoveryPubSubEndPointType(item))
	}
	return result
}
