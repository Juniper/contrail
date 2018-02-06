package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	Created      string          `json:"created,omitempty"`
	RuleUUID     string          `json:"rule_uuid,omitempty"`
	Application  []string        `json:"application,omitempty"`
	Ethertype    EtherType       `json:"ethertype,omitempty"`
	SRCAddresses []*AddressType  `json:"src_addresses,omitempty"`
	Direction    DirectionType   `json:"direction,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
	DSTAddresses []*AddressType  `json:"dst_addresses,omitempty"`
	SRCPorts     []*PortType     `json:"src_ports,omitempty"`
	RuleSequence *SequenceType   `json:"rule_sequence,omitempty"`
	ActionList   *ActionListType `json:"action_list,omitempty"`
	DSTPorts     []*PortType     `json:"dst_ports,omitempty"`
	LastModified string          `json:"last_modified,omitempty"`
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
		Created:     "",
		RuleUUID:    "",
		Application: []string{},
		Ethertype:   MakeEtherType(),

		SRCAddresses: MakeAddressTypeSlice(),

		Direction: MakeDirectionType(),
		Protocol:  "",

		DSTAddresses: MakeAddressTypeSlice(),

		SRCPorts: MakePortTypeSlice(),

		RuleSequence: MakeSequenceType(),
		ActionList:   MakeActionListType(),

		DSTPorts: MakePortTypeSlice(),

		LastModified: "",
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
