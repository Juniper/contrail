package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateLocationRequest struct {
	Location *Location `json:"location"`
}

type CreateLocationResponse struct {
	Location *Location `json:"location"`
}

type UpdateLocationRequest struct {
	Location  *Location       `json:"location"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateLocationResponse struct {
	Location *Location `json:"location"`
}

type DeleteLocationRequest struct {
	ID string `json:"id"`
}

type DeleteLocationResponse struct {
	ID string `json:"id"`
}

type ListLocationRequest struct {
	Spec *ListSpec
}

type ListLocationResponse struct {
	Locations []*Location `json:"locations"`
}

type GetLocationRequest struct {
	ID string `json:"id"`
}

type GetLocationResponse struct {
	Location *Location `json:"location"`
}

func InterfaceToUpdateLocationRequest(i interface{}) *UpdateLocationRequest {
	//TODO implement
	return &UpdateLocationRequest{}
}
