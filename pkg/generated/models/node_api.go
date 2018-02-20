package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateNodeRequest struct {
	Node *Node `json:"node"`
}

type CreateNodeResponse struct {
	Node *Node `json:"node"`
}

type UpdateNodeRequest struct {
	Node      *Node           `json:"node"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateNodeResponse struct {
	Node *Node `json:"node"`
}

type DeleteNodeRequest struct {
	ID string `json:"id"`
}

type DeleteNodeResponse struct {
	ID string `json:"id"`
}

type ListNodeRequest struct {
	Spec *ListSpec
}

type ListNodeResponse struct {
	Nodes []*Node `json:"nodes"`
}

type GetNodeRequest struct {
	ID string `json:"id"`
}

type GetNodeResponse struct {
	Node *Node `json:"node"`
}

func InterfaceToUpdateNodeRequest(i interface{}) *UpdateNodeRequest {
	//TODO implement
	return &UpdateNodeRequest{}
}
