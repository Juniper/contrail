package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateConfigNodeRequest struct {
	ConfigNode *ConfigNode `json:"config-node"`
}

type CreateConfigNodeResponse struct {
	ConfigNode *ConfigNode `json:"config-node"`
}

type UpdateConfigNodeRequest struct {
	ConfigNode *ConfigNode     `json:"config-node"`
	FieldMask  types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateConfigNodeResponse struct {
	ConfigNode *ConfigNode `json:"config-node"`
}

type DeleteConfigNodeRequest struct {
	ID string `json:"id"`
}

type DeleteConfigNodeResponse struct {
	ID string `json:"id"`
}

type ListConfigNodeRequest struct {
	Spec *ListSpec
}

type ListConfigNodeResponse struct {
	ConfigNodes []*ConfigNode `json:"config-nodes"`
}

type GetConfigNodeRequest struct {
	ID string `json:"id"`
}

type GetConfigNodeResponse struct {
	ConfigNode *ConfigNode `json:"config-node"`
}

func InterfaceToUpdateConfigNodeRequest(i interface{}) *UpdateConfigNodeRequest {
	//TODO implement
	return &UpdateConfigNodeRequest{}
}
