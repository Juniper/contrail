package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateDomainRequest struct {
	Domain *Domain `json:"domain"`
}

type CreateDomainResponse struct {
	Domain *Domain `json:"domain"`
}

type UpdateDomainRequest struct {
	Domain    *Domain         `json:"domain"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateDomainResponse struct {
	Domain *Domain `json:"domain"`
}

type DeleteDomainRequest struct {
	ID string `json:"id"`
}

type DeleteDomainResponse struct {
	ID string `json:"id"`
}

type ListDomainRequest struct {
	Spec *ListSpec
}

type ListDomainResponse struct {
	Domains []*Domain `json:"domains"`
}

type GetDomainRequest struct {
	ID string `json:"id"`
}

type GetDomainResponse struct {
	Domain *Domain `json:"domain"`
}

func InterfaceToUpdateDomainRequest(i interface{}) *UpdateDomainRequest {
	//TODO implement
	return &UpdateDomainRequest{}
}
