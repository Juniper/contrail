package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateForwardingClassRequest struct {
	ForwardingClass *ForwardingClass `json:"forwarding-class"`
}

type CreateForwardingClassResponse struct {
	ForwardingClass *ForwardingClass `json:"forwarding-class"`
}

type UpdateForwardingClassRequest struct {
	ForwardingClass *ForwardingClass `json:"forwarding-class"`
	FieldMask       types.FieldMask  `json:"field_mask,omitempty"`
}

type UpdateForwardingClassResponse struct {
	ForwardingClass *ForwardingClass `json:"forwarding-class"`
}

type DeleteForwardingClassRequest struct {
	ID string `json:"id"`
}

type DeleteForwardingClassResponse struct {
	ID string `json:"id"`
}

type ListForwardingClassRequest struct {
	Spec *ListSpec
}

type ListForwardingClassResponse struct {
	ForwardingClasss []*ForwardingClass `json:"forwarding-classs"`
}

type GetForwardingClassRequest struct {
	ID string `json:"id"`
}

type GetForwardingClassResponse struct {
	ForwardingClass *ForwardingClass `json:"forwarding-class"`
}

func InterfaceToUpdateForwardingClassRequest(i interface{}) *UpdateForwardingClassRequest {
	//TODO implement
	return &UpdateForwardingClassRequest{}
}
