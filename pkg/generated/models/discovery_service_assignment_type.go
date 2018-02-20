package models

// DiscoveryServiceAssignmentType

// DiscoveryServiceAssignmentType
//proteus:generate
type DiscoveryServiceAssignmentType struct {
	Subscriber []*DiscoveryPubSubEndPointType `json:"subscriber,omitempty"`
	Publisher  *DiscoveryPubSubEndPointType   `json:"publisher,omitempty"`
}

// MakeDiscoveryServiceAssignmentType makes DiscoveryServiceAssignmentType
func MakeDiscoveryServiceAssignmentType() *DiscoveryServiceAssignmentType {
	return &DiscoveryServiceAssignmentType{
		//TODO(nati): Apply default

		Subscriber: MakeDiscoveryPubSubEndPointTypeSlice(),

		Publisher: MakeDiscoveryPubSubEndPointType(),
	}
}

// MakeDiscoveryServiceAssignmentTypeSlice() makes a slice of DiscoveryServiceAssignmentType
func MakeDiscoveryServiceAssignmentTypeSlice() []*DiscoveryServiceAssignmentType {
	return []*DiscoveryServiceAssignmentType{}
}
