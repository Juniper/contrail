package collector

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

const typeMessageBusNotifyTrace = "MessageBusNotifyTrace"

type payloadMessageBusNotifyTrace struct {
	RequestID string                     `json:"request_id"`
	Operation string                     `json:"operation"`
	Body      *bodyMessageBusNotifyTrace `json:"body"`
	Error     string                     `json:"error"`
}

type bodyMessageBusNotifyTrace struct {
	Operation string      `json:"oper"`
	RequestID string      `json:"request-id"`
	Type      string      `json:"type"`
	UUID      string      `json:"uuid"`
	FQName    []string    `json:"fq_name"`
	ObjDict   interface{} `json:"obj_dict,omitempty"`
}

func (p *payloadMessageBusNotifyTrace) Build() *Message {
	return &Message{
		SandeshType: typeMessageBusNotifyTrace,
		Payload:     p,
	}
}

// MessageBusNotifyTrace sends message with type MessageBusNotifyTrace
func MessageBusNotifyTrace(operation string, obj basemodels.Object) MessageBuilder {
	requestID := "req-" + uuid.NewV4().String()
	var objDict interface{}
	if operation != services.OperationUpdate {
		objDict = obj
	}
	return &payloadMessageBusNotifyTrace{
		RequestID: requestID,
		Operation: operation,
		Body: &bodyMessageBusNotifyTrace{
			Operation: operation,
			RequestID: requestID,
			Type:      basemodels.KindToSchemaID(obj.Kind()),
			UUID:      obj.GetUUID(),
			FQName:    obj.GetFQName(),
			ObjDict:   objDict,
		},
		Error: "",
	}
}

type processor struct {
	collector Collector
}

func (p *processor) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	p.collector.Send(MessageBusNotifyTrace(event.Operation(), event.GetResource()))
	return nil, nil
}

// NewMessageBusProcessor runs etcd events watcher
func NewMessageBusProcessor(collector Collector) error {
	v := &processor{
		collector: collector,
	}
	p, err := etcd.NewEventProducer(v, "collector")
	if err != nil {
		return err
	}
	if err = p.Start(context.Background()); err != nil {
		return err
	}
	return nil
}
