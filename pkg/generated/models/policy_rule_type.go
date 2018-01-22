package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	Direction    DirectionType   `json:"direction,omitempty"`
	ActionList   *ActionListType `json:"action_list,omitempty"`
	Created      string          `json:"created,omitempty"`
	DSTPorts     []*PortType     `json:"dst_ports,omitempty"`
	Application  []string        `json:"application,omitempty"`
	Ethertype    EtherType       `json:"ethertype,omitempty"`
	RuleSequence *SequenceType   `json:"rule_sequence,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
	DSTAddresses []*AddressType  `json:"dst_addresses,omitempty"`
	RuleUUID     string          `json:"rule_uuid,omitempty"`
	LastModified string          `json:"last_modified,omitempty"`
	SRCAddresses []*AddressType  `json:"src_addresses,omitempty"`
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
		Protocol: "",

		DSTAddresses: MakeAddressTypeSlice(),

		RuleUUID:     "",
		LastModified: "",

		SRCAddresses: MakeAddressTypeSlice(),

		SRCPorts: MakePortTypeSlice(),

		Direction:  MakeDirectionType(),
		ActionList: MakeActionListType(),
		Created:    "",

		DSTPorts: MakePortTypeSlice(),

		Application:  []string{},
		Ethertype:    MakeEtherType(),
		RuleSequence: MakeSequenceType(),
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
