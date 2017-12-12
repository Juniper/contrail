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

// InterfaceToBridgeDomainMembershipType makes BridgeDomainMembershipType from interface
func InterfaceToBridgeDomainMembershipType(iData interface{}) *BridgeDomainMembershipType {
	data := iData.(map[string]interface{})
	return &BridgeDomainMembershipType{
		VlanTag: InterfaceToDot1QTagType(data["vlan_tag"]),

		//{"description":"VLAN tag of the incoming packet that maps the                      virtual-machine-interface to bridge domain","type":"integer","minimum":0,"maximum":4094}

	}
}

// InterfaceToBridgeDomainMembershipTypeSlice makes a slice of BridgeDomainMembershipType from interface
func InterfaceToBridgeDomainMembershipTypeSlice(data interface{}) []*BridgeDomainMembershipType {
	list := data.([]interface{})
	result := MakeBridgeDomainMembershipTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToBridgeDomainMembershipType(item))
	}
	return result
}

// MakeBridgeDomainMembershipTypeSlice() makes a slice of BridgeDomainMembershipType
func MakeBridgeDomainMembershipTypeSlice() []*BridgeDomainMembershipType {
	return []*BridgeDomainMembershipType{}
}
