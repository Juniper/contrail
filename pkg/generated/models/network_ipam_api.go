package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateNetworkIpamRequest struct {
	NetworkIpam *NetworkIpam `json:"network-ipam"`
}

type CreateNetworkIpamResponse struct {
	NetworkIpam *NetworkIpam `json:"network-ipam"`
}

type UpdateNetworkIpamRequest struct {
	NetworkIpam *NetworkIpam    `json:"network-ipam"`
	FieldMask   types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateNetworkIpamResponse struct {
	NetworkIpam *NetworkIpam `json:"network-ipam"`
}

type DeleteNetworkIpamRequest struct {
	ID string `json:"id"`
}

type DeleteNetworkIpamResponse struct {
	ID string `json:"id"`
}

type ListNetworkIpamRequest struct {
	Spec *ListSpec
}

type ListNetworkIpamResponse struct {
	NetworkIpams []*NetworkIpam `json:"network-ipams"`
}

type GetNetworkIpamRequest struct {
	ID string `json:"id"`
}

type GetNetworkIpamResponse struct {
	NetworkIpam *NetworkIpam `json:"network-ipam"`
}

func InterfaceToUpdateNetworkIpamRequest(i interface{}) *UpdateNetworkIpamRequest {
	//TODO implement
	return &UpdateNetworkIpamRequest{}
}
