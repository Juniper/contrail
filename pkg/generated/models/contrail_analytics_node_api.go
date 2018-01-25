package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateContrailAnalyticsNodeRequest struct {
	ContrailAnalyticsNode *ContrailAnalyticsNode `json:"contrail-analytics-node"`
}

type CreateContrailAnalyticsNodeResponse struct {
	ContrailAnalyticsNode *ContrailAnalyticsNode `json:"contrail-analytics-node"`
}

type UpdateContrailAnalyticsNodeRequest struct {
	ContrailAnalyticsNode *ContrailAnalyticsNode `json:"contrail-analytics-node"`
	FieldMask             types.FieldMask        `json:"field_mask,omitempty"`
}

type UpdateContrailAnalyticsNodeResponse struct {
	ContrailAnalyticsNode *ContrailAnalyticsNode `json:"contrail-analytics-node"`
}

type DeleteContrailAnalyticsNodeRequest struct {
	ID string `json:"id"`
}

type DeleteContrailAnalyticsNodeResponse struct {
	ID string `json:"id"`
}

type ListContrailAnalyticsNodeRequest struct {
	Spec *ListSpec
}

type ListContrailAnalyticsNodeResponse struct {
	ContrailAnalyticsNodes []*ContrailAnalyticsNode `json:"contrail-analytics-nodes"`
}

type GetContrailAnalyticsNodeRequest struct {
	ID string `json:"id"`
}

type GetContrailAnalyticsNodeResponse struct {
	ContrailAnalyticsNode *ContrailAnalyticsNode `json:"contrail-analytics-node"`
}

func InterfaceToUpdateContrailAnalyticsNodeRequest(i interface{}) *UpdateContrailAnalyticsNodeRequest {
	//TODO implement
	return &UpdateContrailAnalyticsNodeRequest{}
}
