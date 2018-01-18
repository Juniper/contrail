package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	Application  []string        `json:"application,omitempty"`
	Ethertype    EtherType       `json:"ethertype,omitempty"`
	RuleSequence *SequenceType   `json:"rule_sequence,omitempty"`
	Created      string          `json:"created,omitempty"`
	RuleUUID     string          `json:"rule_uuid,omitempty"`
	DSTPorts     []*PortType     `json:"dst_ports,omitempty"`
	LastModified string          `json:"last_modified,omitempty"`
	Direction    DirectionType   `json:"direction,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
	DSTAddresses []*AddressType  `json:"dst_addresses,omitempty"`
	ActionList   *ActionListType `json:"action_list,omitempty"`
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
		Created:  "",
		RuleUUID: "",

		DSTPorts: MakePortTypeSlice(),

		LastModified: "",
		Direction:    MakeDirectionType(),
		Protocol:     "",

		DSTAddresses: MakeAddressTypeSlice(),

		ActionList: MakeActionListType(),

		SRCAddresses: MakeAddressTypeSlice(),

		SRCPorts: MakePortTypeSlice(),

		Application:  []string{},
		Ethertype:    MakeEtherType(),
		RuleSequence: MakeSequenceType(),
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
