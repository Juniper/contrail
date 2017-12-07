package models

// BridgeDomainMembershipType

import "encoding/json"

// BridgeDomainMembershipType
type BridgeDomainMembershipType struct {
	VlanTag Dot1QTagType `json:"vlan_tag"`
}

//  parents relation object

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

		//{"Title":"","Description":"VLAN tag of the incoming packet that maps the                      virtual-machine-interface to bridge domain","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":4094,"Ref":"types.json#/definitions/Dot1QTagType","CollectionType":"","Column":"vlan_tag","Item":null,"GoName":"VlanTag","GoType":"Dot1QTagType","GoPremitive":false}

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
