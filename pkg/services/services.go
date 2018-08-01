package services

import (
	"context"
	"time"

	"github.com/Juniper/contrail/pkg/models"
)

// Chain setup chain of services.
func Chain(services ...Service) {
	if len(services) < 2 {
		return
	}
	previous := services[0]
	for _, current := range services[1:] {
		previous.SetNext(current)
		previous = current
	}
}

// GetObject retrieves object dynamically from ReadService by its schema ID and uuid.
func GetObject(ctx context.Context, rs ReadService, schemaID, uuid string) (models.Object, error) {
	return getObject(ctx, rs, schemaID, uuid)
}

// BaseService is a service that is a link in service chain and has implemented
// all Service methods as noops. Can be embedded in struct to create new service.
type BaseService struct {
	next Service
}

// Next gets next service to call in service chain.
func (service *BaseService) Next() Service {
	return service.next
}

// SetNext sets next service in service chain.
func (service *BaseService) SetNext(next Service) {
	service.next = next
}

// InTransactionDoer executes do function atomically.
type InTransactionDoer interface {
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
}

// RefUpdateToUpdateService is a service that promotes CreateRef and DeleteRef
// methods to Update method by fetching the object and updating reference
// field with fieldmask applied.
type RefUpdateToUpdateService struct {
	BaseService

	ReadService       ReadService
	InTransactionDoer InTransactionDoer
}

//EventProcessor can handle events on generic way.
type EventProcessor interface {
	Process(ctx context.Context, event *Event) (*Event, error)
}

//EventProducerService can dispatch method call for event processor.
type EventProducerService struct {
	BaseService
	Processor EventProcessor
	Timeout   time.Duration
}

//ServiceEventProcessor dispatch event to method call.
type ServiceEventProcessor struct {
	Service Service
}

//Process processes event.
func (p *ServiceEventProcessor) Process(ctx context.Context, event *Event) (*Event, error) {
	return event.Process(ctx, p.Service)
}
