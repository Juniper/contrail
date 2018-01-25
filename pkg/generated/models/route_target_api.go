package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateRouteTargetRequest struct {
	RouteTarget *RouteTarget `json:"route-target"`
}

type CreateRouteTargetResponse struct {
	RouteTarget *RouteTarget `json:"route-target"`
}

type UpdateRouteTargetRequest struct {
	RouteTarget *RouteTarget    `json:"route-target"`
	FieldMask   types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateRouteTargetResponse struct {
	RouteTarget *RouteTarget `json:"route-target"`
}

type DeleteRouteTargetRequest struct {
	ID string `json:"id"`
}

type DeleteRouteTargetResponse struct {
	ID string `json:"id"`
}

type ListRouteTargetRequest struct {
	Spec *ListSpec
}

type ListRouteTargetResponse struct {
	RouteTargets []*RouteTarget `json:"route-targets"`
}

type GetRouteTargetRequest struct {
	ID string `json:"id"`
}

type GetRouteTargetResponse struct {
	RouteTarget *RouteTarget `json:"route-target"`
}

func InterfaceToUpdateRouteTargetRequest(i interface{}) *UpdateRouteTargetRequest {
	//TODO implement
	return &UpdateRouteTargetRequest{}
}
