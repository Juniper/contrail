package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateVirtualIPRequest struct {
	VirtualIP *VirtualIP `json:"virtual-ip"`
}

type CreateVirtualIPResponse struct {
	VirtualIP *VirtualIP `json:"virtual-ip"`
}

type UpdateVirtualIPRequest struct {
	VirtualIP *VirtualIP      `json:"virtual-ip"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateVirtualIPResponse struct {
	VirtualIP *VirtualIP `json:"virtual-ip"`
}

type DeleteVirtualIPRequest struct {
	ID string `json:"id"`
}

type DeleteVirtualIPResponse struct {
	ID string `json:"id"`
}

type ListVirtualIPRequest struct {
	Spec *ListSpec
}

type ListVirtualIPResponse struct {
	VirtualIPs []*VirtualIP `json:"virtual-ips"`
}

type GetVirtualIPRequest struct {
	ID string `json:"id"`
}

type GetVirtualIPResponse struct {
	VirtualIP *VirtualIP `json:"virtual-ip"`
}

func InterfaceToUpdateVirtualIPRequest(i interface{}) *UpdateVirtualIPRequest {
	//TODO implement
	return &UpdateVirtualIPRequest{}
}
