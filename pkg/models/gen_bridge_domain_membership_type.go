package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeBridgeDomainMembershipType makes BridgeDomainMembershipType
// nolint
func MakeBridgeDomainMembershipType() *BridgeDomainMembershipType {
	return &BridgeDomainMembershipType{
		//TODO(nati): Apply default
		VlanTag: 0,
	}
}

// MakeBridgeDomainMembershipType makes BridgeDomainMembershipType
// nolint
func InterfaceToBridgeDomainMembershipType(i interface{}) *BridgeDomainMembershipType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BridgeDomainMembershipType{
		//TODO(nati): Apply default
		VlanTag: common.InterfaceToInt64(m["vlan_tag"]),
	}
}

// MakeBridgeDomainMembershipTypeSlice() makes a slice of BridgeDomainMembershipType
// nolint
func MakeBridgeDomainMembershipTypeSlice() []*BridgeDomainMembershipType {
	return []*BridgeDomainMembershipType{}
}

// InterfaceToBridgeDomainMembershipTypeSlice() makes a slice of BridgeDomainMembershipType
// nolint
func InterfaceToBridgeDomainMembershipTypeSlice(i interface{}) []*BridgeDomainMembershipType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BridgeDomainMembershipType{}
	for _, item := range list {
		result = append(result, InterfaceToBridgeDomainMembershipType(item))
	}
	return result
}
