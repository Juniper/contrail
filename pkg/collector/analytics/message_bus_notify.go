package analytics

import (
	"context"

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/etcd"
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

func (p *payloadMessageBusNotifyTrace) Build() *collector.Message {
	return &collector.Message{
		SandeshType: typeMessageBusNotifyTrace,
		Payload:     p,
	}
}

// MessageBusNotifyTrace sends message with type MessageBusNotifyTrace
func MessageBusNotifyTrace(ctx context.Context, operation string, obj basemodels.Object) collector.MessageBuilder {
	/* TODO: Should be reverted as introspect service for Intent API will be introduced.
	requestID := services.GetRequestID(ctx)
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
	*/
	return collector.NewEmptyMessageBuilder()
}

type processor struct {
	collector collector.Collector
}

func (p *processor) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	p.collector.Send(MessageBusNotifyTrace(ctx, event.Operation(), event.GetResource()))
	return nil, nil
}

// NewMessageBusProcessor runs etcd events watcher
func NewMessageBusProcessor(collector collector.Collector) error {
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
