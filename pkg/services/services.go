package services

import (
	"time"

	context "golang.org/x/net/context"
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
