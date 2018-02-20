package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateBridgeDomainRequest struct {
	BridgeDomain *BridgeDomain `json:"bridge-domain"`
}

type CreateBridgeDomainResponse struct {
	BridgeDomain *BridgeDomain `json:"bridge-domain"`
}

type UpdateBridgeDomainRequest struct {
	BridgeDomain *BridgeDomain   `json:"bridge-domain"`
	FieldMask    types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateBridgeDomainResponse struct {
	BridgeDomain *BridgeDomain `json:"bridge-domain"`
}

type DeleteBridgeDomainRequest struct {
	ID string `json:"id"`
}

type DeleteBridgeDomainResponse struct {
	ID string `json:"id"`
}

type ListBridgeDomainRequest struct {
	Spec *ListSpec
}

type ListBridgeDomainResponse struct {
	BridgeDomains []*BridgeDomain `json:"bridge-domains"`
}

type GetBridgeDomainRequest struct {
	ID string `json:"id"`
}

type GetBridgeDomainResponse struct {
	BridgeDomain *BridgeDomain `json:"bridge-domain"`
}

func InterfaceToUpdateBridgeDomainRequest(i interface{}) *UpdateBridgeDomainRequest {
	//TODO implement
	return &UpdateBridgeDomainRequest{}
}
