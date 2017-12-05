package models

// BridgeDomainMembershipType

import "encoding/json"

type BridgeDomainMembershipType struct {
	VlanTag Dot1QTagType `json:"vlan_tag"`
}

func (model *BridgeDomainMembershipType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeBridgeDomainMembershipType() *BridgeDomainMembershipType {
	return &BridgeDomainMembershipType{
		//TODO(nati): Apply default
		VlanTag: MakeDot1QTagType(),
	}
}

func InterfaceToBridgeDomainMembershipType(iData interface{}) *BridgeDomainMembershipType {
	data := iData.(map[string]interface{})
	return &BridgeDomainMembershipType{
		VlanTag: InterfaceToDot1QTagType(data["vlan_tag"]),

		//{"Title":"","Description":"VLAN tag of the incoming packet that maps the                      virtual-machine-interface to bridge domain","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":4094,"Ref":"types.json#/definitions/Dot1QTagType","CollectionType":"","Column":"vlan_tag","Item":null,"GoName":"VlanTag","GoType":"Dot1QTagType"}

	}
}

func InterfaceToBridgeDomainMembershipTypeSlice(data interface{}) []*BridgeDomainMembershipType {
	list := data.([]interface{})
	result := MakeBridgeDomainMembershipTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToBridgeDomainMembershipType(item))
	}
	return result
}

func MakeBridgeDomainMembershipTypeSlice() []*BridgeDomainMembershipType {
	return []*BridgeDomainMembershipType{}
}
