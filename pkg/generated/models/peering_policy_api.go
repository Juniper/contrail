package models

import (
	"github.com/gogo/protobuf/types"
)

type CreatePeeringPolicyRequest struct {
	PeeringPolicy *PeeringPolicy `json:"peering-policy"`
}

type CreatePeeringPolicyResponse struct {
	PeeringPolicy *PeeringPolicy `json:"peering-policy"`
}

type UpdatePeeringPolicyRequest struct {
	PeeringPolicy *PeeringPolicy  `json:"peering-policy"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdatePeeringPolicyResponse struct {
	PeeringPolicy *PeeringPolicy `json:"peering-policy"`
}

type DeletePeeringPolicyRequest struct {
	ID string `json:"id"`
}

type DeletePeeringPolicyResponse struct {
	ID string `json:"id"`
}

type ListPeeringPolicyRequest struct {
	Spec *ListSpec
}

type ListPeeringPolicyResponse struct {
	PeeringPolicys []*PeeringPolicy `json:"peering-policys"`
}

type GetPeeringPolicyRequest struct {
	ID string `json:"id"`
}

type GetPeeringPolicyResponse struct {
	PeeringPolicy *PeeringPolicy `json:"peering-policy"`
}

func InterfaceToUpdatePeeringPolicyRequest(i interface{}) *UpdatePeeringPolicyRequest {
	//TODO implement
	return &UpdatePeeringPolicyRequest{}
}
