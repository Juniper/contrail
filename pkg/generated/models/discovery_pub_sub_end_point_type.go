package models

// DiscoveryPubSubEndPointType

import "encoding/json"

// DiscoveryPubSubEndPointType
type DiscoveryPubSubEndPointType struct {
	EpVersion string      `json:"ep_version,omitempty"`
	EpID      string      `json:"ep_id,omitempty"`
	EpType    string      `json:"ep_type,omitempty"`
	EpPrefix  *SubnetType `json:"ep_prefix,omitempty"`
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
		EpType:    "",
		EpPrefix:  MakeSubnetType(),
		EpVersion: "",
		EpID:      "",
	}
}

// MakeDiscoveryPubSubEndPointTypeSlice() makes a slice of DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointTypeSlice() []*DiscoveryPubSubEndPointType {
	return []*DiscoveryPubSubEndPointType{}
}
