package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateSecurityGroupRequest struct {
	SecurityGroup *SecurityGroup `json:"security-group"`
}

type CreateSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security-group"`
}

type UpdateSecurityGroupRequest struct {
	SecurityGroup *SecurityGroup  `json:"security-group"`
	FieldMask     types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security-group"`
}

type DeleteSecurityGroupRequest struct {
	ID string `json:"id"`
}

type DeleteSecurityGroupResponse struct {
	ID string `json:"id"`
}

type ListSecurityGroupRequest struct {
	Spec *ListSpec
}

type ListSecurityGroupResponse struct {
	SecurityGroups []*SecurityGroup `json:"security-groups"`
}

type GetSecurityGroupRequest struct {
	ID string `json:"id"`
}

type GetSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security-group"`
}

func InterfaceToUpdateSecurityGroupRequest(i interface{}) *UpdateSecurityGroupRequest {
	//TODO implement
	return &UpdateSecurityGroupRequest{}
}
