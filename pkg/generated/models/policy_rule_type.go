package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	RuleUUID     string          `json:"rule_uuid,omitempty"`
	DSTPorts     []*PortType     `json:"dst_ports,omitempty"`
	Application  []string        `json:"application,omitempty"`
	LastModified string          `json:"last_modified,omitempty"`
	Ethertype    EtherType       `json:"ethertype,omitempty"`
	SRCAddresses []*AddressType  `json:"src_addresses,omitempty"`
	DSTAddresses []*AddressType  `json:"dst_addresses,omitempty"`
	Created      string          `json:"created,omitempty"`
	ActionList   *ActionListType `json:"action_list,omitempty"`
	RuleSequence *SequenceType   `json:"rule_sequence,omitempty"`
	SRCPorts     []*PortType     `json:"src_ports,omitempty"`
	Direction    DirectionType   `json:"direction,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
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
		RuleUUID: "",

		DSTPorts: MakePortTypeSlice(),

		Application:  []string{},
		LastModified: "",
		Ethertype:    MakeEtherType(),

		SRCAddresses: MakeAddressTypeSlice(),

		DSTAddresses: MakeAddressTypeSlice(),

		Created:      "",
		ActionList:   MakeActionListType(),
		RuleSequence: MakeSequenceType(),

		SRCPorts: MakePortTypeSlice(),

		Direction: MakeDirectionType(),
		Protocol:  "",
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
