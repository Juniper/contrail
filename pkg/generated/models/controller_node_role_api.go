package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateControllerNodeRoleRequest struct {
	ControllerNodeRole *ControllerNodeRole `json:"controller-node-role"`
}

type CreateControllerNodeRoleResponse struct {
	ControllerNodeRole *ControllerNodeRole `json:"controller-node-role"`
}

type UpdateControllerNodeRoleRequest struct {
	ControllerNodeRole *ControllerNodeRole `json:"controller-node-role"`
	FieldMask          types.FieldMask     `json:"field_mask,omitempty"`
}

type UpdateControllerNodeRoleResponse struct {
	ControllerNodeRole *ControllerNodeRole `json:"controller-node-role"`
}

type DeleteControllerNodeRoleRequest struct {
	ID string `json:"id"`
}

type DeleteControllerNodeRoleResponse struct {
	ID string `json:"id"`
}

type ListControllerNodeRoleRequest struct {
	Spec *ListSpec
}

type ListControllerNodeRoleResponse struct {
	ControllerNodeRoles []*ControllerNodeRole `json:"controller-node-roles"`
}

type GetControllerNodeRoleRequest struct {
	ID string `json:"id"`
}

type GetControllerNodeRoleResponse struct {
	ControllerNodeRole *ControllerNodeRole `json:"controller-node-role"`
}

func InterfaceToUpdateControllerNodeRoleRequest(i interface{}) *UpdateControllerNodeRoleRequest {
	//TODO implement
	return &UpdateControllerNodeRoleRequest{}
}
