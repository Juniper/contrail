package models

// MatchConditionType

import "encoding/json"

// MatchConditionType
type MatchConditionType struct {
	DSTPort    *PortType    `json:"dst_port"`
	Protocol   string       `json:"protocol"`
	SRCPort    *PortType    `json:"src_port"`
	SRCAddress *AddressType `json:"src_address"`
	Ethertype  EtherType    `json:"ethertype"`
	DSTAddress *AddressType `json:"dst_address"`
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
		DSTAddress: MakeAddressType(),
		DSTPort:    MakePortType(),
		Protocol:   "",
		SRCPort:    MakePortType(),
		SRCAddress: MakeAddressType(),
		Ethertype:  MakeEtherType(),
	}
}

// MakeMatchConditionTypeSlice() makes a slice of MatchConditionType
func MakeMatchConditionTypeSlice() []*MatchConditionType {
	return []*MatchConditionType{}
}
