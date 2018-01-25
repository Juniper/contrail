package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateServiceGroupRequest struct {
	ServiceGroup *ServiceGroup `json:"service-group"`
}

type CreateServiceGroupResponse struct {
	ServiceGroup *ServiceGroup `json:"service-group"`
}

type UpdateServiceGroupRequest struct {
	ServiceGroup *ServiceGroup   `json:"service-group"`
	FieldMask    types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateServiceGroupResponse struct {
	ServiceGroup *ServiceGroup `json:"service-group"`
}

type DeleteServiceGroupRequest struct {
	ID string `json:"id"`
}

type DeleteServiceGroupResponse struct {
	ID string `json:"id"`
}

type ListServiceGroupRequest struct {
	Spec *ListSpec
}

type ListServiceGroupResponse struct {
	ServiceGroups []*ServiceGroup `json:"service-groups"`
}

type GetServiceGroupRequest struct {
	ID string `json:"id"`
}

type GetServiceGroupResponse struct {
	ServiceGroup *ServiceGroup `json:"service-group"`
}

func InterfaceToUpdateServiceGroupRequest(i interface{}) *UpdateServiceGroupRequest {
	//TODO implement
	return &UpdateServiceGroupRequest{}
}
