package models

// MatchConditionType

// MatchConditionType
//proteus:generate
type MatchConditionType struct {
	SRCPort    *PortType    `json:"src_port,omitempty"`
	SRCAddress *AddressType `json:"src_address,omitempty"`
	Ethertype  EtherType    `json:"ethertype,omitempty"`
	DSTAddress *AddressType `json:"dst_address,omitempty"`
	DSTPort    *PortType    `json:"dst_port,omitempty"`
	Protocol   string       `json:"protocol,omitempty"`
}

// MakeMatchConditionType makes MatchConditionType
func MakeMatchConditionType() *MatchConditionType {
	return &MatchConditionType{
		//TODO(nati): Apply default
		SRCPort:    MakePortType(),
		SRCAddress: MakeAddressType(),
		Ethertype:  MakeEtherType(),
		DSTAddress: MakeAddressType(),
		DSTPort:    MakePortType(),
		Protocol:   "",
	}
}

// MakeMatchConditionTypeSlice() makes a slice of MatchConditionType
func MakeMatchConditionTypeSlice() []*MatchConditionType {
	return []*MatchConditionType{}
}
