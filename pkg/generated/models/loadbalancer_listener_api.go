package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateLoadbalancerListenerRequest struct {
	LoadbalancerListener *LoadbalancerListener `json:"loadbalancer-listener"`
}

type CreateLoadbalancerListenerResponse struct {
	LoadbalancerListener *LoadbalancerListener `json:"loadbalancer-listener"`
}

type UpdateLoadbalancerListenerRequest struct {
	LoadbalancerListener *LoadbalancerListener `json:"loadbalancer-listener"`
	FieldMask            types.FieldMask       `json:"field_mask,omitempty"`
}

type UpdateLoadbalancerListenerResponse struct {
	LoadbalancerListener *LoadbalancerListener `json:"loadbalancer-listener"`
}

type DeleteLoadbalancerListenerRequest struct {
	ID string `json:"id"`
}

type DeleteLoadbalancerListenerResponse struct {
	ID string `json:"id"`
}

type ListLoadbalancerListenerRequest struct {
	Spec *ListSpec
}

type ListLoadbalancerListenerResponse struct {
	LoadbalancerListeners []*LoadbalancerListener `json:"loadbalancer-listeners"`
}

type GetLoadbalancerListenerRequest struct {
	ID string `json:"id"`
}

type GetLoadbalancerListenerResponse struct {
	LoadbalancerListener *LoadbalancerListener `json:"loadbalancer-listener"`
}

func InterfaceToUpdateLoadbalancerListenerRequest(i interface{}) *UpdateLoadbalancerListenerRequest {
	//TODO implement
	return &UpdateLoadbalancerListenerRequest{}
}
