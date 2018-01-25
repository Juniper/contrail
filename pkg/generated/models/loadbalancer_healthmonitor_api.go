package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateLoadbalancerHealthmonitorRequest struct {
	LoadbalancerHealthmonitor *LoadbalancerHealthmonitor `json:"loadbalancer-healthmonitor"`
}

type CreateLoadbalancerHealthmonitorResponse struct {
	LoadbalancerHealthmonitor *LoadbalancerHealthmonitor `json:"loadbalancer-healthmonitor"`
}

type UpdateLoadbalancerHealthmonitorRequest struct {
	LoadbalancerHealthmonitor *LoadbalancerHealthmonitor `json:"loadbalancer-healthmonitor"`
	FieldMask                 types.FieldMask            `json:"field_mask,omitempty"`
}

type UpdateLoadbalancerHealthmonitorResponse struct {
	LoadbalancerHealthmonitor *LoadbalancerHealthmonitor `json:"loadbalancer-healthmonitor"`
}

type DeleteLoadbalancerHealthmonitorRequest struct {
	ID string `json:"id"`
}

type DeleteLoadbalancerHealthmonitorResponse struct {
	ID string `json:"id"`
}

type ListLoadbalancerHealthmonitorRequest struct {
	Spec *ListSpec
}

type ListLoadbalancerHealthmonitorResponse struct {
	LoadbalancerHealthmonitors []*LoadbalancerHealthmonitor `json:"loadbalancer-healthmonitors"`
}

type GetLoadbalancerHealthmonitorRequest struct {
	ID string `json:"id"`
}

type GetLoadbalancerHealthmonitorResponse struct {
	LoadbalancerHealthmonitor *LoadbalancerHealthmonitor `json:"loadbalancer-healthmonitor"`
}

func InterfaceToUpdateLoadbalancerHealthmonitorRequest(i interface{}) *UpdateLoadbalancerHealthmonitorRequest {
	//TODO implement
	return &UpdateLoadbalancerHealthmonitorRequest{}
}
