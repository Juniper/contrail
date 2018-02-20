package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateVirtualRouterRequest struct {
	VirtualRouter *VirtualRouter `json:"virtual-router"`
}

type CreateVirtualRouterResponse struct {
	VirtualRouter *VirtualRouter `json:"virtual-router"`
}

type UpdateVirtualRouterRequest struct {
	VirtualRouter *VirtualRouter  `json:"virtual-router"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateVirtualRouterResponse struct {
	VirtualRouter *VirtualRouter `json:"virtual-router"`
}

type DeleteVirtualRouterRequest struct {
	ID string `json:"id"`
}

type DeleteVirtualRouterResponse struct {
	ID string `json:"id"`
}

type ListVirtualRouterRequest struct {
	Spec *ListSpec
}

type ListVirtualRouterResponse struct {
	VirtualRouters []*VirtualRouter `json:"virtual-routers"`
}

type GetVirtualRouterRequest struct {
	ID string `json:"id"`
}

type GetVirtualRouterResponse struct {
	VirtualRouter *VirtualRouter `json:"virtual-router"`
}

func InterfaceToUpdateVirtualRouterRequest(i interface{}) *UpdateVirtualRouterRequest {
	//TODO implement
	return &UpdateVirtualRouterRequest{}
}
