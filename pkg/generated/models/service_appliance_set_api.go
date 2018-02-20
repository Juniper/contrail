package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateServiceApplianceSetRequest struct {
	ServiceApplianceSet *ServiceApplianceSet `json:"service-appliance-set"`
}

type CreateServiceApplianceSetResponse struct {
	ServiceApplianceSet *ServiceApplianceSet `json:"service-appliance-set"`
}

type UpdateServiceApplianceSetRequest struct {
	ServiceApplianceSet *ServiceApplianceSet `json:"service-appliance-set"`
	FieldMask           types.FieldMask      `json:"field_mask,omitempty"`
}

type UpdateServiceApplianceSetResponse struct {
	ServiceApplianceSet *ServiceApplianceSet `json:"service-appliance-set"`
}

type DeleteServiceApplianceSetRequest struct {
	ID string `json:"id"`
}

type DeleteServiceApplianceSetResponse struct {
	ID string `json:"id"`
}

type ListServiceApplianceSetRequest struct {
	Spec *ListSpec
}

type ListServiceApplianceSetResponse struct {
	ServiceApplianceSets []*ServiceApplianceSet `json:"service-appliance-sets"`
}

type GetServiceApplianceSetRequest struct {
	ID string `json:"id"`
}

type GetServiceApplianceSetResponse struct {
	ServiceApplianceSet *ServiceApplianceSet `json:"service-appliance-set"`
}

func InterfaceToUpdateServiceApplianceSetRequest(i interface{}) *UpdateServiceApplianceSetRequest {
	//TODO implement
	return &UpdateServiceApplianceSetRequest{}
}
