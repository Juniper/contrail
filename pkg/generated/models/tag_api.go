package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateTagRequest struct {
	Tag *Tag `json:"tag"`
}

type CreateTagResponse struct {
	Tag *Tag `json:"tag"`
}

type UpdateTagRequest struct {
	Tag       *Tag            `json:"tag"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateTagResponse struct {
	Tag *Tag `json:"tag"`
}

type DeleteTagRequest struct {
	ID string `json:"id"`
}

type DeleteTagResponse struct {
	ID string `json:"id"`
}

type ListTagRequest struct {
	Spec *ListSpec
}

type ListTagResponse struct {
	Tags []*Tag `json:"tags"`
}

type GetTagRequest struct {
	ID string `json:"id"`
}

type GetTagResponse struct {
	Tag *Tag `json:"tag"`
}

func InterfaceToUpdateTagRequest(i interface{}) *UpdateTagRequest {
	//TODO implement
	return &UpdateTagRequest{}
}
