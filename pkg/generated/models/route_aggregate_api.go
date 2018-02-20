package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateRouteAggregateRequest struct {
	RouteAggregate *RouteAggregate `json:"route-aggregate"`
}

type CreateRouteAggregateResponse struct {
	RouteAggregate *RouteAggregate `json:"route-aggregate"`
}

type UpdateRouteAggregateRequest struct {
	RouteAggregate *RouteAggregate `json:"route-aggregate"`
	FieldMask      types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateRouteAggregateResponse struct {
	RouteAggregate *RouteAggregate `json:"route-aggregate"`
}

type DeleteRouteAggregateRequest struct {
	ID string `json:"id"`
}

type DeleteRouteAggregateResponse struct {
	ID string `json:"id"`
}

type ListRouteAggregateRequest struct {
	Spec *ListSpec
}

type ListRouteAggregateResponse struct {
	RouteAggregates []*RouteAggregate `json:"route-aggregates"`
}

type GetRouteAggregateRequest struct {
	ID string `json:"id"`
}

type GetRouteAggregateResponse struct {
	RouteAggregate *RouteAggregate `json:"route-aggregate"`
}

func InterfaceToUpdateRouteAggregateRequest(i interface{}) *UpdateRouteAggregateRequest {
	//TODO implement
	return &UpdateRouteAggregateRequest{}
}
