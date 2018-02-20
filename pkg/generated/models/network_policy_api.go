package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateNetworkPolicyRequest struct {
	NetworkPolicy *NetworkPolicy `json:"network-policy"`
}

type CreateNetworkPolicyResponse struct {
	NetworkPolicy *NetworkPolicy `json:"network-policy"`
}

type UpdateNetworkPolicyRequest struct {
	NetworkPolicy *NetworkPolicy  `json:"network-policy"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateNetworkPolicyResponse struct {
	NetworkPolicy *NetworkPolicy `json:"network-policy"`
}

type DeleteNetworkPolicyRequest struct {
	ID string `json:"id"`
}

type DeleteNetworkPolicyResponse struct {
	ID string `json:"id"`
}

type ListNetworkPolicyRequest struct {
	Spec *ListSpec
}

type ListNetworkPolicyResponse struct {
	NetworkPolicys []*NetworkPolicy `json:"network-policys"`
}

type GetNetworkPolicyRequest struct {
	ID string `json:"id"`
}

type GetNetworkPolicyResponse struct {
	NetworkPolicy *NetworkPolicy `json:"network-policy"`
}

func InterfaceToUpdateNetworkPolicyRequest(i interface{}) *UpdateNetworkPolicyRequest {
	//TODO implement
	return &UpdateNetworkPolicyRequest{}
}
