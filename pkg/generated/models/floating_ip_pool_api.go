package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateFloatingIPPoolRequest struct {
	FloatingIPPool *FloatingIPPool `json:"floating-ip-pool"`
}

type CreateFloatingIPPoolResponse struct {
	FloatingIPPool *FloatingIPPool `json:"floating-ip-pool"`
}

type UpdateFloatingIPPoolRequest struct {
	FloatingIPPool *FloatingIPPool `json:"floating-ip-pool"`
	FieldMask      types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateFloatingIPPoolResponse struct {
	FloatingIPPool *FloatingIPPool `json:"floating-ip-pool"`
}

type DeleteFloatingIPPoolRequest struct {
	ID string `json:"id"`
}

type DeleteFloatingIPPoolResponse struct {
	ID string `json:"id"`
}

type ListFloatingIPPoolRequest struct {
	Spec *ListSpec
}

type ListFloatingIPPoolResponse struct {
	FloatingIPPools []*FloatingIPPool `json:"floating-ip-pools"`
}

type GetFloatingIPPoolRequest struct {
	ID string `json:"id"`
}

type GetFloatingIPPoolResponse struct {
	FloatingIPPool *FloatingIPPool `json:"floating-ip-pool"`
}

func InterfaceToUpdateFloatingIPPoolRequest(i interface{}) *UpdateFloatingIPPoolRequest {
	//TODO implement
	return &UpdateFloatingIPPoolRequest{}
}
