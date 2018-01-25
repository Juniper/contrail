package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateCustomerAttachmentRequest struct {
	CustomerAttachment *CustomerAttachment `json:"customer-attachment"`
}

type CreateCustomerAttachmentResponse struct {
	CustomerAttachment *CustomerAttachment `json:"customer-attachment"`
}

type UpdateCustomerAttachmentRequest struct {
	CustomerAttachment *CustomerAttachment `json:"customer-attachment"`
	FieldMask          types.FieldMask     `json:"field_mask,omitempty"`
}

type UpdateCustomerAttachmentResponse struct {
	CustomerAttachment *CustomerAttachment `json:"customer-attachment"`
}

type DeleteCustomerAttachmentRequest struct {
	ID string `json:"id"`
}

type DeleteCustomerAttachmentResponse struct {
	ID string `json:"id"`
}

type ListCustomerAttachmentRequest struct {
	Spec *ListSpec
}

type ListCustomerAttachmentResponse struct {
	CustomerAttachments []*CustomerAttachment `json:"customer-attachments"`
}

type GetCustomerAttachmentRequest struct {
	ID string `json:"id"`
}

type GetCustomerAttachmentResponse struct {
	CustomerAttachment *CustomerAttachment `json:"customer-attachment"`
}

func InterfaceToUpdateCustomerAttachmentRequest(i interface{}) *UpdateCustomerAttachmentRequest {
	//TODO implement
	return &UpdateCustomerAttachmentRequest{}
}
