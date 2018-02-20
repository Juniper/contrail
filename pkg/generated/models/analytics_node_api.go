package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateAnalyticsNodeRequest struct {
	AnalyticsNode *AnalyticsNode `json:"analytics-node"`
}

type CreateAnalyticsNodeResponse struct {
	AnalyticsNode *AnalyticsNode `json:"analytics-node"`
}

type UpdateAnalyticsNodeRequest struct {
	AnalyticsNode *AnalyticsNode  `json:"analytics-node"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateAnalyticsNodeResponse struct {
	AnalyticsNode *AnalyticsNode `json:"analytics-node"`
}

type DeleteAnalyticsNodeRequest struct {
	ID string `json:"id"`
}

type DeleteAnalyticsNodeResponse struct {
	ID string `json:"id"`
}

type ListAnalyticsNodeRequest struct {
	Spec *ListSpec
}

type ListAnalyticsNodeResponse struct {
	AnalyticsNodes []*AnalyticsNode `json:"analytics-nodes"`
}

type GetAnalyticsNodeRequest struct {
	ID string `json:"id"`
}

type GetAnalyticsNodeResponse struct {
	AnalyticsNode *AnalyticsNode `json:"analytics-node"`
}

func InterfaceToUpdateAnalyticsNodeRequest(i interface{}) *UpdateAnalyticsNodeRequest {
	//TODO implement
	return &UpdateAnalyticsNodeRequest{}
}
