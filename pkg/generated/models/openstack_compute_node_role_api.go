package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateOpenstackComputeNodeRoleRequest struct {
	OpenstackComputeNodeRole *OpenstackComputeNodeRole `json:"openstack-compute-node-role"`
}

type CreateOpenstackComputeNodeRoleResponse struct {
	OpenstackComputeNodeRole *OpenstackComputeNodeRole `json:"openstack-compute-node-role"`
}

type UpdateOpenstackComputeNodeRoleRequest struct {
	OpenstackComputeNodeRole *OpenstackComputeNodeRole `json:"openstack-compute-node-role"`
	FieldMask                types.FieldMask           `json:"field_mask,omitempty"`
}

type UpdateOpenstackComputeNodeRoleResponse struct {
	OpenstackComputeNodeRole *OpenstackComputeNodeRole `json:"openstack-compute-node-role"`
}

type DeleteOpenstackComputeNodeRoleRequest struct {
	ID string `json:"id"`
}

type DeleteOpenstackComputeNodeRoleResponse struct {
	ID string `json:"id"`
}

type ListOpenstackComputeNodeRoleRequest struct {
	Spec *ListSpec
}

type ListOpenstackComputeNodeRoleResponse struct {
	OpenstackComputeNodeRoles []*OpenstackComputeNodeRole `json:"openstack-compute-node-roles"`
}

type GetOpenstackComputeNodeRoleRequest struct {
	ID string `json:"id"`
}

type GetOpenstackComputeNodeRoleResponse struct {
	OpenstackComputeNodeRole *OpenstackComputeNodeRole `json:"openstack-compute-node-role"`
}

func InterfaceToUpdateOpenstackComputeNodeRoleRequest(i interface{}) *UpdateOpenstackComputeNodeRoleRequest {
	//TODO implement
	return &UpdateOpenstackComputeNodeRoleRequest{}
}
