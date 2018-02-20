package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateRoutingPolicyRequest struct {
	RoutingPolicy *RoutingPolicy `json:"routing-policy"`
}

type CreateRoutingPolicyResponse struct {
	RoutingPolicy *RoutingPolicy `json:"routing-policy"`
}

type UpdateRoutingPolicyRequest struct {
	RoutingPolicy *RoutingPolicy  `json:"routing-policy"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateRoutingPolicyResponse struct {
	RoutingPolicy *RoutingPolicy `json:"routing-policy"`
}

type DeleteRoutingPolicyRequest struct {
	ID string `json:"id"`
}

type DeleteRoutingPolicyResponse struct {
	ID string `json:"id"`
}

type ListRoutingPolicyRequest struct {
	Spec *ListSpec
}

type ListRoutingPolicyResponse struct {
	RoutingPolicys []*RoutingPolicy `json:"routing-policys"`
}

type GetRoutingPolicyRequest struct {
	ID string `json:"id"`
}

type GetRoutingPolicyResponse struct {
	RoutingPolicy *RoutingPolicy `json:"routing-policy"`
}

func InterfaceToUpdateRoutingPolicyRequest(i interface{}) *UpdateRoutingPolicyRequest {
	//TODO implement
	return &UpdateRoutingPolicyRequest{}
}
