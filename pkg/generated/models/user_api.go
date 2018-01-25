package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateUserRequest struct {
	User *User `json:"user"`
}

type CreateUserResponse struct {
	User *User `json:"user"`
}

type UpdateUserRequest struct {
	User      *User           `json:"user"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateUserResponse struct {
	User *User `json:"user"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type DeleteUserResponse struct {
	ID string `json:"id"`
}

type ListUserRequest struct {
	Spec *ListSpec
}

type ListUserResponse struct {
	Users []*User `json:"users"`
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

func InterfaceToUpdateUserRequest(i interface{}) *UpdateUserRequest {
	//TODO implement
	return &UpdateUserRequest{}
}
