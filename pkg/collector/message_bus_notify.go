package collector

import (
	uuid "github.com/satori/go.uuid"
)

const typeMessageBusNotifyTrace = "MessageBusNotifyTrace"

type payloadMessageBusNotifyTrace struct {
	RequestID string                     `json:"request_id"`
	Operation string                     `json:"operation"`
	Body      *bodyMessageBusNotifyTrace `json:"body"`
	Error     string                     `json:"error"`
}

type bodyMessageBusNotifyTrace struct {
	Operation string   `json:"oper"`
	RequestID string   `json:"request-id"`
	Type      string   `json:"type"`
	UUID      string   `json:"uuid"`
	FQName    []string `json:"fq_name"`
}

// MessageBusNotifyTrace sends message with type MessageBusNotifyTrace
func (c *Collector) MessageBusNotifyTrace(operation string, typeName string, objUUID string, objFQName []string) {
	requestID := "req-" + uuid.NewV4().String()
	c.sendMessage(&message{
		SandeshType: typeMessageBusNotifyTrace,
		Payload: &payloadMessageBusNotifyTrace{
			RequestID: requestID,
			Operation: operation,
			Body: &bodyMessageBusNotifyTrace{
				Operation: operation,
				RequestID: requestID,
				Type:      typeName,
				UUID:      objUUID,
				FQName:    objFQName,
			},
			Error: "",
		},
	})
}
