package models

import (
	"github.com/gogo/protobuf/types"
)

type CreatePortTupleRequest struct {
	PortTuple *PortTuple `json:"port-tuple"`
}

type CreatePortTupleResponse struct {
	PortTuple *PortTuple `json:"port-tuple"`
}

type UpdatePortTupleRequest struct {
	PortTuple *PortTuple      `json:"port-tuple"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdatePortTupleResponse struct {
	PortTuple *PortTuple `json:"port-tuple"`
}

type DeletePortTupleRequest struct {
	ID string `json:"id"`
}

type DeletePortTupleResponse struct {
	ID string `json:"id"`
}

type ListPortTupleRequest struct {
	Spec *ListSpec
}

type ListPortTupleResponse struct {
	PortTuples []*PortTuple `json:"port-tuples"`
}

type GetPortTupleRequest struct {
	ID string `json:"id"`
}

type GetPortTupleResponse struct {
	PortTuple *PortTuple `json:"port-tuple"`
}

func InterfaceToUpdatePortTupleRequest(i interface{}) *UpdatePortTupleRequest {
	//TODO implement
	return &UpdatePortTupleRequest{}
}
