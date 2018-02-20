package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateRoutingInstanceRequest struct {
	RoutingInstance *RoutingInstance `json:"routing-instance"`
}

type CreateRoutingInstanceResponse struct {
	RoutingInstance *RoutingInstance `json:"routing-instance"`
}

type UpdateRoutingInstanceRequest struct {
	RoutingInstance *RoutingInstance `json:"routing-instance"`
	FieldMask       types.FieldMask  `json:"field_mask,omitempty"`
}

type UpdateRoutingInstanceResponse struct {
	RoutingInstance *RoutingInstance `json:"routing-instance"`
}

type DeleteRoutingInstanceRequest struct {
	ID string `json:"id"`
}

type DeleteRoutingInstanceResponse struct {
	ID string `json:"id"`
}

type ListRoutingInstanceRequest struct {
	Spec *ListSpec
}

type ListRoutingInstanceResponse struct {
	RoutingInstances []*RoutingInstance `json:"routing-instances"`
}

type GetRoutingInstanceRequest struct {
	ID string `json:"id"`
}

type GetRoutingInstanceResponse struct {
	RoutingInstance *RoutingInstance `json:"routing-instance"`
}

func InterfaceToUpdateRoutingInstanceRequest(i interface{}) *UpdateRoutingInstanceRequest {
	//TODO implement
	return &UpdateRoutingInstanceRequest{}
}
