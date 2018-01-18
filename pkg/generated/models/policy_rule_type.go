package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	Protocol     string          `json:"protocol,omitempty"`
	ActionList   *ActionListType `json:"action_list,omitempty"`
	DSTPorts     []*PortType     `json:"dst_ports,omitempty"`
	SRCAddresses []*AddressType  `json:"src_addresses,omitempty"`
	SRCPorts     []*PortType     `json:"src_ports,omitempty"`
	Direction    DirectionType   `json:"direction,omitempty"`
	DSTAddresses []*AddressType  `json:"dst_addresses,omitempty"`
	Created      string          `json:"created,omitempty"`
	RuleUUID     string          `json:"rule_uuid,omitempty"`
	Application  []string        `json:"application,omitempty"`
	LastModified string          `json:"last_modified,omitempty"`
	Ethertype    EtherType       `json:"ethertype,omitempty"`
	RuleSequence *SequenceType   `json:"rule_sequence,omitempty"`
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
		Protocol:   "",
		ActionList: MakeActionListType(),

		DSTPorts: MakePortTypeSlice(),

		SRCAddresses: MakeAddressTypeSlice(),

		SRCPorts: MakePortTypeSlice(),

		RuleSequence: MakeSequenceType(),
		Direction:    MakeDirectionType(),

		DSTAddresses: MakeAddressTypeSlice(),

		Created:      "",
		RuleUUID:     "",
		Application:  []string{},
		LastModified: "",
		Ethertype:    MakeEtherType(),
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
