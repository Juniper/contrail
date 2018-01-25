package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateAlarmRequest struct {
	Alarm *Alarm `json:"alarm"`
}

type CreateAlarmResponse struct {
	Alarm *Alarm `json:"alarm"`
}

type UpdateAlarmRequest struct {
	Alarm     *Alarm          `json:"alarm"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateAlarmResponse struct {
	Alarm *Alarm `json:"alarm"`
}

type DeleteAlarmRequest struct {
	ID string `json:"id"`
}

type DeleteAlarmResponse struct {
	ID string `json:"id"`
}

type ListAlarmRequest struct {
	Spec *ListSpec
}

type ListAlarmResponse struct {
	Alarms []*Alarm `json:"alarms"`
}

type GetAlarmRequest struct {
	ID string `json:"id"`
}

type GetAlarmResponse struct {
	Alarm *Alarm `json:"alarm"`
}

func InterfaceToUpdateAlarmRequest(i interface{}) *UpdateAlarmRequest {
	//TODO implement
	return &UpdateAlarmRequest{}
}
