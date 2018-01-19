package models

// DiscoveryServiceAssignmentType

import "encoding/json"

// DiscoveryServiceAssignmentType
type DiscoveryServiceAssignmentType struct {
	Subscriber []*DiscoveryPubSubEndPointType `json:"subscriber,omitempty"`
	Publisher  *DiscoveryPubSubEndPointType   `json:"publisher,omitempty"`
}

// String returns json representation of the object
func (model *DiscoveryServiceAssignmentType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDiscoveryServiceAssignmentType makes DiscoveryServiceAssignmentType
func MakeDiscoveryServiceAssignmentType() *DiscoveryServiceAssignmentType {
	return &DiscoveryServiceAssignmentType{
		//TODO(nati): Apply default
		Publisher: MakeDiscoveryPubSubEndPointType(),

		Subscriber: MakeDiscoveryPubSubEndPointTypeSlice(),
	}
}

// MakeDiscoveryServiceAssignmentTypeSlice() makes a slice of DiscoveryServiceAssignmentType
func MakeDiscoveryServiceAssignmentTypeSlice() []*DiscoveryServiceAssignmentType {
	return []*DiscoveryServiceAssignmentType{}
}
