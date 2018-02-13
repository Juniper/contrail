package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
//proteus:generate
type PolicyRuleType struct {
	Direction    DirectionType   `json:"direction,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
	DSTAddresses []*AddressType  `json:"dst_addresses,omitempty"`
	ActionList   *ActionListType `json:"action_list,omitempty"`
	Created      string          `json:"created,omitempty"`
	RuleUUID     string          `json:"rule_uuid,omitempty"`
	DSTPorts     []*PortType     `json:"dst_ports,omitempty"`
	Application  []string        `json:"application,omitempty"`
	LastModified string          `json:"last_modified,omitempty"`
	Ethertype    EtherType       `json:"ethertype,omitempty"`
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
		Direction: MakeDirectionType(),
		Protocol:  "",

		DSTAddresses: MakeAddressTypeSlice(),

		ActionList: MakeActionListType(),
		Created:    "",
		RuleUUID:   "",

		DSTPorts: MakePortTypeSlice(),

		Application:  []string{},
		LastModified: "",
		Ethertype:    MakeEtherType(),

		SRCAddresses: MakeAddressTypeSlice(),

		RuleSequence: MakeSequenceType(),

		SRCPorts: MakePortTypeSlice(),
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
