package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeMatchConditionType makes MatchConditionType
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
		Ethertype:  schema.InterfaceToString(m["ethertype"]),
		DSTAddress: InterfaceToAddressType(m["dst_address"]),
		DSTPort:    InterfaceToPortType(m["dst_port"]),
		Protocol:   schema.InterfaceToString(m["protocol"]),
	}
}

// MakeMatchConditionTypeSlice() makes a slice of MatchConditionType
func MakeMatchConditionTypeSlice() []*MatchConditionType {
	return []*MatchConditionType{}
}

// InterfaceToMatchConditionTypeSlice() makes a slice of MatchConditionType
func InterfaceToMatchConditionTypeSlice(i interface{}) []*MatchConditionType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*MatchConditionType{}
	for _, item := range list {
		result = append(result, InterfaceToMatchConditionType(item))
	}
	return result
}
