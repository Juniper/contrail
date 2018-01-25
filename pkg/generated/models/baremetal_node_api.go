package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateBaremetalNodeRequest struct {
	BaremetalNode *BaremetalNode `json:"baremetal-node"`
}

type CreateBaremetalNodeResponse struct {
	BaremetalNode *BaremetalNode `json:"baremetal-node"`
}

type UpdateBaremetalNodeRequest struct {
	BaremetalNode *BaremetalNode  `json:"baremetal-node"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateBaremetalNodeResponse struct {
	BaremetalNode *BaremetalNode `json:"baremetal-node"`
}

type DeleteBaremetalNodeRequest struct {
	ID string `json:"id"`
}

type DeleteBaremetalNodeResponse struct {
	ID string `json:"id"`
}

type ListBaremetalNodeRequest struct {
	Spec *ListSpec
}

type ListBaremetalNodeResponse struct {
	BaremetalNodes []*BaremetalNode `json:"baremetal-nodes"`
}

type GetBaremetalNodeRequest struct {
	ID string `json:"id"`
}

type GetBaremetalNodeResponse struct {
	BaremetalNode *BaremetalNode `json:"baremetal-node"`
}

func InterfaceToUpdateBaremetalNodeRequest(i interface{}) *UpdateBaremetalNodeRequest {
	//TODO implement
	return &UpdateBaremetalNodeRequest{}
}
