package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateDsaRuleRequest struct {
	DsaRule *DsaRule `json:"dsa-rule"`
}

type CreateDsaRuleResponse struct {
	DsaRule *DsaRule `json:"dsa-rule"`
}

type UpdateDsaRuleRequest struct {
	DsaRule   *DsaRule        `json:"dsa-rule"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateDsaRuleResponse struct {
	DsaRule *DsaRule `json:"dsa-rule"`
}

type DeleteDsaRuleRequest struct {
	ID string `json:"id"`
}

type DeleteDsaRuleResponse struct {
	ID string `json:"id"`
}

type ListDsaRuleRequest struct {
	Spec *ListSpec
}

type ListDsaRuleResponse struct {
	DsaRules []*DsaRule `json:"dsa-rules"`
}

type GetDsaRuleRequest struct {
	ID string `json:"id"`
}

type GetDsaRuleResponse struct {
	DsaRule *DsaRule `json:"dsa-rule"`
}

func InterfaceToUpdateDsaRuleRequest(i interface{}) *UpdateDsaRuleRequest {
	//TODO implement
	return &UpdateDsaRuleRequest{}
}
