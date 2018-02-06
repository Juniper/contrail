package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateOsImageRequest struct {
	OsImage *OsImage `json:"os-image"`
}

type CreateOsImageResponse struct {
	OsImage *OsImage `json:"os-image"`
}

type UpdateOsImageRequest struct {
	OsImage   *OsImage        `json:"os-image"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateOsImageResponse struct {
	OsImage *OsImage `json:"os-image"`
}

type DeleteOsImageRequest struct {
	ID string `json:"id"`
}

type DeleteOsImageResponse struct {
	ID string `json:"id"`
}

type ListOsImageRequest struct {
	Spec *ListSpec
}

type ListOsImageResponse struct {
	OsImages []*OsImage `json:"os-images"`
}

type GetOsImageRequest struct {
	ID string `json:"id"`
}

type GetOsImageResponse struct {
	OsImage *OsImage `json:"os-image"`
}

func InterfaceToUpdateOsImageRequest(i interface{}) *UpdateOsImageRequest {
	//TODO implement
	return &UpdateOsImageRequest{}
}
