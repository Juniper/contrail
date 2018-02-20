package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateServerRequest struct {
	Server *Server `json:"server"`
}

type CreateServerResponse struct {
	Server *Server `json:"server"`
}

type UpdateServerRequest struct {
	Server    *Server         `json:"server"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateServerResponse struct {
	Server *Server `json:"server"`
}

type DeleteServerRequest struct {
	ID string `json:"id"`
}

type DeleteServerResponse struct {
	ID string `json:"id"`
}

type ListServerRequest struct {
	Spec *ListSpec
}

type ListServerResponse struct {
	Servers []*Server `json:"servers"`
}

type GetServerRequest struct {
	ID string `json:"id"`
}

type GetServerResponse struct {
	Server *Server `json:"server"`
}

func InterfaceToUpdateServerRequest(i interface{}) *UpdateServerRequest {
	//TODO implement
	return &UpdateServerRequest{}
}
