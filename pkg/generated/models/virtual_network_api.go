package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateVirtualNetworkRequest struct {
	VirtualNetwork *VirtualNetwork `json:"virtual-network"`
}

type CreateVirtualNetworkResponse struct {
	VirtualNetwork *VirtualNetwork `json:"virtual-network"`
}

type UpdateVirtualNetworkRequest struct {
	VirtualNetwork *VirtualNetwork `json:"virtual-network"`
	FieldMask      types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateVirtualNetworkResponse struct {
	VirtualNetwork *VirtualNetwork `json:"virtual-network"`
}

type DeleteVirtualNetworkRequest struct {
	ID string `json:"id"`
}

type DeleteVirtualNetworkResponse struct {
	ID string `json:"id"`
}

type ListVirtualNetworkRequest struct {
	Spec *ListSpec
}

type ListVirtualNetworkResponse struct {
	VirtualNetworks []*VirtualNetwork `json:"virtual-networks"`
}

type GetVirtualNetworkRequest struct {
	ID string `json:"id"`
}

type GetVirtualNetworkResponse struct {
	VirtualNetwork *VirtualNetwork `json:"virtual-network"`
}

func InterfaceToUpdateVirtualNetworkRequest(i interface{}) *UpdateVirtualNetworkRequest {
	//TODO implement
	return &UpdateVirtualNetworkRequest{}
}
