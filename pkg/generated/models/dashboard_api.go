package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateDashboardRequest struct {
	Dashboard *Dashboard `json:"dashboard"`
}

type CreateDashboardResponse struct {
	Dashboard *Dashboard `json:"dashboard"`
}

type UpdateDashboardRequest struct {
	Dashboard *Dashboard      `json:"dashboard"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateDashboardResponse struct {
	Dashboard *Dashboard `json:"dashboard"`
}

type DeleteDashboardRequest struct {
	ID string `json:"id"`
}

type DeleteDashboardResponse struct {
	ID string `json:"id"`
}

type ListDashboardRequest struct {
	Spec *ListSpec
}

type ListDashboardResponse struct {
	Dashboards []*Dashboard `json:"dashboards"`
}

type GetDashboardRequest struct {
	ID string `json:"id"`
}

type GetDashboardResponse struct {
	Dashboard *Dashboard `json:"dashboard"`
}

func InterfaceToUpdateDashboardRequest(i interface{}) *UpdateDashboardRequest {
	//TODO implement
	return &UpdateDashboardRequest{}
}
