package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateKeypairRequest struct {
	Keypair *Keypair `json:"keypair"`
}

type CreateKeypairResponse struct {
	Keypair *Keypair `json:"keypair"`
}

type UpdateKeypairRequest struct {
	Keypair   *Keypair        `json:"keypair"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateKeypairResponse struct {
	Keypair *Keypair `json:"keypair"`
}

type DeleteKeypairRequest struct {
	ID string `json:"id"`
}

type DeleteKeypairResponse struct {
	ID string `json:"id"`
}

type ListKeypairRequest struct {
	Spec *ListSpec
}

type ListKeypairResponse struct {
	Keypairs []*Keypair `json:"keypairs"`
}

type GetKeypairRequest struct {
	ID string `json:"id"`
}

type GetKeypairResponse struct {
	Keypair *Keypair `json:"keypair"`
}

func InterfaceToUpdateKeypairRequest(i interface{}) *UpdateKeypairRequest {
	//TODO implement
	return &UpdateKeypairRequest{}
}
