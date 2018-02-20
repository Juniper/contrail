package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateTagTypeRequest struct {
	TagType *TagType `json:"tag-type"`
}

type CreateTagTypeResponse struct {
	TagType *TagType `json:"tag-type"`
}

type UpdateTagTypeRequest struct {
	TagType   *TagType        `json:"tag-type"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateTagTypeResponse struct {
	TagType *TagType `json:"tag-type"`
}

type DeleteTagTypeRequest struct {
	ID string `json:"id"`
}

type DeleteTagTypeResponse struct {
	ID string `json:"id"`
}

type ListTagTypeRequest struct {
	Spec *ListSpec
}

type ListTagTypeResponse struct {
	TagTypes []*TagType `json:"tag-types"`
}

type GetTagTypeRequest struct {
	ID string `json:"id"`
}

type GetTagTypeResponse struct {
	TagType *TagType `json:"tag-type"`
}

func InterfaceToUpdateTagTypeRequest(i interface{}) *UpdateTagTypeRequest {
	//TODO implement
	return &UpdateTagTypeRequest{}
}
