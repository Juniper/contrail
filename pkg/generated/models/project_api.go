package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateProjectRequest struct {
	Project *Project `json:"project"`
}

type CreateProjectResponse struct {
	Project *Project `json:"project"`
}

type UpdateProjectRequest struct {
	Project   *Project        `json:"project"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateProjectResponse struct {
	Project *Project `json:"project"`
}

type DeleteProjectRequest struct {
	ID string `json:"id"`
}

type DeleteProjectResponse struct {
	ID string `json:"id"`
}

type ListProjectRequest struct {
	Spec *ListSpec
}

type ListProjectResponse struct {
	Projects []*Project `json:"projects"`
}

type GetProjectRequest struct {
	ID string `json:"id"`
}

type GetProjectResponse struct {
	Project *Project `json:"project"`
}

func InterfaceToUpdateProjectRequest(i interface{}) *UpdateProjectRequest {
	//TODO implement
	return &UpdateProjectRequest{}
}
