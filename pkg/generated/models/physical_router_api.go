package models

import (
	"github.com/gogo/protobuf/types"
)

type CreatePhysicalRouterRequest struct {
	PhysicalRouter *PhysicalRouter `json:"physical-router"`
}

type CreatePhysicalRouterResponse struct {
	PhysicalRouter *PhysicalRouter `json:"physical-router"`
}

type UpdatePhysicalRouterRequest struct {
	PhysicalRouter *PhysicalRouter `json:"physical-router"`
	FieldMask      types.FieldMask `json:"field_mask,omitempty"`
}

type UpdatePhysicalRouterResponse struct {
	PhysicalRouter *PhysicalRouter `json:"physical-router"`
}

type DeletePhysicalRouterRequest struct {
	ID string `json:"id"`
}

type DeletePhysicalRouterResponse struct {
	ID string `json:"id"`
}

type ListPhysicalRouterRequest struct {
	Spec *ListSpec
}

type ListPhysicalRouterResponse struct {
	PhysicalRouters []*PhysicalRouter `json:"physical-routers"`
}

type GetPhysicalRouterRequest struct {
	ID string `json:"id"`
}

type GetPhysicalRouterResponse struct {
	PhysicalRouter *PhysicalRouter `json:"physical-router"`
}

func InterfaceToUpdatePhysicalRouterRequest(i interface{}) *UpdatePhysicalRouterRequest {
	//TODO implement
	return &UpdatePhysicalRouterRequest{}
}
