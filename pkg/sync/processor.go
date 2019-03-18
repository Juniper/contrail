package sync

import (
	"context"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

type ProcessorWithCache struct {
	services.EventProcessor
	idToFQName fqNameCache
}

type fqNameCache map[string][]string

func NewProcessorWithCache(p services.EventProcessor) *ProcessorWithCache {
	return &ProcessorWithCache{
		EventProcessor: p,
		idToFQName:     fqNameCache{},
	}
}

func (p *ProcessorWithCache) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	p.updateFQNameCache(event)

	if err := p.sanitizeCreateRefEvent(event); err != nil {
		return nil, errors.Wrapf(err, "failed to sanitize reference fqName, event: %v", event)
	}

	return p.EventProcessor.Process(ctx, event)
}

func (p *ProcessorWithCache) updateFQNameCache(event *services.Event) {
	switch request := event.ConcreteRequest().(type) {
	case resourceRequest:
		r := request.GetResource()
		p.idToFQName[r.GetUUID()] = r.GetFQName()
	case deleteRequest:
		delete(p.idToFQName, request.GetID())
	}
}

type resourceRequest interface {
	GetResource() basemodels.Object
}

type deleteRequest interface {
	GetID() string
	Kind() string
}

func (p *ProcessorWithCache) sanitizeCreateRefEvent(event *services.Event) error {
	refRequest, ok := event.ConcreteRequest().(services.CreateRefRequest)
	if !ok {
		return nil
	}
	reference := refRequest.GetReference()
	fqName, ok := p.idToFQName[reference.GetUUID()]
	if !ok {
		return errors.Errorf("failed to fetched reference fq_name for uuid: '%v'", reference.GetUUID())
	}
	reference.SetTo(fqName)
	return nil
}
