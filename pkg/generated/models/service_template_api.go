package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateServiceTemplateRequest struct {
	ServiceTemplate *ServiceTemplate `json:"service-template"`
}

type CreateServiceTemplateResponse struct {
	ServiceTemplate *ServiceTemplate `json:"service-template"`
}

type UpdateServiceTemplateRequest struct {
	ServiceTemplate *ServiceTemplate `json:"service-template"`
	FieldMask       types.FieldMask  `json:"field_mask,omitempty"`
}

type UpdateServiceTemplateResponse struct {
	ServiceTemplate *ServiceTemplate `json:"service-template"`
}

type DeleteServiceTemplateRequest struct {
	ID string `json:"id"`
}

type DeleteServiceTemplateResponse struct {
	ID string `json:"id"`
}

type ListServiceTemplateRequest struct {
	Spec *ListSpec
}

type ListServiceTemplateResponse struct {
	ServiceTemplates []*ServiceTemplate `json:"service-templates"`
}

type GetServiceTemplateRequest struct {
	ID string `json:"id"`
}

type GetServiceTemplateResponse struct {
	ServiceTemplate *ServiceTemplate `json:"service-template"`
}

func InterfaceToUpdateServiceTemplateRequest(i interface{}) *UpdateServiceTemplateRequest {
	//TODO implement
	return &UpdateServiceTemplateRequest{}
}
