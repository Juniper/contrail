package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateOpenstackStorageNodeRoleRequest struct {
	OpenstackStorageNodeRole *OpenstackStorageNodeRole `json:"openstack-storage-node-role"`
}

type CreateOpenstackStorageNodeRoleResponse struct {
	OpenstackStorageNodeRole *OpenstackStorageNodeRole `json:"openstack-storage-node-role"`
}

type UpdateOpenstackStorageNodeRoleRequest struct {
	OpenstackStorageNodeRole *OpenstackStorageNodeRole `json:"openstack-storage-node-role"`
	FieldMask                types.FieldMask           `json:"field_mask,omitempty"`
}

type UpdateOpenstackStorageNodeRoleResponse struct {
	OpenstackStorageNodeRole *OpenstackStorageNodeRole `json:"openstack-storage-node-role"`
}

type DeleteOpenstackStorageNodeRoleRequest struct {
	ID string `json:"id"`
}

type DeleteOpenstackStorageNodeRoleResponse struct {
	ID string `json:"id"`
}

type ListOpenstackStorageNodeRoleRequest struct {
	Spec *ListSpec
}

type ListOpenstackStorageNodeRoleResponse struct {
	OpenstackStorageNodeRoles []*OpenstackStorageNodeRole `json:"openstack-storage-node-roles"`
}

type GetOpenstackStorageNodeRoleRequest struct {
	ID string `json:"id"`
}

type GetOpenstackStorageNodeRoleResponse struct {
	OpenstackStorageNodeRole *OpenstackStorageNodeRole `json:"openstack-storage-node-role"`
}

func InterfaceToUpdateOpenstackStorageNodeRoleRequest(i interface{}) *UpdateOpenstackStorageNodeRoleRequest {
	//TODO implement
	return &UpdateOpenstackStorageNodeRoleRequest{}
}
