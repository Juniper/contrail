package sink

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// EventProcessorSink is a Sink that dispatches events to processor.
type EventProcessorSink struct {
	services.EventProcessor
}

// Create dispatches OperationCreate event to processor.
func (e *EventProcessorSink) Create(ctx context.Context, resourceName string, pk string, obj basemodels.Object) error {
	ev, err := services.NewEvent(&services.EventOption{
		UUID:      pk,
		Kind:      resourceName,
		Data:      obj.ToMap(),
		Operation: services.OperationCreate,
	})
	if err != nil {
		return err
	}
	return e.process(ctx, ev)
}

// CreateRef dispatches OperationCreate event to processor for ref_ tables.
func (e *EventProcessorSink) CreateRef(
	ctx context.Context,
	resourceName string,
	pk []string,
	obj basemodels.Object,
) error {
	if len(pk) != 2 {
		return errors.Errorf("expecting primary key with 2 items, got %d instead", len(pk))
	}
	typeName, typeRef := resolveReferenceTable(resourceName)
	ev, err := services.NewEventFromRefUpdate(&services.RefUpdate{
		Operation: services.RefOperationAdd,
		Type:      typeName,
		UUID:      pk[0],
		RefType:   typeRef,
		RefUUID:   pk[1],
	})
	if err != nil {
		return err
	}
	return e.process(ctx, ev)
}

// Update dispatches OperationUpdate event to processor.
func (e *EventProcessorSink) Update(ctx context.Context, resourceName string, pk string, obj basemodels.Object) error {
	ev, err := services.NewEvent(&services.EventOption{
		UUID:      pk,
		Kind:      resourceName,
		Data:      obj.ToMap(),
		Operation: services.OperationUpdate,
	})
	if err != nil {
		return err
	}
	return e.process(ctx, ev)
}

// Delete dispatches OperationDelete event to processor.
func (e *EventProcessorSink) Delete(ctx context.Context, resourceName string, pk string) error {
	ev, err := services.NewEvent(&services.EventOption{
		UUID:      pk,
		Kind:      resourceName,
		Operation: services.OperationDelete,
	})
	if err != nil {
		return err
	}
	return e.process(ctx, ev)
}

// DeleteRef dispatches OperationDelete event to processor for ref_ tables.
func (e *EventProcessorSink) DeleteRef(ctx context.Context, resourceName string, pk []string) error {
	if len(pk) != 2 {
		return errors.Errorf("expecting primary key with 2 items, got %d instead", len(pk))
	}
	typeName, typeRef := resolveReferenceTable(resourceName)
	ev, err := services.NewEventFromRefUpdate(&services.RefUpdate{
		Operation: services.RefOperationAdd,
		Type:      typeName,
		UUID:      pk[0],
		RefType:   typeRef,
		RefUUID:   pk[1],
	})
	if err != nil {
		return err
	}
	return e.process(ctx, ev)
}

func (e *EventProcessorSink) process(ctx context.Context, ev *services.Event) error {
	_, err := e.Process(ctx, ev)
	return err
}
