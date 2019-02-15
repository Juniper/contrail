package collector

import (
	"context"
	"strings"

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
	Operation string   `json:"oper"`
	RequestID string   `json:"request-id"`
	Type      string   `json:"type"`
	UUID      string   `json:"uuid"`
	FQName    []string `json:"fq_name"`
}

func (p *payloadMessageBusNotifyTrace) Get() *message {
	return &message{
		SandeshType: typeMessageBusNotifyTrace,
		Payload:     p,
	}
}

// MessageBusNotifyTrace sends message with type MessageBusNotifyTrace
func MessageBusNotifyTrace(operation string, typeName string, objUUID string, objFQName []string) messager {
	requestID := "req-" + uuid.NewV4().String()
	return &payloadMessageBusNotifyTrace{
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
	}
}

type processor struct {
	collector Collector
}

func (p *processor) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	p.collector.Send(MessageBusNotifyTrace(
		strings.ToLower(event.Operation()),
		basemodels.KindToSchemaID(event.GetResource().Kind()),
		event.GetResource().GetUUID(),
		event.GetResource().GetFQName(),
	))
	return nil, nil
}

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
