package models

// ProviderDetails

import "encoding/json"

// ProviderDetails
type ProviderDetails struct {
	PhysicalNetwork string     `json:"physical_network,omitempty"`
	SegmentationID  VlanIdType `json:"segmentation_id,omitempty"`
}

// String returns json representation of the object
func (model *ProviderDetails) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeProviderDetails makes ProviderDetails
func MakeProviderDetails() *ProviderDetails {
	return &ProviderDetails{
		//TODO(nati): Apply default
		SegmentationID:  MakeVlanIdType(),
		PhysicalNetwork: "",
	}
}

// MakeProviderDetailsSlice() makes a slice of ProviderDetails
func MakeProviderDetailsSlice() []*ProviderDetails {
	return []*ProviderDetails{}
}
