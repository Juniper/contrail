package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateLoadbalancerPoolRequest struct {
	LoadbalancerPool *LoadbalancerPool `json:"loadbalancer-pool"`
}

type CreateLoadbalancerPoolResponse struct {
	LoadbalancerPool *LoadbalancerPool `json:"loadbalancer-pool"`
}

type UpdateLoadbalancerPoolRequest struct {
	LoadbalancerPool *LoadbalancerPool `json:"loadbalancer-pool"`
	FieldMask        types.FieldMask   `json:"field_mask,omitempty"`
}

type UpdateLoadbalancerPoolResponse struct {
	LoadbalancerPool *LoadbalancerPool `json:"loadbalancer-pool"`
}

type DeleteLoadbalancerPoolRequest struct {
	ID string `json:"id"`
}

type DeleteLoadbalancerPoolResponse struct {
	ID string `json:"id"`
}

type ListLoadbalancerPoolRequest struct {
	Spec *ListSpec
}

type ListLoadbalancerPoolResponse struct {
	LoadbalancerPools []*LoadbalancerPool `json:"loadbalancer-pools"`
}

type GetLoadbalancerPoolRequest struct {
	ID string `json:"id"`
}

type GetLoadbalancerPoolResponse struct {
	LoadbalancerPool *LoadbalancerPool `json:"loadbalancer-pool"`
}

func InterfaceToUpdateLoadbalancerPoolRequest(i interface{}) *UpdateLoadbalancerPoolRequest {
	//TODO implement
	return &UpdateLoadbalancerPoolRequest{}
}
