package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateServiceHealthCheckRequest struct {
	ServiceHealthCheck *ServiceHealthCheck `json:"service-health-check"`
}

type CreateServiceHealthCheckResponse struct {
	ServiceHealthCheck *ServiceHealthCheck `json:"service-health-check"`
}

type UpdateServiceHealthCheckRequest struct {
	ServiceHealthCheck *ServiceHealthCheck `json:"service-health-check"`
	FieldMask          types.FieldMask     `json:"field_mask,omitempty"`
}

type UpdateServiceHealthCheckResponse struct {
	ServiceHealthCheck *ServiceHealthCheck `json:"service-health-check"`
}

type DeleteServiceHealthCheckRequest struct {
	ID string `json:"id"`
}

type DeleteServiceHealthCheckResponse struct {
	ID string `json:"id"`
}

type ListServiceHealthCheckRequest struct {
	Spec *ListSpec
}

type ListServiceHealthCheckResponse struct {
	ServiceHealthChecks []*ServiceHealthCheck `json:"service-health-checks"`
}

type GetServiceHealthCheckRequest struct {
	ID string `json:"id"`
}

type GetServiceHealthCheckResponse struct {
	ServiceHealthCheck *ServiceHealthCheck `json:"service-health-check"`
}

func InterfaceToUpdateServiceHealthCheckRequest(i interface{}) *UpdateServiceHealthCheckRequest {
	//TODO implement
	return &UpdateServiceHealthCheckRequest{}
}
