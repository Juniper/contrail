package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateFirewallRuleRequest struct {
	FirewallRule *FirewallRule `json:"firewall-rule"`
}

type CreateFirewallRuleResponse struct {
	FirewallRule *FirewallRule `json:"firewall-rule"`
}

type UpdateFirewallRuleRequest struct {
	FirewallRule *FirewallRule   `json:"firewall-rule"`
	FieldMask    types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateFirewallRuleResponse struct {
	FirewallRule *FirewallRule `json:"firewall-rule"`
}

type DeleteFirewallRuleRequest struct {
	ID string `json:"id"`
}

type DeleteFirewallRuleResponse struct {
	ID string `json:"id"`
}

type ListFirewallRuleRequest struct {
	Spec *ListSpec
}

type ListFirewallRuleResponse struct {
	FirewallRules []*FirewallRule `json:"firewall-rules"`
}

type GetFirewallRuleRequest struct {
	ID string `json:"id"`
}

type GetFirewallRuleResponse struct {
	FirewallRule *FirewallRule `json:"firewall-rule"`
}

func InterfaceToUpdateFirewallRuleRequest(i interface{}) *UpdateFirewallRuleRequest {
	//TODO implement
	return &UpdateFirewallRuleRequest{}
}
