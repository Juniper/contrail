package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateLoadbalancerRequest struct {
	Loadbalancer *Loadbalancer `json:"loadbalancer"`
}

type CreateLoadbalancerResponse struct {
	Loadbalancer *Loadbalancer `json:"loadbalancer"`
}

type UpdateLoadbalancerRequest struct {
	Loadbalancer *Loadbalancer   `json:"loadbalancer"`
	FieldMask    types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateLoadbalancerResponse struct {
	Loadbalancer *Loadbalancer `json:"loadbalancer"`
}

type DeleteLoadbalancerRequest struct {
	ID string `json:"id"`
}

type DeleteLoadbalancerResponse struct {
	ID string `json:"id"`
}

type ListLoadbalancerRequest struct {
	Spec *ListSpec
}

type ListLoadbalancerResponse struct {
	Loadbalancers []*Loadbalancer `json:"loadbalancers"`
}

type GetLoadbalancerRequest struct {
	ID string `json:"id"`
}

type GetLoadbalancerResponse struct {
	Loadbalancer *Loadbalancer `json:"loadbalancer"`
}

func InterfaceToUpdateLoadbalancerRequest(i interface{}) *UpdateLoadbalancerRequest {
	//TODO implement
	return &UpdateLoadbalancerRequest{}
}
