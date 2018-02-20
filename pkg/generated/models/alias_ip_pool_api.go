package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateAliasIPPoolRequest struct {
	AliasIPPool *AliasIPPool `json:"alias-ip-pool"`
}

type CreateAliasIPPoolResponse struct {
	AliasIPPool *AliasIPPool `json:"alias-ip-pool"`
}

type UpdateAliasIPPoolRequest struct {
	AliasIPPool *AliasIPPool    `json:"alias-ip-pool"`
	FieldMask   types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateAliasIPPoolResponse struct {
	AliasIPPool *AliasIPPool `json:"alias-ip-pool"`
}

type DeleteAliasIPPoolRequest struct {
	ID string `json:"id"`
}

type DeleteAliasIPPoolResponse struct {
	ID string `json:"id"`
}

type ListAliasIPPoolRequest struct {
	Spec *ListSpec
}

type ListAliasIPPoolResponse struct {
	AliasIPPools []*AliasIPPool `json:"alias-ip-pools"`
}

type GetAliasIPPoolRequest struct {
	ID string `json:"id"`
}

type GetAliasIPPoolResponse struct {
	AliasIPPool *AliasIPPool `json:"alias-ip-pool"`
}

func InterfaceToUpdateAliasIPPoolRequest(i interface{}) *UpdateAliasIPPoolRequest {
	//TODO implement
	return &UpdateAliasIPPoolRequest{}
}
