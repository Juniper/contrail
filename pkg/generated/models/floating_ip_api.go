package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateFloatingIPRequest struct {
	FloatingIP *FloatingIP `json:"floating-ip"`
}

type CreateFloatingIPResponse struct {
	FloatingIP *FloatingIP `json:"floating-ip"`
}

type UpdateFloatingIPRequest struct {
	FloatingIP *FloatingIP     `json:"floating-ip"`
	FieldMask  types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateFloatingIPResponse struct {
	FloatingIP *FloatingIP `json:"floating-ip"`
}

type DeleteFloatingIPRequest struct {
	ID string `json:"id"`
}

type DeleteFloatingIPResponse struct {
	ID string `json:"id"`
}

type ListFloatingIPRequest struct {
	Spec *ListSpec
}

type ListFloatingIPResponse struct {
	FloatingIPs []*FloatingIP `json:"floating-ips"`
}

type GetFloatingIPRequest struct {
	ID string `json:"id"`
}

type GetFloatingIPResponse struct {
	FloatingIP *FloatingIP `json:"floating-ip"`
}

func InterfaceToUpdateFloatingIPRequest(i interface{}) *UpdateFloatingIPRequest {
	//TODO implement
	return &UpdateFloatingIPRequest{}
}
