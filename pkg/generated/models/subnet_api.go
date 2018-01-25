package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateSubnetRequest struct {
	Subnet *Subnet `json:"subnet"`
}

type CreateSubnetResponse struct {
	Subnet *Subnet `json:"subnet"`
}

type UpdateSubnetRequest struct {
	Subnet    *Subnet         `json:"subnet"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateSubnetResponse struct {
	Subnet *Subnet `json:"subnet"`
}

type DeleteSubnetRequest struct {
	ID string `json:"id"`
}

type DeleteSubnetResponse struct {
	ID string `json:"id"`
}

type ListSubnetRequest struct {
	Spec *ListSpec
}

type ListSubnetResponse struct {
	Subnets []*Subnet `json:"subnets"`
}

type GetSubnetRequest struct {
	ID string `json:"id"`
}

type GetSubnetResponse struct {
	Subnet *Subnet `json:"subnet"`
}

func InterfaceToUpdateSubnetRequest(i interface{}) *UpdateSubnetRequest {
	//TODO implement
	return &UpdateSubnetRequest{}
}
