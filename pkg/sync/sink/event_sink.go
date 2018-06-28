package sink

import (
	"context"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/services"
)

// EventProcessorSink is a Sink that dispatches events to processor.
type EventProcessorSink struct {
	services.EventProcessor
}

// Create dispatches OperationCreate event to processor.
func (e *EventProcessorSink) Create(ctx context.Context, resourceName string, pk string, obj db.Object) error {
	ev := services.NewEvent(&services.EventOption{
		UUID:      pk,
		Kind:      resourceName,
		Data:      obj.ToMap(),
		Operation: services.OperationCreate,
	})
	return e.process(ctx, ev)
}

// Update dispatches OperationUpdate event to processor.
func (e *EventProcessorSink) Update(ctx context.Context, resourceName string, pk string, obj db.Object) error {
	ev := services.NewEvent(&services.EventOption{
		UUID:      pk,
		Kind:      resourceName,
		Data:      obj.ToMap(),
		Operation: services.OperationUpdate,
	})
	return e.process(ctx, ev)
}

// Delete dispatches OperationDelete event to processor.
func (e *EventProcessorSink) Delete(ctx context.Context, resourceName string, pk string) error {
	ev := services.NewEvent(&services.EventOption{
		UUID:      pk,
		Kind:      resourceName,
		Operation: services.OperationDelete,
	})
	return e.process(ctx, ev)
}

func (e *EventProcessorSink) process(ctx context.Context, ev *services.Event) error {
	_, err := e.Process(ctx, ev)
	return err
}
