package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateE2ServiceProviderRequest struct {
	E2ServiceProvider *E2ServiceProvider `json:"e2-service-provider"`
}

type CreateE2ServiceProviderResponse struct {
	E2ServiceProvider *E2ServiceProvider `json:"e2-service-provider"`
}

type UpdateE2ServiceProviderRequest struct {
	E2ServiceProvider *E2ServiceProvider `json:"e2-service-provider"`
	FieldMask         types.FieldMask    `json:"field_mask,omitempty"`
}

type UpdateE2ServiceProviderResponse struct {
	E2ServiceProvider *E2ServiceProvider `json:"e2-service-provider"`
}

type DeleteE2ServiceProviderRequest struct {
	ID string `json:"id"`
}

type DeleteE2ServiceProviderResponse struct {
	ID string `json:"id"`
}

type ListE2ServiceProviderRequest struct {
	Spec *ListSpec
}

type ListE2ServiceProviderResponse struct {
	E2ServiceProviders []*E2ServiceProvider `json:"e2-service-providers"`
}

type GetE2ServiceProviderRequest struct {
	ID string `json:"id"`
}

type GetE2ServiceProviderResponse struct {
	E2ServiceProvider *E2ServiceProvider `json:"e2-service-provider"`
}

func InterfaceToUpdateE2ServiceProviderRequest(i interface{}) *UpdateE2ServiceProviderRequest {
	//TODO implement
	return &UpdateE2ServiceProviderRequest{}
}
