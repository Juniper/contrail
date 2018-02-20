package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateBaremetalPortRequest struct {
	BaremetalPort *BaremetalPort `json:"baremetal-port"`
}

type CreateBaremetalPortResponse struct {
	BaremetalPort *BaremetalPort `json:"baremetal-port"`
}

type UpdateBaremetalPortRequest struct {
	BaremetalPort *BaremetalPort  `json:"baremetal-port"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateBaremetalPortResponse struct {
	BaremetalPort *BaremetalPort `json:"baremetal-port"`
}

type DeleteBaremetalPortRequest struct {
	ID string `json:"id"`
}

type DeleteBaremetalPortResponse struct {
	ID string `json:"id"`
}

type ListBaremetalPortRequest struct {
	Spec *ListSpec
}

type ListBaremetalPortResponse struct {
	BaremetalPorts []*BaremetalPort `json:"baremetal-ports"`
}

type GetBaremetalPortRequest struct {
	ID string `json:"id"`
}

type GetBaremetalPortResponse struct {
	BaremetalPort *BaremetalPort `json:"baremetal-port"`
}

func InterfaceToUpdateBaremetalPortRequest(i interface{}) *UpdateBaremetalPortRequest {
	//TODO implement
	return &UpdateBaremetalPortRequest{}
}
