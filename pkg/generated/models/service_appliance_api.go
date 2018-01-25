package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateServiceApplianceRequest struct {
	ServiceAppliance *ServiceAppliance `json:"service-appliance"`
}

type CreateServiceApplianceResponse struct {
	ServiceAppliance *ServiceAppliance `json:"service-appliance"`
}

type UpdateServiceApplianceRequest struct {
	ServiceAppliance *ServiceAppliance `json:"service-appliance"`
	FieldMask        types.FieldMask   `json:"field_mask,omitempty"`
}

type UpdateServiceApplianceResponse struct {
	ServiceAppliance *ServiceAppliance `json:"service-appliance"`
}

type DeleteServiceApplianceRequest struct {
	ID string `json:"id"`
}

type DeleteServiceApplianceResponse struct {
	ID string `json:"id"`
}

type ListServiceApplianceRequest struct {
	Spec *ListSpec
}

type ListServiceApplianceResponse struct {
	ServiceAppliances []*ServiceAppliance `json:"service-appliances"`
}

type GetServiceApplianceRequest struct {
	ID string `json:"id"`
}

type GetServiceApplianceResponse struct {
	ServiceAppliance *ServiceAppliance `json:"service-appliance"`
}

func InterfaceToUpdateServiceApplianceRequest(i interface{}) *UpdateServiceApplianceRequest {
	//TODO implement
	return &UpdateServiceApplianceRequest{}
}
