package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	Direction    DirectionType   `json:"direction,omitempty"`
	ActionList   *ActionListType `json:"action_list,omitempty"`
	Application  []string        `json:"application,omitempty"`
	LastModified string          `json:"last_modified,omitempty"`
	Ethertype    EtherType       `json:"ethertype,omitempty"`
	RuleSequence *SequenceType   `json:"rule_sequence,omitempty"`
	SRCPorts     []*PortType     `json:"src_ports,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
	DSTAddresses []*AddressType  `json:"dst_addresses,omitempty"`
	Created      string          `json:"created,omitempty"`
	RuleUUID     string          `json:"rule_uuid,omitempty"`
	DSTPorts     []*PortType     `json:"dst_ports,omitempty"`
	SRCAddresses []*AddressType  `json:"src_addresses,omitempty"`
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

		Created:  "",
		RuleUUID: "",

		DSTPorts: MakePortTypeSlice(),

		SRCAddresses: MakeAddressTypeSlice(),

		SRCPorts: MakePortTypeSlice(),

		Direction:    MakeDirectionType(),
		ActionList:   MakeActionListType(),
		Application:  []string{},
		LastModified: "",
		Ethertype:    MakeEtherType(),
		RuleSequence: MakeSequenceType(),
	}
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
