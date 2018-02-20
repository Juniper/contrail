package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateAppformixNodeRoleRequest struct {
	AppformixNodeRole *AppformixNodeRole `json:"appformix-node-role"`
}

type CreateAppformixNodeRoleResponse struct {
	AppformixNodeRole *AppformixNodeRole `json:"appformix-node-role"`
}

type UpdateAppformixNodeRoleRequest struct {
	AppformixNodeRole *AppformixNodeRole `json:"appformix-node-role"`
	FieldMask         types.FieldMask    `json:"field_mask,omitempty"`
}

type UpdateAppformixNodeRoleResponse struct {
	AppformixNodeRole *AppformixNodeRole `json:"appformix-node-role"`
}

type DeleteAppformixNodeRoleRequest struct {
	ID string `json:"id"`
}

type DeleteAppformixNodeRoleResponse struct {
	ID string `json:"id"`
}

type ListAppformixNodeRoleRequest struct {
	Spec *ListSpec
}

type ListAppformixNodeRoleResponse struct {
	AppformixNodeRoles []*AppformixNodeRole `json:"appformix-node-roles"`
}

type GetAppformixNodeRoleRequest struct {
	ID string `json:"id"`
}

type GetAppformixNodeRoleResponse struct {
	AppformixNodeRole *AppformixNodeRole `json:"appformix-node-role"`
}

func InterfaceToUpdateAppformixNodeRoleRequest(i interface{}) *UpdateAppformixNodeRoleRequest {
	//TODO implement
	return &UpdateAppformixNodeRoleRequest{}
}
