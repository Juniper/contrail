package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeMatchConditionType makes MatchConditionType
// nolint
func MakeMatchConditionType() *MatchConditionType {
	return &MatchConditionType{
		//TODO(nati): Apply default
		SRCPort:    MakePortType(),
		SRCAddress: MakeAddressType(),
		Ethertype:  "",
		DSTAddress: MakeAddressType(),
		DSTPort:    MakePortType(),
		Protocol:   "",
	}
}

// MakeMatchConditionType makes MatchConditionType
// nolint
func InterfaceToMatchConditionType(i interface{}) *MatchConditionType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &MatchConditionType{
		//TODO(nati): Apply default
		SRCPort:    InterfaceToPortType(m["src_port"]),
		SRCAddress: InterfaceToAddressType(m["src_address"]),
		Ethertype:  common.InterfaceToString(m["ethertype"]),
		DSTAddress: InterfaceToAddressType(m["dst_address"]),
		DSTPort:    InterfaceToPortType(m["dst_port"]),
		Protocol:   common.InterfaceToString(m["protocol"]),
	}
}

// MakeMatchConditionTypeSlice() makes a slice of MatchConditionType
// nolint
func MakeMatchConditionTypeSlice() []*MatchConditionType {
	return []*MatchConditionType{}
}

// InterfaceToMatchConditionTypeSlice() makes a slice of MatchConditionType
// nolint
func InterfaceToMatchConditionTypeSlice(i interface{}) []*MatchConditionType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*MatchConditionType{}
	for _, item := range list {
		result = append(result, InterfaceToMatchConditionType(item))
	}
	return result
}
