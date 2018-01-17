package models

// MatchConditionType

import "encoding/json"

// MatchConditionType
type MatchConditionType struct {
	Protocol   string       `json:"protocol,omitempty"`
	SRCPort    *PortType    `json:"src_port,omitempty"`
	SRCAddress *AddressType `json:"src_address,omitempty"`
	Ethertype  EtherType    `json:"ethertype,omitempty"`
	DSTAddress *AddressType `json:"dst_address,omitempty"`
	DSTPort    *PortType    `json:"dst_port,omitempty"`
}

// String returns json representation of the object
func (model *MatchConditionType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeMatchConditionType makes MatchConditionType
func MakeMatchConditionType() *MatchConditionType {
	return &MatchConditionType{
		//TODO(nati): Apply default
		DSTPort:    MakePortType(),
		Protocol:   "",
		SRCPort:    MakePortType(),
		SRCAddress: MakeAddressType(),
		Ethertype:  MakeEtherType(),
		DSTAddress: MakeAddressType(),
	}
}

// MakeMatchConditionTypeSlice() makes a slice of MatchConditionType
func MakeMatchConditionTypeSlice() []*MatchConditionType {
	return []*MatchConditionType{}
}
