package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateBGPRouterRequest struct {
	BGPRouter *BGPRouter `json:"bgp-router"`
}

type CreateBGPRouterResponse struct {
	BGPRouter *BGPRouter `json:"bgp-router"`
}

type UpdateBGPRouterRequest struct {
	BGPRouter *BGPRouter      `json:"bgp-router"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateBGPRouterResponse struct {
	BGPRouter *BGPRouter `json:"bgp-router"`
}

type DeleteBGPRouterRequest struct {
	ID string `json:"id"`
}

type DeleteBGPRouterResponse struct {
	ID string `json:"id"`
}

type ListBGPRouterRequest struct {
	Spec *ListSpec
}

type ListBGPRouterResponse struct {
	BGPRouters []*BGPRouter `json:"bgp-routers"`
}

type GetBGPRouterRequest struct {
	ID string `json:"id"`
}

type GetBGPRouterResponse struct {
	BGPRouter *BGPRouter `json:"bgp-router"`
}

func InterfaceToUpdateBGPRouterRequest(i interface{}) *UpdateBGPRouterRequest {
	//TODO implement
	return &UpdateBGPRouterRequest{}
}
