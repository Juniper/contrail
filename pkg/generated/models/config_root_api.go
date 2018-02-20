package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateConfigRootRequest struct {
	ConfigRoot *ConfigRoot `json:"config-root"`
}

type CreateConfigRootResponse struct {
	ConfigRoot *ConfigRoot `json:"config-root"`
}

type UpdateConfigRootRequest struct {
	ConfigRoot *ConfigRoot     `json:"config-root"`
	FieldMask  types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateConfigRootResponse struct {
	ConfigRoot *ConfigRoot `json:"config-root"`
}

type DeleteConfigRootRequest struct {
	ID string `json:"id"`
}

type DeleteConfigRootResponse struct {
	ID string `json:"id"`
}

type ListConfigRootRequest struct {
	Spec *ListSpec
}

type ListConfigRootResponse struct {
	ConfigRoots []*ConfigRoot `json:"config-roots"`
}

type GetConfigRootRequest struct {
	ID string `json:"id"`
}

type GetConfigRootResponse struct {
	ConfigRoot *ConfigRoot `json:"config-root"`
}

func InterfaceToUpdateConfigRootRequest(i interface{}) *UpdateConfigRootRequest {
	//TODO implement
	return &UpdateConfigRootRequest{}
}
