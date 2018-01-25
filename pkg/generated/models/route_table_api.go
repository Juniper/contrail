package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateRouteTableRequest struct {
	RouteTable *RouteTable `json:"route-table"`
}

type CreateRouteTableResponse struct {
	RouteTable *RouteTable `json:"route-table"`
}

type UpdateRouteTableRequest struct {
	RouteTable *RouteTable     `json:"route-table"`
	FieldMask  types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateRouteTableResponse struct {
	RouteTable *RouteTable `json:"route-table"`
}

type DeleteRouteTableRequest struct {
	ID string `json:"id"`
}

type DeleteRouteTableResponse struct {
	ID string `json:"id"`
}

type ListRouteTableRequest struct {
	Spec *ListSpec
}

type ListRouteTableResponse struct {
	RouteTables []*RouteTable `json:"route-tables"`
}

type GetRouteTableRequest struct {
	ID string `json:"id"`
}

type GetRouteTableResponse struct {
	RouteTable *RouteTable `json:"route-table"`
}

func InterfaceToUpdateRouteTableRequest(i interface{}) *UpdateRouteTableRequest {
	//TODO implement
	return &UpdateRouteTableRequest{}
}
