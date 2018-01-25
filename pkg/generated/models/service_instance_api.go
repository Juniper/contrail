package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateServiceInstanceRequest struct {
	ServiceInstance *ServiceInstance `json:"service-instance"`
}

type CreateServiceInstanceResponse struct {
	ServiceInstance *ServiceInstance `json:"service-instance"`
}

type UpdateServiceInstanceRequest struct {
	ServiceInstance *ServiceInstance `json:"service-instance"`
	FieldMask       types.FieldMask  `json:"field_mask,omitempty"`
}

type UpdateServiceInstanceResponse struct {
	ServiceInstance *ServiceInstance `json:"service-instance"`
}

type DeleteServiceInstanceRequest struct {
	ID string `json:"id"`
}

type DeleteServiceInstanceResponse struct {
	ID string `json:"id"`
}

type ListServiceInstanceRequest struct {
	Spec *ListSpec
}

type ListServiceInstanceResponse struct {
	ServiceInstances []*ServiceInstance `json:"service-instances"`
}

type GetServiceInstanceRequest struct {
	ID string `json:"id"`
}

type GetServiceInstanceResponse struct {
	ServiceInstance *ServiceInstance `json:"service-instance"`
}

func InterfaceToUpdateServiceInstanceRequest(i interface{}) *UpdateServiceInstanceRequest {
	//TODO implement
	return &UpdateServiceInstanceRequest{}
}
