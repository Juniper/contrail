package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	DSTAddresses []*AddressType  `json:"dst_addresses,omitempty"`
	ActionList   *ActionListType `json:"action_list,omitempty"`
	Created      string          `json:"created,omitempty"`
	Application  []string        `json:"application,omitempty"`
	Ethertype    EtherType       `json:"ethertype,omitempty"`
	Direction    DirectionType   `json:"direction,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
	RuleUUID     string          `json:"rule_uuid,omitempty"`
	DSTPorts     []*PortType     `json:"dst_ports,omitempty"`
	LastModified string          `json:"last_modified,omitempty"`
	SRCAddresses []*AddressType  `json:"src_addresses,omitempty"`
	RuleSequence *SequenceType   `json:"rule_sequence,omitempty"`
	SRCPorts     []*PortType     `json:"src_ports,omitempty"`
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

		DSTAddresses: MakeAddressTypeSlice(),

		ActionList:   MakeActionListType(),
		Created:      "",
		Application:  []string{},
		Ethertype:    MakeEtherType(),
		RuleSequence: MakeSequenceType(),

		SRCPorts: MakePortTypeSlice(),

		Direction: MakeDirectionType(),
		Protocol:  "",
		RuleUUID:  "",

		DSTPorts: MakePortTypeSlice(),

		LastModified: "",

		SRCAddresses: MakeAddressTypeSlice(),
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
