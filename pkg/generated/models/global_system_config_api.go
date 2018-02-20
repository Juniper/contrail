package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateGlobalSystemConfigRequest struct {
	GlobalSystemConfig *GlobalSystemConfig `json:"global-system-config"`
}

type CreateGlobalSystemConfigResponse struct {
	GlobalSystemConfig *GlobalSystemConfig `json:"global-system-config"`
}

type UpdateGlobalSystemConfigRequest struct {
	GlobalSystemConfig *GlobalSystemConfig `json:"global-system-config"`
	FieldMask          types.FieldMask     `json:"field_mask,omitempty"`
}

type UpdateGlobalSystemConfigResponse struct {
	GlobalSystemConfig *GlobalSystemConfig `json:"global-system-config"`
}

type DeleteGlobalSystemConfigRequest struct {
	ID string `json:"id"`
}

type DeleteGlobalSystemConfigResponse struct {
	ID string `json:"id"`
}

type ListGlobalSystemConfigRequest struct {
	Spec *ListSpec
}

type ListGlobalSystemConfigResponse struct {
	GlobalSystemConfigs []*GlobalSystemConfig `json:"global-system-configs"`
}

type GetGlobalSystemConfigRequest struct {
	ID string `json:"id"`
}

type GetGlobalSystemConfigResponse struct {
	GlobalSystemConfig *GlobalSystemConfig `json:"global-system-config"`
}

func InterfaceToUpdateGlobalSystemConfigRequest(i interface{}) *UpdateGlobalSystemConfigRequest {
	//TODO implement
	return &UpdateGlobalSystemConfigRequest{}
}
