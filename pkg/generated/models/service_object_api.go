package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateServiceObjectRequest struct {
	ServiceObject *ServiceObject `json:"service-object"`
}

type CreateServiceObjectResponse struct {
	ServiceObject *ServiceObject `json:"service-object"`
}

type UpdateServiceObjectRequest struct {
	ServiceObject *ServiceObject  `json:"service-object"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateServiceObjectResponse struct {
	ServiceObject *ServiceObject `json:"service-object"`
}

type DeleteServiceObjectRequest struct {
	ID string `json:"id"`
}

type DeleteServiceObjectResponse struct {
	ID string `json:"id"`
}

type ListServiceObjectRequest struct {
	Spec *ListSpec
}

type ListServiceObjectResponse struct {
	ServiceObjects []*ServiceObject `json:"service-objects"`
}

type GetServiceObjectRequest struct {
	ID string `json:"id"`
}

type GetServiceObjectResponse struct {
	ServiceObject *ServiceObject `json:"service-object"`
}

func InterfaceToUpdateServiceObjectRequest(i interface{}) *UpdateServiceObjectRequest {
	//TODO implement
	return &UpdateServiceObjectRequest{}
}
