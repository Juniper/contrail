package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeDiscoveryPubSubEndPointType makes DiscoveryPubSubEndPointType
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
func InterfaceToDiscoveryPubSubEndPointType(i interface{}) *DiscoveryPubSubEndPointType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DiscoveryPubSubEndPointType{
		//TODO(nati): Apply default
		EpVersion: schema.InterfaceToString(m["ep_version"]),
		EpID:      schema.InterfaceToString(m["ep_id"]),
		EpType:    schema.InterfaceToString(m["ep_type"]),
		EpPrefix:  InterfaceToSubnetType(m["ep_prefix"]),
	}
}

// MakeDiscoveryPubSubEndPointTypeSlice() makes a slice of DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointTypeSlice() []*DiscoveryPubSubEndPointType {
	return []*DiscoveryPubSubEndPointType{}
}

// InterfaceToDiscoveryPubSubEndPointTypeSlice() makes a slice of DiscoveryPubSubEndPointType
func InterfaceToDiscoveryPubSubEndPointTypeSlice(i interface{}) []*DiscoveryPubSubEndPointType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DiscoveryPubSubEndPointType{}
	for _, item := range list {
		result = append(result, InterfaceToDiscoveryPubSubEndPointType(item))
	}
	return result
}
