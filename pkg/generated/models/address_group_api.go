package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateAddressGroupRequest struct {
	AddressGroup *AddressGroup `json:"address-group"`
}

type CreateAddressGroupResponse struct {
	AddressGroup *AddressGroup `json:"address-group"`
}

type UpdateAddressGroupRequest struct {
	AddressGroup *AddressGroup   `json:"address-group"`
	FieldMask    types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateAddressGroupResponse struct {
	AddressGroup *AddressGroup `json:"address-group"`
}

type DeleteAddressGroupRequest struct {
	ID string `json:"id"`
}

type DeleteAddressGroupResponse struct {
	ID string `json:"id"`
}

type ListAddressGroupRequest struct {
	Spec *ListSpec
}

type ListAddressGroupResponse struct {
	AddressGroups []*AddressGroup `json:"address-groups"`
}

type GetAddressGroupRequest struct {
	ID string `json:"id"`
}

type GetAddressGroupResponse struct {
	AddressGroup *AddressGroup `json:"address-group"`
}

func InterfaceToUpdateAddressGroupRequest(i interface{}) *UpdateAddressGroupRequest {
	//TODO implement
	return &UpdateAddressGroupRequest{}
}
