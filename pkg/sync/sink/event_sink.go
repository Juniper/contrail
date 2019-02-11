package sink

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/logutil"
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
		log:       logutil.NewLogger("event-processor-sink"),
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
	referenceName string,
	pk []string,
	attr basedb.Object,
) error {
	if len(pk) != 2 {
		return errors.Errorf("expecting primary key with 2 items, got %d instead", len(pk))
	}
	var m map[string]interface{}
	if attr != nil {
		m = attr.ToMap()
	}
	ev, err := services.NewRefEvent(&services.RefEventOption{
		Operation: services.RefOperationAdd,
		RefType:   referenceName,
		FromUUID:  pk[0],
		ToUUID:    pk[1],
		Attr:      m,
	})
	if err != nil {
		return err
	}
	return e.process(ctx, ev)
}

// Update dispatches OperationUpdate event to processor.
func (e *EventProcessorSink) Update(
	ctx context.Context, resourceName string, pk string, obj basedb.Object, fm types.FieldMask,
) error {
	ev, err := services.NewEvent(&services.EventOption{
		UUID:      pk,
		Kind:      resourceName,
		Data:      obj.ToMap(),
		Operation: services.OperationUpdate,
		FieldMask: &fm,
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
func (e *EventProcessorSink) DeleteRef(ctx context.Context, referenceName string, pk []string) error {
	if len(pk) != 2 {
		return errors.Errorf("expecting primary key with 2 items, got %d instead", len(pk))
	}
	ev, err := services.NewRefEvent(&services.RefEventOption{
		Operation: services.RefOperationDelete,
		RefType:   referenceName,
		FromUUID:  pk[0],
		ToUUID:    pk[1],
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
