package models

// DiscoveryPubSubEndPointType

import "encoding/json"

// DiscoveryPubSubEndPointType
type DiscoveryPubSubEndPointType struct {
	EpVersion string      `json:"ep_version"`
	EpID      string      `json:"ep_id"`
	EpType    string      `json:"ep_type"`
	EpPrefix  *SubnetType `json:"ep_prefix"`
}

// String returns json representation of the object
func (model *DiscoveryPubSubEndPointType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDiscoveryPubSubEndPointType makes DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointType() *DiscoveryPubSubEndPointType {
	return &DiscoveryPubSubEndPointType{
		//TODO(nati): Apply default
		EpID:      "",
		EpType:    "",
		EpPrefix:  MakeSubnetType(),
		EpVersion: "",
	}
}

// MakeDiscoveryPubSubEndPointTypeSlice() makes a slice of DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointTypeSlice() []*DiscoveryPubSubEndPointType {
	return []*DiscoveryPubSubEndPointType{}
}
