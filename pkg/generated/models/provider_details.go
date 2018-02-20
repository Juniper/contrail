package models

// ProviderDetails

// ProviderDetails
//proteus:generate
type ProviderDetails struct {
	SegmentationID  VlanIdType `json:"segmentation_id,omitempty"`
	PhysicalNetwork string     `json:"physical_network,omitempty"`
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
