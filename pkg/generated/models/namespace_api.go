package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateNamespaceRequest struct {
	Namespace *Namespace `json:"namespace"`
}

type CreateNamespaceResponse struct {
	Namespace *Namespace `json:"namespace"`
}

type UpdateNamespaceRequest struct {
	Namespace *Namespace      `json:"namespace"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateNamespaceResponse struct {
	Namespace *Namespace `json:"namespace"`
}

type DeleteNamespaceRequest struct {
	ID string `json:"id"`
}

type DeleteNamespaceResponse struct {
	ID string `json:"id"`
}

type ListNamespaceRequest struct {
	Spec *ListSpec
}

type ListNamespaceResponse struct {
	Namespaces []*Namespace `json:"namespaces"`
}

type GetNamespaceRequest struct {
	ID string `json:"id"`
}

type GetNamespaceResponse struct {
	Namespace *Namespace `json:"namespace"`
}

func InterfaceToUpdateNamespaceRequest(i interface{}) *UpdateNamespaceRequest {
	//TODO implement
	return &UpdateNamespaceRequest{}
}
