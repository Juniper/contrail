package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	SRCPorts     []*PortType     `json:"src_ports,omitempty"`
	DSTAddresses []*AddressType  `json:"dst_addresses,omitempty"`
	DSTPorts     []*PortType     `json:"dst_ports,omitempty"`
	LastModified string          `json:"last_modified,omitempty"`
	Ethertype    EtherType       `json:"ethertype,omitempty"`
	SRCAddresses []*AddressType  `json:"src_addresses,omitempty"`
	RuleSequence *SequenceType   `json:"rule_sequence,omitempty"`
	Direction    DirectionType   `json:"direction,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
	ActionList   *ActionListType `json:"action_list,omitempty"`
	Created      string          `json:"created,omitempty"`
	RuleUUID     string          `json:"rule_uuid,omitempty"`
	Application  []string        `json:"application,omitempty"`
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
		RuleSequence: MakeSequenceType(),

		SRCPorts: MakePortTypeSlice(),

		DSTAddresses: MakeAddressTypeSlice(),

		DSTPorts: MakePortTypeSlice(),

		LastModified: "",
		Ethertype:    MakeEtherType(),

		SRCAddresses: MakeAddressTypeSlice(),

		Application: []string{},
		Direction:   MakeDirectionType(),
		Protocol:    "",
		ActionList:  MakeActionListType(),
		Created:     "",
		RuleUUID:    "",
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
