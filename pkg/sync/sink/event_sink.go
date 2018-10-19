package sink

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/parsing"
	"github.com/Juniper/contrail/pkg/services"
)

// EventProcessorSink is a Sink that dispatches it's method calls as events
// to EventProcessor.
type EventProcessorSink struct {
	processor services.EventProcessor
	log       *logrus.Entry
}

// NewEventProcessorSink creates new EventProcessorSink that uses recevied EventProcessor.
func NewEventProcessorSink(processor services.EventProcessor) *EventProcessorSink {
	return &EventProcessorSink{
		processor: processor,
		log:       log.NewLogger("event-processor-sink"),
	}
}

// Create dispatches OperationCreate event to processor.
func (e *EventProcessorSink) Create(ctx context.Context, resourceName string, pk string, obj basedb.Object) error {
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
	attr basedb.Object,
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
		Attr:      json.RawMessage(parsing.MustJSON(attr)),
	})
	if err != nil {
		return err
	}
	return e.process(ctx, ev)
}

// Update dispatches OperationUpdate event to processor.
func (e *EventProcessorSink) Update(ctx context.Context, resourceName string, pk string, obj basedb.Object) error {
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
		Operation: services.RefOperationDelete,
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
	e.log.WithField("type", fmt.Sprintf("%T", ev.Request)).Debug("Dispatching event")
	_, err := e.processor.Process(ctx, ev)
	return err
}
