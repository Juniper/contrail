package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateVirtualDNSRequest struct {
	VirtualDNS *VirtualDNS `json:"virtual-DNS"`
}

type CreateVirtualDNSResponse struct {
	VirtualDNS *VirtualDNS `json:"virtual-DNS"`
}

type UpdateVirtualDNSRequest struct {
	VirtualDNS *VirtualDNS     `json:"virtual-DNS"`
	FieldMask  types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateVirtualDNSResponse struct {
	VirtualDNS *VirtualDNS `json:"virtual-DNS"`
}

type DeleteVirtualDNSRequest struct {
	ID string `json:"id"`
}

type DeleteVirtualDNSResponse struct {
	ID string `json:"id"`
}

type ListVirtualDNSRequest struct {
	Spec *ListSpec
}

type ListVirtualDNSResponse struct {
	VirtualDNSs []*VirtualDNS `json:"virtual-DNSs"`
}

type GetVirtualDNSRequest struct {
	ID string `json:"id"`
}

type GetVirtualDNSResponse struct {
	VirtualDNS *VirtualDNS `json:"virtual-DNS"`
}

func InterfaceToUpdateVirtualDNSRequest(i interface{}) *UpdateVirtualDNSRequest {
	//TODO implement
	return &UpdateVirtualDNSRequest{}
}
