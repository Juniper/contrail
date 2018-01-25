package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateWidgetRequest struct {
	Widget *Widget `json:"widget"`
}

type CreateWidgetResponse struct {
	Widget *Widget `json:"widget"`
}

type UpdateWidgetRequest struct {
	Widget    *Widget         `json:"widget"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateWidgetResponse struct {
	Widget *Widget `json:"widget"`
}

type DeleteWidgetRequest struct {
	ID string `json:"id"`
}

type DeleteWidgetResponse struct {
	ID string `json:"id"`
}

type ListWidgetRequest struct {
	Spec *ListSpec
}

type ListWidgetResponse struct {
	Widgets []*Widget `json:"widgets"`
}

type GetWidgetRequest struct {
	ID string `json:"id"`
}

type GetWidgetResponse struct {
	Widget *Widget `json:"widget"`
}

func InterfaceToUpdateWidgetRequest(i interface{}) *UpdateWidgetRequest {
	//TODO implement
	return &UpdateWidgetRequest{}
}
