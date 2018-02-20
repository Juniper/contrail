package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateServiceConnectionModuleRequest struct {
	ServiceConnectionModule *ServiceConnectionModule `json:"service-connection-module"`
}

type CreateServiceConnectionModuleResponse struct {
	ServiceConnectionModule *ServiceConnectionModule `json:"service-connection-module"`
}

type UpdateServiceConnectionModuleRequest struct {
	ServiceConnectionModule *ServiceConnectionModule `json:"service-connection-module"`
	FieldMask               types.FieldMask          `json:"field_mask,omitempty"`
}

type UpdateServiceConnectionModuleResponse struct {
	ServiceConnectionModule *ServiceConnectionModule `json:"service-connection-module"`
}

type DeleteServiceConnectionModuleRequest struct {
	ID string `json:"id"`
}

type DeleteServiceConnectionModuleResponse struct {
	ID string `json:"id"`
}

type ListServiceConnectionModuleRequest struct {
	Spec *ListSpec
}

type ListServiceConnectionModuleResponse struct {
	ServiceConnectionModules []*ServiceConnectionModule `json:"service-connection-modules"`
}

type GetServiceConnectionModuleRequest struct {
	ID string `json:"id"`
}

type GetServiceConnectionModuleResponse struct {
	ServiceConnectionModule *ServiceConnectionModule `json:"service-connection-module"`
}

func InterfaceToUpdateServiceConnectionModuleRequest(i interface{}) *UpdateServiceConnectionModuleRequest {
	//TODO implement
	return &UpdateServiceConnectionModuleRequest{}
}
