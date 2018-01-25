package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateBGPVPNRequest struct {
	BGPVPN *BGPVPN `json:"bgpvpn"`
}

type CreateBGPVPNResponse struct {
	BGPVPN *BGPVPN `json:"bgpvpn"`
}

type UpdateBGPVPNRequest struct {
	BGPVPN    *BGPVPN         `json:"bgpvpn"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateBGPVPNResponse struct {
	BGPVPN *BGPVPN `json:"bgpvpn"`
}

type DeleteBGPVPNRequest struct {
	ID string `json:"id"`
}

type DeleteBGPVPNResponse struct {
	ID string `json:"id"`
}

type ListBGPVPNRequest struct {
	Spec *ListSpec
}

type ListBGPVPNResponse struct {
	BGPVPNs []*BGPVPN `json:"bgpvpns"`
}

type GetBGPVPNRequest struct {
	ID string `json:"id"`
}

type GetBGPVPNResponse struct {
	BGPVPN *BGPVPN `json:"bgpvpn"`
}

func InterfaceToUpdateBGPVPNRequest(i interface{}) *UpdateBGPVPNRequest {
	//TODO implement
	return &UpdateBGPVPNRequest{}
}
