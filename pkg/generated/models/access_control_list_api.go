package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateAccessControlListRequest struct {
	AccessControlList *AccessControlList `json:"access-control-list"`
}

type CreateAccessControlListResponse struct {
	AccessControlList *AccessControlList `json:"access-control-list"`
}

type UpdateAccessControlListRequest struct {
	AccessControlList *AccessControlList `json:"access-control-list"`
	FieldMask         types.FieldMask    `json:"field_mask,omitempty"`
}

type UpdateAccessControlListResponse struct {
	AccessControlList *AccessControlList `json:"access-control-list"`
}

type DeleteAccessControlListRequest struct {
	ID string `json:"id"`
}

type DeleteAccessControlListResponse struct {
	ID string `json:"id"`
}

type ListAccessControlListRequest struct {
	Spec *ListSpec
}

type ListAccessControlListResponse struct {
	AccessControlLists []*AccessControlList `json:"access-control-lists"`
}

type GetAccessControlListRequest struct {
	ID string `json:"id"`
}

type GetAccessControlListResponse struct {
	AccessControlList *AccessControlList `json:"access-control-list"`
}

func InterfaceToUpdateAccessControlListRequest(i interface{}) *UpdateAccessControlListRequest {
	//TODO implement
	return &UpdateAccessControlListRequest{}
}
