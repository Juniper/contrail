package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePolicyRuleType makes PolicyRuleType
// nolint
func MakePolicyRuleType() *PolicyRuleType {
	return &PolicyRuleType{
		//TODO(nati): Apply default
		Direction: "",
		Protocol:  "",

		DSTAddresses: MakeAddressTypeSlice(),

		ActionList: MakeActionListType(),
		Created:    "",
		RuleUUID:   "",

		DSTPorts: MakePortTypeSlice(),

		Application:  []string{},
		LastModified: "",
		Ethertype:    "",

		SRCAddresses: MakeAddressTypeSlice(),

		RuleSequence: MakeSequenceType(),

		SRCPorts: MakePortTypeSlice(),
	}
}

// MakePolicyRuleType makes PolicyRuleType
// nolint
func InterfaceToPolicyRuleType(i interface{}) *PolicyRuleType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PolicyRuleType{
		//TODO(nati): Apply default
		Direction: common.InterfaceToString(m["direction"]),
		Protocol:  common.InterfaceToString(m["protocol"]),

		DSTAddresses: InterfaceToAddressTypeSlice(m["dst_addresses"]),

		ActionList: InterfaceToActionListType(m["action_list"]),
		Created:    common.InterfaceToString(m["created"]),
		RuleUUID:   common.InterfaceToString(m["rule_uuid"]),

		DSTPorts: InterfaceToPortTypeSlice(m["dst_ports"]),

		Application:  common.InterfaceToStringList(m["application"]),
		LastModified: common.InterfaceToString(m["last_modified"]),
		Ethertype:    common.InterfaceToString(m["ethertype"]),

		SRCAddresses: InterfaceToAddressTypeSlice(m["src_addresses"]),

		RuleSequence: InterfaceToSequenceType(m["rule_sequence"]),

		SRCPorts: InterfaceToPortTypeSlice(m["src_ports"]),
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
// nolint
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}

// InterfaceToPolicyRuleTypeSlice() makes a slice of PolicyRuleType
// nolint
func InterfaceToPolicyRuleTypeSlice(i interface{}) []*PolicyRuleType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PolicyRuleType{}
	for _, item := range list {
		result = append(result, InterfaceToPolicyRuleType(item))
	}
	return result
}
