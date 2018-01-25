package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateBGPAsAServiceRequest struct {
	BGPAsAService *BGPAsAService `json:"bgp-as-a-service"`
}

type CreateBGPAsAServiceResponse struct {
	BGPAsAService *BGPAsAService `json:"bgp-as-a-service"`
}

type UpdateBGPAsAServiceRequest struct {
	BGPAsAService *BGPAsAService  `json:"bgp-as-a-service"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateBGPAsAServiceResponse struct {
	BGPAsAService *BGPAsAService `json:"bgp-as-a-service"`
}

type DeleteBGPAsAServiceRequest struct {
	ID string `json:"id"`
}

type DeleteBGPAsAServiceResponse struct {
	ID string `json:"id"`
}

type ListBGPAsAServiceRequest struct {
	Spec *ListSpec
}

type ListBGPAsAServiceResponse struct {
	BGPAsAServices []*BGPAsAService `json:"bgp-as-a-services"`
}

type GetBGPAsAServiceRequest struct {
	ID string `json:"id"`
}

type GetBGPAsAServiceResponse struct {
	BGPAsAService *BGPAsAService `json:"bgp-as-a-service"`
}

func InterfaceToUpdateBGPAsAServiceRequest(i interface{}) *UpdateBGPAsAServiceRequest {
	//TODO implement
	return &UpdateBGPAsAServiceRequest{}
}
