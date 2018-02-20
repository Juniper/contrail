package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateProviderAttachmentRequest struct {
	ProviderAttachment *ProviderAttachment `json:"provider-attachment"`
}

type CreateProviderAttachmentResponse struct {
	ProviderAttachment *ProviderAttachment `json:"provider-attachment"`
}

type UpdateProviderAttachmentRequest struct {
	ProviderAttachment *ProviderAttachment `json:"provider-attachment"`
	FieldMask          types.FieldMask     `json:"field_mask,omitempty"`
}

type UpdateProviderAttachmentResponse struct {
	ProviderAttachment *ProviderAttachment `json:"provider-attachment"`
}

type DeleteProviderAttachmentRequest struct {
	ID string `json:"id"`
}

type DeleteProviderAttachmentResponse struct {
	ID string `json:"id"`
}

type ListProviderAttachmentRequest struct {
	Spec *ListSpec
}

type ListProviderAttachmentResponse struct {
	ProviderAttachments []*ProviderAttachment `json:"provider-attachments"`
}

type GetProviderAttachmentRequest struct {
	ID string `json:"id"`
}

type GetProviderAttachmentResponse struct {
	ProviderAttachment *ProviderAttachment `json:"provider-attachment"`
}

func InterfaceToUpdateProviderAttachmentRequest(i interface{}) *UpdateProviderAttachmentRequest {
	//TODO implement
	return &UpdateProviderAttachmentRequest{}
}
