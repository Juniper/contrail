package models


// MakeProviderDetails makes ProviderDetails
func MakeProviderDetails() *ProviderDetails{
    return &ProviderDetails{
    //TODO(nati): Apply default
    SegmentationID: 0,
        PhysicalNetwork: "",
        
    }
}

// MakeProviderDetailsSlice() makes a slice of ProviderDetails
func MakeProviderDetailsSlice() []*ProviderDetails {
    return []*ProviderDetails{}
}


