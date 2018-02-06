package models

// DiscoveryPubSubEndPointType

import "encoding/json"

// DiscoveryPubSubEndPointType
type DiscoveryPubSubEndPointType struct {
	EpPrefix  *SubnetType `json:"ep_prefix,omitempty"`
	EpVersion string      `json:"ep_version,omitempty"`
	EpID      string      `json:"ep_id,omitempty"`
	EpType    string      `json:"ep_type,omitempty"`
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
		EpPrefix:  MakeSubnetType(),
		EpVersion: "",
		EpID:      "",
		EpType:    "",
	}
}

// MakeDiscoveryPubSubEndPointTypeSlice() makes a slice of DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointTypeSlice() []*DiscoveryPubSubEndPointType {
	return []*DiscoveryPubSubEndPointType{}
}
