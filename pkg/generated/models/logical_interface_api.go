package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateLogicalInterfaceRequest struct {
	LogicalInterface *LogicalInterface `json:"logical-interface"`
}

type CreateLogicalInterfaceResponse struct {
	LogicalInterface *LogicalInterface `json:"logical-interface"`
}

type UpdateLogicalInterfaceRequest struct {
	LogicalInterface *LogicalInterface `json:"logical-interface"`
	FieldMask        types.FieldMask   `json:"field_mask,omitempty"`
}

type UpdateLogicalInterfaceResponse struct {
	LogicalInterface *LogicalInterface `json:"logical-interface"`
}

type DeleteLogicalInterfaceRequest struct {
	ID string `json:"id"`
}

type DeleteLogicalInterfaceResponse struct {
	ID string `json:"id"`
}

type ListLogicalInterfaceRequest struct {
	Spec *ListSpec
}

type ListLogicalInterfaceResponse struct {
	LogicalInterfaces []*LogicalInterface `json:"logical-interfaces"`
}

type GetLogicalInterfaceRequest struct {
	ID string `json:"id"`
}

type GetLogicalInterfaceResponse struct {
	LogicalInterface *LogicalInterface `json:"logical-interface"`
}

func InterfaceToUpdateLogicalInterfaceRequest(i interface{}) *UpdateLogicalInterfaceRequest {
	//TODO implement
	return &UpdateLogicalInterfaceRequest{}
}
