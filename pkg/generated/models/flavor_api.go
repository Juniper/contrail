package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateFlavorRequest struct {
	Flavor *Flavor `json:"flavor"`
}

type CreateFlavorResponse struct {
	Flavor *Flavor `json:"flavor"`
}

type UpdateFlavorRequest struct {
	Flavor    *Flavor         `json:"flavor"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateFlavorResponse struct {
	Flavor *Flavor `json:"flavor"`
}

type DeleteFlavorRequest struct {
	ID string `json:"id"`
}

type DeleteFlavorResponse struct {
	ID string `json:"id"`
}

type ListFlavorRequest struct {
	Spec *ListSpec
}

type ListFlavorResponse struct {
	Flavors []*Flavor `json:"flavors"`
}

type GetFlavorRequest struct {
	ID string `json:"id"`
}

type GetFlavorResponse struct {
	Flavor *Flavor `json:"flavor"`
}

func InterfaceToUpdateFlavorRequest(i interface{}) *UpdateFlavorRequest {
	//TODO implement
	return &UpdateFlavorRequest{}
}
