package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateDatabaseNodeRequest struct {
	DatabaseNode *DatabaseNode `json:"database-node"`
}

type CreateDatabaseNodeResponse struct {
	DatabaseNode *DatabaseNode `json:"database-node"`
}

type UpdateDatabaseNodeRequest struct {
	DatabaseNode *DatabaseNode   `json:"database-node"`
	FieldMask    types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateDatabaseNodeResponse struct {
	DatabaseNode *DatabaseNode `json:"database-node"`
}

type DeleteDatabaseNodeRequest struct {
	ID string `json:"id"`
}

type DeleteDatabaseNodeResponse struct {
	ID string `json:"id"`
}

type ListDatabaseNodeRequest struct {
	Spec *ListSpec
}

type ListDatabaseNodeResponse struct {
	DatabaseNodes []*DatabaseNode `json:"database-nodes"`
}

type GetDatabaseNodeRequest struct {
	ID string `json:"id"`
}

type GetDatabaseNodeResponse struct {
	DatabaseNode *DatabaseNode `json:"database-node"`
}

func InterfaceToUpdateDatabaseNodeRequest(i interface{}) *UpdateDatabaseNodeRequest {
	//TODO implement
	return &UpdateDatabaseNodeRequest{}
}
