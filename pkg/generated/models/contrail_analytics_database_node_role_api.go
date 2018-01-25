package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateContrailAnalyticsDatabaseNodeRoleRequest struct {
	ContrailAnalyticsDatabaseNodeRole *ContrailAnalyticsDatabaseNodeRole `json:"contrail-analytics-database-node-role"`
}

type CreateContrailAnalyticsDatabaseNodeRoleResponse struct {
	ContrailAnalyticsDatabaseNodeRole *ContrailAnalyticsDatabaseNodeRole `json:"contrail-analytics-database-node-role"`
}

type UpdateContrailAnalyticsDatabaseNodeRoleRequest struct {
	ContrailAnalyticsDatabaseNodeRole *ContrailAnalyticsDatabaseNodeRole `json:"contrail-analytics-database-node-role"`
	FieldMask                         types.FieldMask                    `json:"field_mask,omitempty"`
}

type UpdateContrailAnalyticsDatabaseNodeRoleResponse struct {
	ContrailAnalyticsDatabaseNodeRole *ContrailAnalyticsDatabaseNodeRole `json:"contrail-analytics-database-node-role"`
}

type DeleteContrailAnalyticsDatabaseNodeRoleRequest struct {
	ID string `json:"id"`
}

type DeleteContrailAnalyticsDatabaseNodeRoleResponse struct {
	ID string `json:"id"`
}

type ListContrailAnalyticsDatabaseNodeRoleRequest struct {
	Spec *ListSpec
}

type ListContrailAnalyticsDatabaseNodeRoleResponse struct {
	ContrailAnalyticsDatabaseNodeRoles []*ContrailAnalyticsDatabaseNodeRole `json:"contrail-analytics-database-node-roles"`
}

type GetContrailAnalyticsDatabaseNodeRoleRequest struct {
	ID string `json:"id"`
}

type GetContrailAnalyticsDatabaseNodeRoleResponse struct {
	ContrailAnalyticsDatabaseNodeRole *ContrailAnalyticsDatabaseNodeRole `json:"contrail-analytics-database-node-role"`
}

func InterfaceToUpdateContrailAnalyticsDatabaseNodeRoleRequest(i interface{}) *UpdateContrailAnalyticsDatabaseNodeRoleRequest {
	//TODO implement
	return &UpdateContrailAnalyticsDatabaseNodeRoleRequest{}
}
