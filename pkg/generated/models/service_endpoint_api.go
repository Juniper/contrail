package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateServiceEndpointRequest struct {
	ServiceEndpoint *ServiceEndpoint `json:"service-endpoint"`
}

type CreateServiceEndpointResponse struct {
	ServiceEndpoint *ServiceEndpoint `json:"service-endpoint"`
}

type UpdateServiceEndpointRequest struct {
	ServiceEndpoint *ServiceEndpoint `json:"service-endpoint"`
	FieldMask       types.FieldMask  `json:"field_mask,omitempty"`
}

type UpdateServiceEndpointResponse struct {
	ServiceEndpoint *ServiceEndpoint `json:"service-endpoint"`
}

type DeleteServiceEndpointRequest struct {
	ID string `json:"id"`
}

type DeleteServiceEndpointResponse struct {
	ID string `json:"id"`
}

type ListServiceEndpointRequest struct {
	Spec *ListSpec
}

type ListServiceEndpointResponse struct {
	ServiceEndpoints []*ServiceEndpoint `json:"service-endpoints"`
}

type GetServiceEndpointRequest struct {
	ID string `json:"id"`
}

type GetServiceEndpointResponse struct {
	ServiceEndpoint *ServiceEndpoint `json:"service-endpoint"`
}

func InterfaceToUpdateServiceEndpointRequest(i interface{}) *UpdateServiceEndpointRequest {
	//TODO implement
	return &UpdateServiceEndpointRequest{}
}
