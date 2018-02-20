package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateSecurityLoggingObjectRequest struct {
	SecurityLoggingObject *SecurityLoggingObject `json:"security-logging-object"`
}

type CreateSecurityLoggingObjectResponse struct {
	SecurityLoggingObject *SecurityLoggingObject `json:"security-logging-object"`
}

type UpdateSecurityLoggingObjectRequest struct {
	SecurityLoggingObject *SecurityLoggingObject `json:"security-logging-object"`
	FieldMask             types.FieldMask        `json:"field_mask,omitempty"`
}

type UpdateSecurityLoggingObjectResponse struct {
	SecurityLoggingObject *SecurityLoggingObject `json:"security-logging-object"`
}

type DeleteSecurityLoggingObjectRequest struct {
	ID string `json:"id"`
}

type DeleteSecurityLoggingObjectResponse struct {
	ID string `json:"id"`
}

type ListSecurityLoggingObjectRequest struct {
	Spec *ListSpec
}

type ListSecurityLoggingObjectResponse struct {
	SecurityLoggingObjects []*SecurityLoggingObject `json:"security-logging-objects"`
}

type GetSecurityLoggingObjectRequest struct {
	ID string `json:"id"`
}

type GetSecurityLoggingObjectResponse struct {
	SecurityLoggingObject *SecurityLoggingObject `json:"security-logging-object"`
}

func InterfaceToUpdateSecurityLoggingObjectRequest(i interface{}) *UpdateSecurityLoggingObjectRequest {
	//TODO implement
	return &UpdateSecurityLoggingObjectRequest{}
}
