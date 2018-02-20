package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateGlobalVrouterConfigRequest struct {
	GlobalVrouterConfig *GlobalVrouterConfig `json:"global-vrouter-config"`
}

type CreateGlobalVrouterConfigResponse struct {
	GlobalVrouterConfig *GlobalVrouterConfig `json:"global-vrouter-config"`
}

type UpdateGlobalVrouterConfigRequest struct {
	GlobalVrouterConfig *GlobalVrouterConfig `json:"global-vrouter-config"`
	FieldMask           types.FieldMask      `json:"field_mask,omitempty"`
}

type UpdateGlobalVrouterConfigResponse struct {
	GlobalVrouterConfig *GlobalVrouterConfig `json:"global-vrouter-config"`
}

type DeleteGlobalVrouterConfigRequest struct {
	ID string `json:"id"`
}

type DeleteGlobalVrouterConfigResponse struct {
	ID string `json:"id"`
}

type ListGlobalVrouterConfigRequest struct {
	Spec *ListSpec
}

type ListGlobalVrouterConfigResponse struct {
	GlobalVrouterConfigs []*GlobalVrouterConfig `json:"global-vrouter-configs"`
}

type GetGlobalVrouterConfigRequest struct {
	ID string `json:"id"`
}

type GetGlobalVrouterConfigResponse struct {
	GlobalVrouterConfig *GlobalVrouterConfig `json:"global-vrouter-config"`
}

func InterfaceToUpdateGlobalVrouterConfigRequest(i interface{}) *UpdateGlobalVrouterConfigRequest {
	//TODO implement
	return &UpdateGlobalVrouterConfigRequest{}
}
