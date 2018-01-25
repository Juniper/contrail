package models

import (
	"github.com/gogo/protobuf/types"
)

type CreatePolicyManagementRequest struct {
	PolicyManagement *PolicyManagement `json:"policy-management"`
}

type CreatePolicyManagementResponse struct {
	PolicyManagement *PolicyManagement `json:"policy-management"`
}

type UpdatePolicyManagementRequest struct {
	PolicyManagement *PolicyManagement `json:"policy-management"`
	FieldMask        types.FieldMask   `json:"field_mask,omitempty"`
}

type UpdatePolicyManagementResponse struct {
	PolicyManagement *PolicyManagement `json:"policy-management"`
}

type DeletePolicyManagementRequest struct {
	ID string `json:"id"`
}

type DeletePolicyManagementResponse struct {
	ID string `json:"id"`
}

type ListPolicyManagementRequest struct {
	Spec *ListSpec
}

type ListPolicyManagementResponse struct {
	PolicyManagements []*PolicyManagement `json:"policy-managements"`
}

type GetPolicyManagementRequest struct {
	ID string `json:"id"`
}

type GetPolicyManagementResponse struct {
	PolicyManagement *PolicyManagement `json:"policy-management"`
}

func InterfaceToUpdatePolicyManagementRequest(i interface{}) *UpdatePolicyManagementRequest {
	//TODO implement
	return &UpdatePolicyManagementRequest{}
}
