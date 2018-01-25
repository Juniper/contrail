package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateInstanceIPRequest struct {
	InstanceIP *InstanceIP `json:"instance-ip"`
}

type CreateInstanceIPResponse struct {
	InstanceIP *InstanceIP `json:"instance-ip"`
}

type UpdateInstanceIPRequest struct {
	InstanceIP *InstanceIP     `json:"instance-ip"`
	FieldMask  types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateInstanceIPResponse struct {
	InstanceIP *InstanceIP `json:"instance-ip"`
}

type DeleteInstanceIPRequest struct {
	ID string `json:"id"`
}

type DeleteInstanceIPResponse struct {
	ID string `json:"id"`
}

type ListInstanceIPRequest struct {
	Spec *ListSpec
}

type ListInstanceIPResponse struct {
	InstanceIPs []*InstanceIP `json:"instance-ips"`
}

type GetInstanceIPRequest struct {
	ID string `json:"id"`
}

type GetInstanceIPResponse struct {
	InstanceIP *InstanceIP `json:"instance-ip"`
}

func InterfaceToUpdateInstanceIPRequest(i interface{}) *UpdateInstanceIPRequest {
	//TODO implement
	return &UpdateInstanceIPRequest{}
}
