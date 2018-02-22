package models


// MakeBridgeDomainMembershipType makes BridgeDomainMembershipType
func MakeBridgeDomainMembershipType() *BridgeDomainMembershipType{
    return &BridgeDomainMembershipType{
    //TODO(nati): Apply default
    VlanTag: 0,
        
    }
}

// MakeBridgeDomainMembershipTypeSlice() makes a slice of BridgeDomainMembershipType
func MakeBridgeDomainMembershipTypeSlice() []*BridgeDomainMembershipType {
    return []*BridgeDomainMembershipType{}
}


