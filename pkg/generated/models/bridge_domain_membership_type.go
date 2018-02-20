package models

// BridgeDomainMembershipType

// BridgeDomainMembershipType
//proteus:generate
type BridgeDomainMembershipType struct {
	VlanTag Dot1QTagType `json:"vlan_tag,omitempty"`
}

// MakeBridgeDomainMembershipType makes BridgeDomainMembershipType
func MakeBridgeDomainMembershipType() *BridgeDomainMembershipType {
	return &BridgeDomainMembershipType{
		//TODO(nati): Apply default
		VlanTag: MakeDot1QTagType(),
	}
}

// MakeBridgeDomainMembershipTypeSlice() makes a slice of BridgeDomainMembershipType
func MakeBridgeDomainMembershipTypeSlice() []*BridgeDomainMembershipType {
	return []*BridgeDomainMembershipType{}
}
