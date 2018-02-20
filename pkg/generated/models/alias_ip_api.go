package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateAliasIPRequest struct {
	AliasIP *AliasIP `json:"alias-ip"`
}

type CreateAliasIPResponse struct {
	AliasIP *AliasIP `json:"alias-ip"`
}

type UpdateAliasIPRequest struct {
	AliasIP   *AliasIP        `json:"alias-ip"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateAliasIPResponse struct {
	AliasIP *AliasIP `json:"alias-ip"`
}

type DeleteAliasIPRequest struct {
	ID string `json:"id"`
}

type DeleteAliasIPResponse struct {
	ID string `json:"id"`
}

type ListAliasIPRequest struct {
	Spec *ListSpec
}

type ListAliasIPResponse struct {
	AliasIPs []*AliasIP `json:"alias-ips"`
}

type GetAliasIPRequest struct {
	ID string `json:"id"`
}

type GetAliasIPResponse struct {
	AliasIP *AliasIP `json:"alias-ip"`
}

func InterfaceToUpdateAliasIPRequest(i interface{}) *UpdateAliasIPRequest {
	//TODO implement
	return &UpdateAliasIPRequest{}
}
