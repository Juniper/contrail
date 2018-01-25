package models

import (
	"github.com/gogo/protobuf/types"
)

type CreatePhysicalInterfaceRequest struct {
	PhysicalInterface *PhysicalInterface `json:"physical-interface"`
}

type CreatePhysicalInterfaceResponse struct {
	PhysicalInterface *PhysicalInterface `json:"physical-interface"`
}

type UpdatePhysicalInterfaceRequest struct {
	PhysicalInterface *PhysicalInterface `json:"physical-interface"`
	FieldMask         types.FieldMask    `json:"field_mask,omitempty"`
}

type UpdatePhysicalInterfaceResponse struct {
	PhysicalInterface *PhysicalInterface `json:"physical-interface"`
}

type DeletePhysicalInterfaceRequest struct {
	ID string `json:"id"`
}

type DeletePhysicalInterfaceResponse struct {
	ID string `json:"id"`
}

type ListPhysicalInterfaceRequest struct {
	Spec *ListSpec
}

type ListPhysicalInterfaceResponse struct {
	PhysicalInterfaces []*PhysicalInterface `json:"physical-interfaces"`
}

type GetPhysicalInterfaceRequest struct {
	ID string `json:"id"`
}

type GetPhysicalInterfaceResponse struct {
	PhysicalInterface *PhysicalInterface `json:"physical-interface"`
}

func InterfaceToUpdatePhysicalInterfaceRequest(i interface{}) *UpdatePhysicalInterfaceRequest {
	//TODO implement
	return &UpdatePhysicalInterfaceRequest{}
}
