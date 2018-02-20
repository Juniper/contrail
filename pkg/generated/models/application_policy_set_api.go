package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateApplicationPolicySetRequest struct {
	ApplicationPolicySet *ApplicationPolicySet `json:"application-policy-set"`
}

type CreateApplicationPolicySetResponse struct {
	ApplicationPolicySet *ApplicationPolicySet `json:"application-policy-set"`
}

type UpdateApplicationPolicySetRequest struct {
	ApplicationPolicySet *ApplicationPolicySet `json:"application-policy-set"`
	FieldMask            types.FieldMask       `json:"field_mask,omitempty"`
}

type UpdateApplicationPolicySetResponse struct {
	ApplicationPolicySet *ApplicationPolicySet `json:"application-policy-set"`
}

type DeleteApplicationPolicySetRequest struct {
	ID string `json:"id"`
}

type DeleteApplicationPolicySetResponse struct {
	ID string `json:"id"`
}

type ListApplicationPolicySetRequest struct {
	Spec *ListSpec
}

type ListApplicationPolicySetResponse struct {
	ApplicationPolicySets []*ApplicationPolicySet `json:"application-policy-sets"`
}

type GetApplicationPolicySetRequest struct {
	ID string `json:"id"`
}

type GetApplicationPolicySetResponse struct {
	ApplicationPolicySet *ApplicationPolicySet `json:"application-policy-set"`
}

func InterfaceToUpdateApplicationPolicySetRequest(i interface{}) *UpdateApplicationPolicySetRequest {
	//TODO implement
	return &UpdateApplicationPolicySetRequest{}
}
