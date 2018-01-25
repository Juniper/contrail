package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateAPIAccessListRequest struct {
	APIAccessList *APIAccessList `json:"api-access-list"`
}

type CreateAPIAccessListResponse struct {
	APIAccessList *APIAccessList `json:"api-access-list"`
}

type UpdateAPIAccessListRequest struct {
	APIAccessList *APIAccessList  `json:"api-access-list"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateAPIAccessListResponse struct {
	APIAccessList *APIAccessList `json:"api-access-list"`
}

type DeleteAPIAccessListRequest struct {
	ID string `json:"id"`
}

type DeleteAPIAccessListResponse struct {
	ID string `json:"id"`
}

type ListAPIAccessListRequest struct {
	Spec *ListSpec
}

type ListAPIAccessListResponse struct {
	APIAccessLists []*APIAccessList `json:"api-access-lists"`
}

type GetAPIAccessListRequest struct {
	ID string `json:"id"`
}

type GetAPIAccessListResponse struct {
	APIAccessList *APIAccessList `json:"api-access-list"`
}

func InterfaceToUpdateAPIAccessListRequest(i interface{}) *UpdateAPIAccessListRequest {
	//TODO implement
	return &UpdateAPIAccessListRequest{}
}
