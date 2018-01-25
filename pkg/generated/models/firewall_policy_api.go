package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateFirewallPolicyRequest struct {
	FirewallPolicy *FirewallPolicy `json:"firewall-policy"`
}

type CreateFirewallPolicyResponse struct {
	FirewallPolicy *FirewallPolicy `json:"firewall-policy"`
}

type UpdateFirewallPolicyRequest struct {
	FirewallPolicy *FirewallPolicy `json:"firewall-policy"`
	FieldMask      types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateFirewallPolicyResponse struct {
	FirewallPolicy *FirewallPolicy `json:"firewall-policy"`
}

type DeleteFirewallPolicyRequest struct {
	ID string `json:"id"`
}

type DeleteFirewallPolicyResponse struct {
	ID string `json:"id"`
}

type ListFirewallPolicyRequest struct {
	Spec *ListSpec
}

type ListFirewallPolicyResponse struct {
	FirewallPolicys []*FirewallPolicy `json:"firewall-policys"`
}

type GetFirewallPolicyRequest struct {
	ID string `json:"id"`
}

type GetFirewallPolicyResponse struct {
	FirewallPolicy *FirewallPolicy `json:"firewall-policy"`
}

func InterfaceToUpdateFirewallPolicyRequest(i interface{}) *UpdateFirewallPolicyRequest {
	//TODO implement
	return &UpdateFirewallPolicyRequest{}
}
