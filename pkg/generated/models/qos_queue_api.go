package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateQosQueueRequest struct {
	QosQueue *QosQueue `json:"qos-queue"`
}

type CreateQosQueueResponse struct {
	QosQueue *QosQueue `json:"qos-queue"`
}

type UpdateQosQueueRequest struct {
	QosQueue  *QosQueue       `json:"qos-queue"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateQosQueueResponse struct {
	QosQueue *QosQueue `json:"qos-queue"`
}

type DeleteQosQueueRequest struct {
	ID string `json:"id"`
}

type DeleteQosQueueResponse struct {
	ID string `json:"id"`
}

type ListQosQueueRequest struct {
	Spec *ListSpec
}

type ListQosQueueResponse struct {
	QosQueues []*QosQueue `json:"qos-queues"`
}

type GetQosQueueRequest struct {
	ID string `json:"id"`
}

type GetQosQueueResponse struct {
	QosQueue *QosQueue `json:"qos-queue"`
}

func InterfaceToUpdateQosQueueRequest(i interface{}) *UpdateQosQueueRequest {
	//TODO implement
	return &UpdateQosQueueRequest{}
}
