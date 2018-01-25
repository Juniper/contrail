package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateVPNGroupRequest struct {
	VPNGroup *VPNGroup `json:"vpn-group"`
}

type CreateVPNGroupResponse struct {
	VPNGroup *VPNGroup `json:"vpn-group"`
}

type UpdateVPNGroupRequest struct {
	VPNGroup  *VPNGroup       `json:"vpn-group"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateVPNGroupResponse struct {
	VPNGroup *VPNGroup `json:"vpn-group"`
}

type DeleteVPNGroupRequest struct {
	ID string `json:"id"`
}

type DeleteVPNGroupResponse struct {
	ID string `json:"id"`
}

type ListVPNGroupRequest struct {
	Spec *ListSpec
}

type ListVPNGroupResponse struct {
	VPNGroups []*VPNGroup `json:"vpn-groups"`
}

type GetVPNGroupRequest struct {
	ID string `json:"id"`
}

type GetVPNGroupResponse struct {
	VPNGroup *VPNGroup `json:"vpn-group"`
}

func InterfaceToUpdateVPNGroupRequest(i interface{}) *UpdateVPNGroupRequest {
	//TODO implement
	return &UpdateVPNGroupRequest{}
}
