package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateLogicalRouterRequest struct {
	LogicalRouter *LogicalRouter `json:"logical-router"`
}

type CreateLogicalRouterResponse struct {
	LogicalRouter *LogicalRouter `json:"logical-router"`
}

type UpdateLogicalRouterRequest struct {
	LogicalRouter *LogicalRouter  `json:"logical-router"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateLogicalRouterResponse struct {
	LogicalRouter *LogicalRouter `json:"logical-router"`
}

type DeleteLogicalRouterRequest struct {
	ID string `json:"id"`
}

type DeleteLogicalRouterResponse struct {
	ID string `json:"id"`
}

type ListLogicalRouterRequest struct {
	Spec *ListSpec
}

type ListLogicalRouterResponse struct {
	LogicalRouters []*LogicalRouter `json:"logical-routers"`
}

type GetLogicalRouterRequest struct {
	ID string `json:"id"`
}

type GetLogicalRouterResponse struct {
	LogicalRouter *LogicalRouter `json:"logical-router"`
}

func InterfaceToUpdateLogicalRouterRequest(i interface{}) *UpdateLogicalRouterRequest {
	//TODO implement
	return &UpdateLogicalRouterRequest{}
}
