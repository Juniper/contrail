package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	ActionList   *ActionListType `json:"action_list,omitempty"`
	Created      string          `json:"created,omitempty"`
	DSTPorts     []*PortType     `json:"dst_ports,omitempty"`
	Application  []string        `json:"application,omitempty"`
	LastModified string          `json:"last_modified,omitempty"`
	SRCPorts     []*PortType     `json:"src_ports,omitempty"`
	Direction    DirectionType   `json:"direction,omitempty"`
	DSTAddresses []*AddressType  `json:"dst_addresses,omitempty"`
	Ethertype    EtherType       `json:"ethertype,omitempty"`
	SRCAddresses []*AddressType  `json:"src_addresses,omitempty"`
	RuleSequence *SequenceType   `json:"rule_sequence,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
	RuleUUID     string          `json:"rule_uuid,omitempty"`
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
		Created: "",

		DSTPorts: MakePortTypeSlice(),

		Application:  []string{},
		LastModified: "",

		SRCPorts: MakePortTypeSlice(),

		Direction: MakeDirectionType(),

		DSTAddresses: MakeAddressTypeSlice(),

		ActionList: MakeActionListType(),

		SRCAddresses: MakeAddressTypeSlice(),

		RuleSequence: MakeSequenceType(),
		Protocol:     "",
		RuleUUID:     "",
		Ethertype:    MakeEtherType(),
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
