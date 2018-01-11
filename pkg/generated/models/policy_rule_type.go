package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	DSTPorts     []*PortType     `json:"dst_ports"`
	Application  []string        `json:"application"`
	LastModified string          `json:"last_modified"`
	Ethertype    EtherType       `json:"ethertype"`
	SRCPorts     []*PortType     `json:"src_ports"`
	DSTAddresses []*AddressType  `json:"dst_addresses"`
	Created      string          `json:"created"`
	RuleUUID     string          `json:"rule_uuid"`
	SRCAddresses []*AddressType  `json:"src_addresses"`
	RuleSequence *SequenceType   `json:"rule_sequence"`
	Direction    DirectionType   `json:"direction"`
	Protocol     string          `json:"protocol"`
	ActionList   *ActionListType `json:"action_list"`
}

// String returns json representation of the object
func (model *PolicyRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePolicyRuleType makes PolicyRuleType
func MakePolicyRuleType() *PolicyRuleType {
	return &PolicyRuleType{
		//TODO(nati): Apply default
		Direction:  MakeDirectionType(),
		Protocol:   "",
		ActionList: MakeActionListType(),

		SRCAddresses: MakeAddressTypeSlice(),

		RuleSequence: MakeSequenceType(),
		LastModified: "",
		Ethertype:    MakeEtherType(),

		SRCPorts: MakePortTypeSlice(),

		DSTAddresses: MakeAddressTypeSlice(),

		Created:  "",
		RuleUUID: "",

		DSTPorts: MakePortTypeSlice(),

		Application: []string{},
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
