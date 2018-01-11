package models

// BridgeDomainMembershipType

import "encoding/json"

// BridgeDomainMembershipType
type BridgeDomainMembershipType struct {
	VlanTag Dot1QTagType `json:"vlan_tag"`
}

// String returns json representation of the object
func (model *BridgeDomainMembershipType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
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
