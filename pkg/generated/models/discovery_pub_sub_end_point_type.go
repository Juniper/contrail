package models


// MakeDiscoveryPubSubEndPointType makes DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointType() *DiscoveryPubSubEndPointType{
    return &DiscoveryPubSubEndPointType{
    //TODO(nati): Apply default
    EpVersion: "",
        EpID: "",
        EpType: "",
        EpPrefix: MakeSubnetType(),
        
    }
}

// MakeDiscoveryPubSubEndPointTypeSlice() makes a slice of DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointTypeSlice() []*DiscoveryPubSubEndPointType {
    return []*DiscoveryPubSubEndPointType{}
}


