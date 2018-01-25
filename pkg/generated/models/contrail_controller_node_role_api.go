package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateContrailControllerNodeRoleRequest struct {
	ContrailControllerNodeRole *ContrailControllerNodeRole `json:"contrail-controller-node-role"`
}

type CreateContrailControllerNodeRoleResponse struct {
	ContrailControllerNodeRole *ContrailControllerNodeRole `json:"contrail-controller-node-role"`
}

type UpdateContrailControllerNodeRoleRequest struct {
	ContrailControllerNodeRole *ContrailControllerNodeRole `json:"contrail-controller-node-role"`
	FieldMask                  types.FieldMask             `json:"field_mask,omitempty"`
}

type UpdateContrailControllerNodeRoleResponse struct {
	ContrailControllerNodeRole *ContrailControllerNodeRole `json:"contrail-controller-node-role"`
}

type DeleteContrailControllerNodeRoleRequest struct {
	ID string `json:"id"`
}

type DeleteContrailControllerNodeRoleResponse struct {
	ID string `json:"id"`
}

type ListContrailControllerNodeRoleRequest struct {
	Spec *ListSpec
}

type ListContrailControllerNodeRoleResponse struct {
	ContrailControllerNodeRoles []*ContrailControllerNodeRole `json:"contrail-controller-node-roles"`
}

type GetContrailControllerNodeRoleRequest struct {
	ID string `json:"id"`
}

type GetContrailControllerNodeRoleResponse struct {
	ContrailControllerNodeRole *ContrailControllerNodeRole `json:"contrail-controller-node-role"`
}

func InterfaceToUpdateContrailControllerNodeRoleRequest(i interface{}) *UpdateContrailControllerNodeRoleRequest {
	//TODO implement
	return &UpdateContrailControllerNodeRoleRequest{}
}
